package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/input"
	"ascii-art/internal/render"
	"fmt"
	"io"
	"os"
)

const defaultBannerPath = "assets/banners/ascii_library.txt"

// Run wires input parsing, banner loading, and rendering.
func Run(args []string, stdout io.Writer) error {
	parsed, err := input.ParseInput(args)
	if err != nil {
		return err
	}

	b, err := banner.LoadBanner(defaultBannerPath)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(stdout, render.Render(parsed, b))
	return err
}

func main() {
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
