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

	// For now, assume the first positional arg is the input string.
	// Complex logic (Banner selection, Color substring) will be added in Cycle 3.
	cfg.Input = positionalArgs[0]

	// Basic validation for input string (reuse existing logic if needed, or simple check)
	if cfg.Input == "" {
		// Depending on requirements, empty string might be valid or not.
		// Current parser allows it but returns empty output.
	}

	return cfg, nil
}