package input

import (
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