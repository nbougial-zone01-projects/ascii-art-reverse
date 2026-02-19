package render

import (
	"ascii-art/pkg/model"
	"testing"
)

func TestRender_Simple(t *testing.T) {
	// 1. Mock the Banner (Data Layer)
	// We manually construct a banner with 1 character ('A') and 8 dummy lines.
	mockBanner := model.Banner{
		'A': []string{"Line1", "Line2", "Line3", "Line4", "Line5", "Line6", "Line7", "Line8"},
	}

	// 2. Define Input and Expected Output
	input := "A"
	expected := "Line1\nLine2\nLine3\nLine4\nLine5\nLine6\nLine7\nLine8"

	// 3. Call the function (This will fail to compile initially because Render doesn't exist)
	got := Render(input, mockBanner)

	if got != expected {
		t.Errorf("Render() = %q, want %q", got, expected)
	}
}