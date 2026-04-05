package input

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/pkg/model"
)

// ParseArgs parses the command line arguments into a Config struct.
// It handles flags for color, output, and alignment, and validates usage.
func ParseArgs(args []string) (*model.Config, error) {
	config := &model.Config{
		BannerFile: "standard",
		Align:      "left",
	}

	var remainingArgs []string

	// 1. Parse Flags
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg, "=", 2)
			flag := parts[0]
			if flag == "--reverse" {
				if len(parts) != 2 || parts[1] == "" {
					return nil, fmt.Errorf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
				}
				if _, err := os.Stat(parts[1]); os.IsNotExist(err) {
					return nil, fmt.Errorf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n\nFile not found: %s", parts[1])
				}
				config.ReverseFile = parts[1]
			} else if flag == "--color" {
				if len(parts) != 2 {
					return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
				}
				config.Color = parts[1]
			} else if flag == "--output" {
				if len(parts) != 2 {
					return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run ./cmd/ascii-art --output=<fileName.txt> something standard")
				}
				config.OutputFile = parts[1]
			} else if flag == "--align" {
				if len(parts) != 2 || !isAlign(parts[1]) {
					return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				}
				config.Align = parts[1]
			} else {
				// Treat unknown flags as arguments or ignore, depending on strictness.
				// For now, we assume they might be part of the string if not matched.
				remainingArgs = append(remainingArgs, arg)
			}
		} else {
			remainingArgs = append(remainingArgs, arg)
		}
	}

	if len(remainingArgs) == 0 {
		if config.ReverseFile != "" {
			return config, nil
		}
		return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run . something standard")
	}

	// In reverse mode, the only optional positional arg is the banner name
	if config.ReverseFile != "" {
		if len(remainingArgs) == 1 && isBanner(remainingArgs[0]) {
			config.BannerFile = remainingArgs[0]
		} else if len(remainingArgs) > 0 {
			return nil, fmt.Errorf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		}
		return config, nil
	}

	// 2. Parse Positional Arguments
	if config.Color != "" {
		// Color mode has specific argument structure: [ColorSubstr] [Input] [Banner]
		switch len(remainingArgs) {
		case 1:
			config.Input = remainingArgs[0]
		case 2:
			if isBanner(remainingArgs[1]) {
				config.Input = remainingArgs[0]
				config.BannerFile = remainingArgs[1]
			} else {
				config.ColorSubstr = remainingArgs[0]
				config.Input = remainingArgs[1]
			}
		case 3:
			config.ColorSubstr = remainingArgs[0]
			config.Input = remainingArgs[1]
			config.BannerFile = remainingArgs[2]
		default:
			return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <substring> <string> <banner>")
		}
	} else {
		// Standard mode: [Input] [Banner]
		switch len(remainingArgs) {
		case 1:
			config.Input = remainingArgs[0]
		case 2:
			config.Input = remainingArgs[0]
			config.BannerFile = remainingArgs[1]
		default:
			// If too many args, return specific error based on active flag
			if config.OutputFile != "" {
				return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			}
			if config.Align != "left" {
				return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			}
			return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run . something standard")
		}
	}

	// 3. Process Input (Handle escaped newlines)
	config.Input = strings.ReplaceAll(config.Input, "\\n", "\n")

	if !isBanner(config.BannerFile) {
		return nil, fmt.Errorf("Usage: go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]\n\nEX: go run . something standard\n\nAvailable banners: standard, shadow, thinkertoy")
	}

	return config, nil
}

func isBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}

func isAlign(name string) bool {
	switch name {
	case "left", "center", "right", "justify":
		return true
	}
	return false
}
