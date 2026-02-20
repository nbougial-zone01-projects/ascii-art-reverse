# ASCII Art Generator (Go)

CLI tool that converts a single input string into ASCII art using a banner template.

## Features

- Supports ASCII characters `32..126`.
- Converts literal `\n` in input into real line breaks.
- Preserves case, spacing, and consecutive empty lines.
- Loads banner glyphs from `assets/banners/ascii_library.txt`.
- Includes unit tests and golden integration tests.

## Project Layout

- `cmd/ascii-art/main.go`: entrypoint and app wiring.
- `internal/input`: CLI input validation and `\n` parsing.
- `internal/banner`: banner file loading/parsing.
- `internal/render`: ASCII rendering engine.
- `pkg/model`: shared `Banner` type.
- `test/integration_test.go`: golden regression suite.
- `test/golden/*.txt`: expected outputs (GT-01..GT-10).

## Requirements

- Go `1.21+`

## Run

```bash
go run ./cmd/ascii-art "hello"
```

Or using `make`:

```bash
make run ARGS="hello"
```

To run with the default argument ("Hello"):
```bash
make run
```

Multiline example:

```bash
go run ./cmd/ascii-art "Hello\nThere"
```

Special characters example (bash-safe for `!`):

```bash
go run ./cmd/ascii-art '!@#$%^&*()'
```

## Test

Run all tests:

```bash
go test ./...
```

This includes:
- unit tests for `internal/banner`, `internal/input`, `internal/render`
- golden integration tests in `test/integration_test.go`

## Error Cases

The app exits with code `1` and prints an error when:

- argument count is not exactly 1
- input contains non-ASCII characters (outside `32..126`, excluding newline)
- banner file cannot be read or is malformed/empty

## Notes

- Default banner path is fixed to `assets/banners/ascii_library.txt`.
- If rendering logic changes intentionally, regenerate golden files and rerun `go test ./...`.
- In interactive `bash`, prefer single quotes when input contains `!` to avoid history expansion errors.
