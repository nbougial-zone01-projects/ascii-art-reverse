package main

import (
	"fmt"
	"io"
	"os"

	"ascii-art/internal/banner"
	"ascii-art/internal/input"
	"ascii-art/internal/output"
	"ascii-art/internal/render"
	"ascii-art/internal/reverse"
)

// Run wires input parsing, banner loading, and rendering.
// It writes the result to the provided writer (usually stdout) or to a file if OutputFile is set.
func Run(args []string, stdout io.Writer) error {
	// 1. Parse and validate input arguments
	cfg, err := input.ParseArgs(args)
	if err != nil {
		return err
	}

	// 2. Resolve and load the selected banner font
	bannerPath, err := banner.GetBannerPath(cfg.BannerFile)
	if err != nil {
		return err
	}
	b, err := banner.LoadBanner(bannerPath)
	if err != nil {
		return err
	}

	// 3. Reverse mode: decode ASCII art file back to text
	if cfg.ReverseFile != "" {
		data, err := os.ReadFile(cfg.ReverseFile)
		if err != nil {
			return err
		}
		result, err := reverse.Reverse(string(data), b)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(stdout, result)
		return err
	}

	// 4. Render the string using the loaded banner
	result, err := render.Render(cfg, b)
	if err != nil {
		return err
	}

	// 4. Write to file or stdout based on Config.OutputFile
	if cfg.OutputFile != "" {
		return output.WriteOutput(cfg.OutputFile, result)
	}

	_, err = fmt.Fprintln(stdout, result)
	return err
}

func main() {
	// Delegate to Run for better testability and error handling
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
