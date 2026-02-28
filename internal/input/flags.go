package input

import (
	"ascii-art/pkg/model"
	"strings"
)

// ParseArgs parses command line arguments into a Config object.
// It handles flags (--color, --output, --align) and positional arguments.
func ParseArgs(args []string) (*model.Config, error) {
	cfg := &model.Config{
		BannerFile: "standard",
		Align:      "left",
	}

	var positionalArgs []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			// Handle Flags
			parts := strings.SplitN(arg[2:], "=", 2)
			key := parts[0]
			val := ""
			if len(parts) > 1 {
				val = parts[1]
			}

			switch key {
			case "output":
				cfg.OutputFile = val
			case "color":
				cfg.Color = val
			case "align":
				cfg.Align = val
			}
		} else {
			// Handle Positional Args
			positionalArgs = append(positionalArgs, arg)
		}
	}

	// Validate and assign positional arguments
	if len(positionalArgs) == 0 {
		return nil, ErrUsage
	}

	switch len(positionalArgs) {
	case 1:
		// [STRING]
		cfg.Input = positionalArgs[0]
	case 2:
		// If --color is set, it could be [SUBSTRING] [STRING]
		// Otherwise, it is [STRING] [BANNER]
		if cfg.Color != "" {
			cfg.ColorSubstr = positionalArgs[0]
			cfg.Input = positionalArgs[1]
		} else {
			cfg.Input = positionalArgs[0]
			cfg.BannerFile = positionalArgs[1]
		}
	case 3:
		// Must be [SUBSTRING] [STRING] [BANNER] and --color must be set
		if cfg.Color != "" {
			cfg.ColorSubstr = positionalArgs[0]
			cfg.Input = positionalArgs[1]
			cfg.BannerFile = positionalArgs[2]
		} else {
			return nil, ErrUsage
		}
	default:
		return nil, ErrUsage
	}

	return cfg, nil
}