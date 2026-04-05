package reverse

import (
	"ascii-art/pkg/model"
	"fmt"
	"strings"
)

// invertBanner builds a lookup map from a serialized glyph string to its rune.
func invertBanner(b model.Banner) map[string]rune {
	inv := make(map[string]rune, len(b))
	for r, lines := range b {
		inv[strings.Join(lines, "\n")] = r
	}
	return inv
}

// Reverse reconstructs the original input string from ASCII art produced by this tool.
// art is the full art string (8 rows per input line, blocks separated by a blank line).
func Reverse(art string, b model.Banner) (string, error) {
	if art == "" {
		return "", nil
	}

	inv := invertBanner(b)

	// Split into per-input-line blocks separated by a blank line between 8-row groups.
	// Each block is exactly 8 rows followed by an optional blank line separator.
	rows := strings.Split(art, "\n")
	var blocks [][]string
	for i := 0; i < len(rows); {
		if i+8 > len(rows) {
			break
		}
		blocks = append(blocks, rows[i:i+8])
		i += 8
		// Skip the blank separator line if present
		if i < len(rows) && rows[i] == "" {
			i++
		}
	}

	var resultLines []string
	for _, block := range blocks {
		line, err := decodeLine(block, inv)
		if err != nil {
			return "", err
		}
		resultLines = append(resultLines, line)
	}

	return strings.Join(resultLines, "\n"), nil
}

// decodeLine reconstructs a single line of text from its 8-row ASCII art block.
func decodeLine(rows []string, inv map[string]rune) (string, error) {
	if len(rows) == 0 {
		return "", nil
	}

	// Determine glyph width by finding the width of the first glyph in the banner.
	// All glyphs in a block share the same column structure.
	// We detect character boundaries by slicing columns of equal width.
	// Since glyph widths vary, we use the inverted map to match greedily.
	lineLen := len(rows[0])
	var result strings.Builder

	col := 0
	for col < lineLen {
		matched := false
		// Try increasing widths until a match is found
		for width := 1; width <= lineLen-col; width++ {
			var glyphLines [8]string
			for r := 0; r < 8; r++ {
				if col+width <= len(rows[r]) {
					glyphLines[r] = rows[r][col : col+width]
				}
			}
			key := strings.Join(glyphLines[:], "\n")
			if r, ok := inv[key]; ok {
				result.WriteRune(r)
				col += width
				matched = true
				break
			}
		}
		if !matched {
			return "", fmt.Errorf("could not decode character at column %d", col)
		}
	}

	return result.String(), nil
}
