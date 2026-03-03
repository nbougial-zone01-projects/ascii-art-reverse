package render

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

const defaultTerminalWidth = 80

var terminalWidthProvider = getTerminalWidth

func getTerminalWidth() int {
	// Try to get the width using 'tput', which is more reliable than COLUMNS in some shells.
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		if n, err := strconv.Atoi(strings.TrimSpace(string(out))); err == nil && n > 0 {
			return n
		}
	}

	if cols := os.Getenv("COLUMNS"); cols != "" {
		n, err := strconv.Atoi(cols)
		if err == nil && n > 0 {
			return n
		}
	}
	return defaultTerminalWidth
}

func calculatePadding(align string, artWidth, terminalWidth int) int {
	if terminalWidth <= artWidth {
		return 0
	}

	switch strings.ToLower(align) {
	case "center":
		return (terminalWidth - artWidth) / 2
	case "right":
		return terminalWidth - artWidth
	default:
		return 0
	}
}

func applyAlign(lines []string, align string, terminalWidth int) []string {
	if len(lines) == 0 {
		return lines
	}

	mode := strings.ToLower(align)
	if mode == "" || mode == "left" {
		return lines
	}

	out := make([]string, len(lines))
	copy(out, lines)

	if mode == "justify" {
		for i := range out {
			out[i] = justifyLine(out[i], terminalWidth)
		}
		return out
	}

	maxWidth := 0
	for _, line := range out {
		w := visibleWidth(line)
		if w > maxWidth {
			maxWidth = w
		}
	}

	padding := calculatePadding(mode, maxWidth, terminalWidth)
	if padding <= 0 {
		return out
	}

	pad := strings.Repeat(" ", padding)
	for i := range out {
		out[i] = pad + out[i]
	}
	return out
}

func justifyLine(line string, terminalWidth int) string {
	width := visibleWidth(line)
	if width >= terminalWidth {
		return line
	}

	type gap struct {
		start int
		end   int
	}
	var gaps []gap
	for i := 0; i < len(line); {
		if unicode.IsSpace(rune(line[i])) {
			start := i
			for i < len(line) && unicode.IsSpace(rune(line[i])) {
				i++
			}
			// Ignore leading/trailing spaces, spread only internal gaps.
			if start > 0 && i < len(line) {
				gaps = append(gaps, gap{start: start, end: i})
			}
			continue
		}
		i++
	}
	if len(gaps) == 0 {
		return line
	}

	extra := terminalWidth - width
	base := extra / len(gaps)
	rem := extra % len(gaps)

	var sb strings.Builder
	last := 0
	for i, g := range gaps {
		sb.WriteString(line[last:g.end])
		add := base
		if i < rem {
			add++
		}
		sb.WriteString(strings.Repeat(" ", add))
		last = g.end
	}
	sb.WriteString(line[last:])
	return sb.String()
}

func visibleWidth(s string) int {
	// Remove ANSI CSI escape sequences to avoid skewed width calculations.
	var cleaned strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == 0x1b && i+1 < len(s) && s[i+1] == '[' {
			i += 2
			for i < len(s) {
				c := s[i]
				if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
					break
				}
				i++
			}
			continue
		}
		cleaned.WriteByte(s[i])
	}
	return len(cleaned.String())
}
