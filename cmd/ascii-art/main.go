package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/input"
	"ascii-art/internal/render"
	"ascii-art/pkg/model"
	"fmt"
	"io"
	"os"
)

const defaultBannerPath = "assets/banners/standard.txt"

// Run wires input parsing, banner loading, and rendering.
// It writes the result to the provided writer (usually stdout).
func Run(args []string, stdout io.Writer) error {
	// 1. Parse and validate input arguments
	parsed, err := input.ParseInput(args)
	if err != nil {
		return err
	}

	// 2. Load the banner font from the assets directory
	b, err := banner.LoadBanner(defaultBannerPath)
	if err != nil {
		return err
	}

	// 3. Render the string using the loaded banner and print to output
	cfg := &model.Config{
		Input:      parsed,
		BannerFile: "standard",
	}
	_, err = fmt.Fprint(stdout, render.Render(cfg, b))
	return err
}

func main() {
	// Delegate to Run for better testability and error handling
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
