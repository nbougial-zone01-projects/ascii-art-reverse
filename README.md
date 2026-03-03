# ASCII Art Generator (Go)

CLI tool that converts input strings into ASCII art using selectable banner templates.

## Features

- Supports ASCII characters `32..126`.
- Converts literal `\n` in input into real line breaks.
- Preserves case, spacing, and consecutive empty lines.
- Supports banner selection: `standard`, `shadow`, `thinkertoy`.
- Supports output colorization with `--color=<format>` (ANSI name, Hex, RGB, HSL).
- Supports output redirection to file via `--output=<file>`.
- Supports alignment with `--align=left|center|right|justify`.
- Includes unit tests and golden integration tests.

## Project Layout

- `cmd/ascii-art/main.go`: entrypoint and app wiring.
- `internal/input`: CLI input validation and `\n` parsing.
- `internal/banner`: banner file loading/parsing.
- `internal/output`: file output writer.
- `internal/render`: ASCII rendering engine.
- `pkg/model`: shared `Banner` and `Config` types.
- `test/integration_test.go`: golden regression suite.
- `test/golden/*.txt`: expected outputs (GT-01..GT-15 + banner/output cases).

## Requirements

- Go `1.21+`

## Usage

```bash
go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]
```

**Note:** Flags must use the `=` separator (e.g., `--color=red`). Space-separated flags are not supported.

| Flag | Description | Format |
|------|-------------|--------|
| `--color` | Colorize the output | `--color=<color>` |
| `--output` | Save output to a file | `--output=<file>` |
| `--align` | Align the output | `--align=<type>` |

## Run

```bash
go run ./cmd/ascii-art "hello"
```

With banner selection:

```bash
go run ./cmd/ascii-art "hello" shadow
```

With color (whole string):

```bash
go run ./cmd/ascii-art --color=red "hello"
```

With color (substring mode):

```bash
go run ./cmd/ascii-art --color=green "ll" "hello"
```

With alignment:

```bash
go run ./cmd/ascii-art --align=center "hello"
```

With file output:

```bash
go run ./cmd/ascii-art --output=result.txt "hello"
```

Or using `make`:

```bash
make run ARGS="hello"
```

To run with the default argument ("Hello"):
```bash
make run
```

To run with multiple examples:
```bash
make run-examples
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

- input arguments are missing or invalid
- input contains non-ASCII characters (outside `32..126`, excluding newline)
- banner name is unknown (prints available banners) or banner file cannot be read/is malformed
- color format is invalid or unrecognized

## Notes

- Default banner is `standard` (`assets/banners/standard.txt`).
- Terminal width for alignment is read from `COLUMNS` (fallback: `80`).
- If rendering logic changes intentionally, regenerate golden files and rerun `go test ./...`.
- In interactive `bash`, prefer single quotes when input contains `!` to avoid history expansion errors.
