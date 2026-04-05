package reverse

import (
	"ascii-art/pkg/model"
	"testing"
)

// buildArt constructs the expected ASCII art string for a given string
// using the provided mock banner, mirroring how the renderer works.
func buildArt(input string, b model.Banner) string {
	var rows [8]string
	for _, ch := range input {
		glyph := b[ch]
		for i := 0; i < 8; i++ {
			rows[i] += glyph[i]
		}
	}
	result := ""
	for i := 0; i < 8; i++ {
		if i > 0 {
			result += "\n"
		}
		result += rows[i]
	}
	return result
}

func mockBanner() model.Banner {
	return model.Banner{
		'h': []string{"h1", "h2", "h3", "h4", "h5", "h6", "h7", "h8"},
		'e': []string{"e1", "e2", "e3", "e4", "e5", "e6", "e7", "e8"},
		'l': []string{"l1", "l2", "l3", "l4", "l5", "l6", "l7", "l8"},
		'o': []string{"o1", "o2", "o3", "o4", "o5", "o6", "o7", "o8"},
	}
}

func TestReverse_Simple(t *testing.T) {
	b := mockBanner()
	art := buildArt("hello", b)

	got, err := Reverse(art, b)
	if err != nil {
		t.Fatalf("Reverse() unexpected error: %v", err)
	}
	if got != "hello" {
		t.Errorf("Reverse() = %q, want %q", got, "hello")
	}
}

func TestReverse_Empty(t *testing.T) {
	b := mockBanner()

	got, err := Reverse("", b)
	if err != nil {
		t.Fatalf("Reverse() unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("Reverse() = %q, want %q", got, "")
	}
}

func TestReverse_Multiline(t *testing.T) {
	b := mockBanner()
	// Build art for "he" on first line and "lo" on second line
	line1 := buildArt("he", b)
	line2 := buildArt("lo", b)
	art := line1 + "\n" + line2

	got, err := Reverse(art, b)
	if err != nil {
		t.Fatalf("Reverse() unexpected error: %v", err)
	}
	if got != "he\nlo" {
		t.Errorf("Reverse() = %q, want %q", got, "he\nlo")
	}
}
