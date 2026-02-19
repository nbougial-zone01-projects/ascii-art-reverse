package banner

import (
	"ascii-art/pkg/model"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	asciiStart  = 32
	asciiEnd    = 126
	glyphHeight = 8
)

var (
	ErrEmptyBanner     = errors.New("banner file is empty")
	ErrMalformedBanner = errors.New("banner file format is invalid")
)

// LoadBanner reads a banner file from disk and parses it into a rune->glyph map.
func LoadBanner(filename string) (model.Banner, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read banner file %q: %w", filename, err)
	}
	if len(data) == 0 {
		return nil, ErrEmptyBanner
	}

	normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	lines := strings.Split(normalized, "\n")

	b := make(model.Banner, asciiEnd-asciiStart+1)
	index := 0
	if len(lines) > 0 && lines[0] == "" {
		index = 1
	}

	for ascii := asciiStart; ascii <= asciiEnd; ascii++ {
		if index+glyphHeight > len(lines) {
			return nil, ErrMalformedBanner
		}

		glyph := make([]string, glyphHeight)
		copy(glyph, lines[index:index+glyphHeight])
		b[rune(ascii)] = glyph
		index += glyphHeight

		if index < len(lines) && lines[index] == "" {
			index++
		}
	}

	for ; index < len(lines); index++ {
		if lines[index] != "" {
			return nil, ErrMalformedBanner
		}
	}

	return b, nil
}
