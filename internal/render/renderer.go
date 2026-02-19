package render

import (
	"ascii-art/pkg/model"
	"strings"
)

// Render converts the input string into an ASCII art string using the provided banner.
func Render(input string, banner model.Banner) string {
	inputLines := strings.Split(input, "\n")
	var sb strings.Builder

	for i, line := range inputLines {
		if i > 0 {
			sb.WriteByte('\n')
		}
		if line == "" {
			continue
		}
		var lines [8]strings.Builder
		for _, char := range line {
			if asciiLines, ok := banner[char]; ok {
				for row := 0; row < 8; row++ {
					lines[row].WriteString(asciiLines[row])
				}
			}
		}
		for row := 0; row < 8; row++ {
			sb.WriteString(lines[row].String())
			if row < 7 {
				sb.WriteByte('\n')
			}
		}
	}

	return sb.String()
}