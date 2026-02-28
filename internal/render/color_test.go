package render

import "testing"

func TestGetColorCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Red", "red", "\033[31m"},
		{"Green", "green", "\033[32m"},
		{"Blue", "blue", "\033[34m"},
		{"Orange", "orange", "\033[38;5;208m"},
		{"Empty", "", ""},
		{"Unknown", "mystic", ""},
		{"CaseInsensitive", "ReD", "\033[31m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetColorCode(tt.input)
			if got != tt.expected {
				t.Errorf("GetColorCode(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestApplyColor(t *testing.T) {
	colorCode := "\033[31m"
	input := "A"
	expected := "\033[31mA\033[0m"

	got := ApplyColor(input, colorCode)
	if got != expected {
		t.Errorf("ApplyColor(%q, %q) = %q, want %q", input, colorCode, got, expected)
	}
}

func TestIdentifyColorIndices(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sub      string
		expected map[int]bool // Key is byte index
	}{
		{
			name:     "WholeString",
			input:    "abc",
			sub:      "",
			expected: map[int]bool{0: true, 1: true, 2: true},
		},
		{
			name:     "SubstringMiddle",
			input:    "hello",
			sub:      "ll",
			expected: map[int]bool{2: true, 3: true},
		},
		{
			name:     "MultipleOccurrences",
			input:    "banana",
			sub:      "a",
			expected: map[int]bool{1: true, 3: true, 5: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IdentifyColorIndices(tt.input, tt.sub)
			for idx, want := range tt.expected {
				if !got[idx] && want {
					t.Errorf("Index %d expected true, got false", idx)
				}
			}
			if len(got) != len(tt.expected) {
				t.Errorf("Expected %d indices, got %d", len(tt.expected), len(got))
			}
		})
	}
}