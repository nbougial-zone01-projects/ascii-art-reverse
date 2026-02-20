package render

import (
	"ascii-art/pkg/model"
	"strings"
)

// Render converts the input string into an ASCII art string using the provided banner.
func Render(input string, banner model.Banner) string {
	// Split input into lines to handle newlines correctly
	inputLines := strings.Split(input, "\n")
	var sb strings.Builder

	for i, line := range inputLines {
		// Add a newline between blocks of text, but not before the first one
		if i > 0 {
			sb.WriteByte('\n')
		}
		// If the line is empty, we just needed the newline added above (if i > 0)
		if line == "" {
			continue
		}
		var lines [8]strings.Builder
		// Iterate over each character in the current line of text
		for _, char := range line {
			if asciiLines, ok := banner[char]; ok {
				// Append each of the 8 rows of the character to the corresponding builder
				for row := 0; row < 8; row++ {
					lines[row].WriteString(asciiLines[row])
				}
			}
		}
		// Combine the 8 rows into the final output builder
		for row := 0; row < 8; row++ {
			sb.WriteString(lines[row].String())
			if row < 7 {
				sb.WriteByte('\n')
			}
		}
	}

	return sb.String()
}