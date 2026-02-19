package render

import (
	"ascii-art/pkg/model"
	"strings"
)

// Render converts the input string into an ASCII art string using the provided banner.
func Render(input string, banner model.Banner) string {
	outputLines := make([]string, 8)

	for _, char := range input {
		if lines, ok := banner[char]; ok {
			for i := 0; i < 8; i++ {
				outputLines[i] += lines[i]
			}
		}
	}

	return strings.Join(outputLines, "\n")
}