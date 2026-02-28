package render

import "strings"

const Reset = "\033[0m"

// GetColorCode returns the ANSI escape sequence for a supported color name.
// Returns empty string if color is unknown or empty.
func GetColorCode(name string) string {
	switch strings.ToLower(name) {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "orange":
		return "\033[38;5;208m"
	default:
		return ""
	}
}

// ApplyColor wraps the given string s in the colorCode and resets it afterwards.
func ApplyColor(s, colorCode string) string {
	if colorCode == "" {
		return s
	}
	return colorCode + s + Reset
}

// IdentifyColorIndices returns a map of byte indices in the input string that should be colored.
// If sub is empty, all indices are marked as true.
func IdentifyColorIndices(input, sub string) map[int]bool {
	indices := make(map[int]bool)

	// If no substring specified, color everything
	if sub == "" {
		for i := range input {
			indices[i] = true
		}
		return indices
	}

	// Find all non-overlapping occurrences of sub
	start := 0
	for {
		idx := strings.Index(input[start:], sub)
		if idx == -1 {
			break
		}
		actualIdx := start + idx
		for i := 0; i < len(sub); i++ {
			indices[actualIdx+i] = true
		}
		start = actualIdx + len(sub)
	}

	return indices
}