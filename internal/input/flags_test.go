package input

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"ascii-art/pkg/model"
)

func TestParseArgs_Defaults(t *testing.T) {
	// Test backward compatibility: simple input should yield default config
	args := []string{"hello"}
	expected := &model.Config{
		Input:      "hello",
		BannerFile: "standard",
		Align:      "left",
		Color:      "",
		OutputFile: "",
	}

	got, err := ParseArgs(args)
	if err != nil {
		t.Fatalf("ParseArgs() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseArgs() = %+v, want %+v", got, expected)
	}
}

func TestParseArgs_Flags(t *testing.T) {
	// Test parsing of all supported flags
	args := []string{
		"--color=red",
		"--align=right",
		"--output=result.txt",
		"hello",
	}

	expected := &model.Config{
		Input:      "hello",
		BannerFile: "standard",
		Align:      "right",
		Color:      "red",
		OutputFile: "result.txt",
	}

	got, err := ParseArgs(args)
	if err != nil {
		t.Fatalf("ParseArgs() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseArgs() = %+v, want %+v", got, expected)
	}
}

func TestParseArgs_Empty(t *testing.T) {
	_, err := ParseArgs([]string{})
	if err == nil {
		t.Error("ParseArgs(empty) expected error, got nil")
	}
}

func TestParseArgs_Positional(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected *model.Config
		wantErr  bool
	}{
		{
			name: "String and Banner",
			args: []string{"hello", "shadow"},
			expected: &model.Config{
				Input:      "hello",
				BannerFile: "shadow",
				Align:      "left",
				Color:      "",
				OutputFile: "",
			},
		},
		{
			name: "Color Substring and String",
			args: []string{"--color=red", "he", "hello"},
			expected: &model.Config{
				Input:       "hello",
				BannerFile:  "standard",
				Align:       "left",
				Color:       "red",
				ColorSubstr: "he",
				OutputFile:  "",
			},
		},
		{
			name: "Color Input and Banner",
			args: []string{"--color=red", "hello", "shadow"},
			expected: &model.Config{
				Input:       "hello",
				BannerFile:  "shadow",
				Align:       "left",
				Color:       "red",
				ColorSubstr: "",
				OutputFile:  "",
			},
		},
		{
			name: "Color Substring, String, and Banner",
			args: []string{"--color=green", "H", "Hello", "thinkertoy"},
			expected: &model.Config{
				Input:       "Hello",
				BannerFile:  "thinkertoy",
				Align:       "left",
				Color:       "green",
				ColorSubstr: "H",
				OutputFile:  "",
			},
		},
		{
			name:    "Invalid Banner",
			args:    []string{"hello", "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseArgs() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}

func TestParseArgs_ReverseFlag(t *testing.T) {
	// Create a real temp file so the existence check passes
	dir := t.TempDir()
	path := filepath.Join(dir, "art.txt")
	if err := os.WriteFile(path, []byte("art"), 0o644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	got, err := ParseArgs([]string{"--reverse=" + path})
	if err != nil {
		t.Fatalf("ParseArgs() unexpected error: %v", err)
	}
	if got.ReverseFile != path {
		t.Errorf("ReverseFile = %q, want %q", got.ReverseFile, path)
	}
}

func TestParseArgs_ReverseFlagMissingFile(t *testing.T) {
	_, err := ParseArgs([]string{"--reverse=nonexistent.txt"})
	if err == nil {
		t.Fatal("ParseArgs() expected error for missing file, got nil")
	}
	if !contains(err.Error(), "File not found") {
		t.Errorf("error = %q, want it to contain %q", err.Error(), "File not found")
	}
}

func TestParseArgs_ReverseFlagBadFormat(t *testing.T) {
	_, err := ParseArgs([]string{"--reverse"})
	if err == nil {
		t.Fatal("ParseArgs() expected error for bad flag format, got nil")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}