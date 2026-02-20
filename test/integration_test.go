package test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type goldenCase struct {
	name       string
	input      string
	goldenFile string
}

func TestGolden(t *testing.T) {
	testCases := []goldenCase{
		{name: "GT-01", input: "hello", goldenFile: "hello.txt"},
		{name: "GT-02", input: "HELLO", goldenFile: "HELLO.txt"},
		{name: "GT-03", input: "HeLlo WoRlD", goldenFile: "mixed_case.txt"},
		{name: "GT-04", input: "1234567890", goldenFile: "numbers.txt"},
		{name: "GT-05", input: "!@#$%^&*()", goldenFile: "special_chars.txt"},
		{name: "GT-06", input: "Hello\\nThere", goldenFile: "multiline.txt"},
		{name: "GT-07", input: "\\n", goldenFile: "newline_only.txt"},
		{name: "GT-08", input: "", goldenFile: "empty.txt"},
		{name: "GT-09", input: "Hello\\n\\nWorld", goldenFile: "double_newline.txt"},
		{name: "GT-10", input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", goldenFile: "all_upper.txt"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Execute the main program as a subprocess
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tc.input)
			cmd.Dir = ".."
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("command failed: %v, stderr: %s", err, stderr.String())
			}

			// Read the expected output from the golden file
			goldenPath := filepath.Join("golden", tc.goldenFile)
			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden file %q: %v", goldenPath, err)
			}

			// Compare actual stdout with the content of the golden file
			if stdout.String() != string(expected) {
				t.Fatalf("output mismatch\nexpected:\n%s\nactual:\n%s", string(expected), stdout.String())
			}
		})
	}
}
