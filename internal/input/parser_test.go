package input

import (
	"errors"
	"testing"
)

func TestParseInput_NoArgs(t *testing.T) {
	_, err := ParseInput([]string{})
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected ErrUsage, got %v", err)
	}
}

func TestParseInput_TooManyArgs(t *testing.T) {
	_, err := ParseInput([]string{"one", "two"})
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected ErrUsage, got %v", err)
	}
}

func TestParseInput_ValidString(t *testing.T) {
	got, err := ParseInput([]string{"Hello 123 !"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "Hello 123 !" {
		t.Fatalf("ParseInput() = %q, want %q", got, "Hello 123 !")
	}
}

func TestParseInput_EscapedNewline(t *testing.T) {
	got, err := ParseInput([]string{"Hello\\nWorld"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "Hello\nWorld" {
		t.Fatalf("ParseInput() = %q, want %q", got, "Hello\nWorld")
	}
}

func TestParseInput_InvalidChar(t *testing.T) {
	_, err := ParseInput([]string{"Hell€"})
	if !errors.Is(err, ErrInvalidASCII) {
		t.Fatalf("expected ErrInvalidASCII, got %v", err)
	}
}

func TestParseInput_EmptyString(t *testing.T) {
	got, err := ParseInput([]string{""})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "" {
		t.Fatalf("ParseInput() = %q, want empty string", got)
	}
}
