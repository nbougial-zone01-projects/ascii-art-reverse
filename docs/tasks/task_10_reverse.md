# Task 10: Reverse Feature

**Objective:** Implement the `--reverse=<fileName>` flag that reads a file containing ASCII art produced by this tool and reconstructs the original input string, printing it to stdout.

---

## Background

The reverse process is the inverse of rendering:
- The renderer maps each input rune → 8 lines of ASCII art.
- The reverser maps each 8-line block back → the original rune.

The file is read, split into 8-row blocks (one per input character column), each block is looked up in the banner glyph map, and the matched rune is appended to the result string.

---

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Core Reverse Logic (`internal/reverse`)

1. **RED (Write Test):**
   - Create `internal/reverse/reverse_test.go`.
   - Write `TestReverse_Simple`:
     - Build a mock `model.Banner` with known glyphs for `'h'`, `'e'`, `'l'`, `'o'`.
     - Construct the expected 8-line art string for `"hello"` from the mock banner.
     - Call `Reverse(artString, banner)` and assert result == `"hello"`.
   - Write `TestReverse_Empty`:
     - Pass an empty string, assert result == `""`.
   - Run `go test ./internal/reverse/...` — confirm **FAIL** (package does not exist yet).

2. **GREEN (Write Code):**
   - Create `internal/reverse/reverse.go`.
   - Implement `Reverse(art string, b model.Banner) (string, error)`:
     - Split `art` by `"\n"` into rows.
     - Group rows into 8-line blocks (one block per input line).
     - For each block, slice column-by-column to extract per-character glyph slices.
     - Look up each glyph slice in the banner map (inverted lookup).
     - Append matched rune to result; return error if no match found.
   - Run `go test ./internal/reverse/...` — confirm **PASS**.

3. **REFACTOR:**
   - Extract the inverted banner map build into a helper `invertBanner(b model.Banner) map[string]rune`.
   - Ensure no logic duplication.
   - Run `go test ./internal/reverse/...` — confirm **PASS**.

---

### Cycle 2: Multiline Reverse

1. **RED (Write Test):**
   - Add `TestReverse_Multiline` to `internal/reverse/reverse_test.go`.
   - Input: art string for `"Hello\nThere"` built from mock banner.
   - Assert result == `"Hello\nThere"`.
   - Run `go test ./internal/reverse/...` — confirm **FAIL**.

2. **GREEN (Write Code):**
   - Update `Reverse` to handle multiple 8-line blocks separated by a blank line.
   - Each blank-line separator between blocks maps to a `\n` in the output.
   - Run `go test ./internal/reverse/...` — confirm **PASS**.

---

### Cycle 3: Flag Parsing (`--reverse`)

1. **RED (Write Test):**
   - Add to `internal/input/flags_test.go`:
     - `TestParseArgs_ReverseFlag`: assert `--reverse=file.txt` sets `Config.ReverseFile = "file.txt"` (use a real temp file so the existence check passes).
     - `TestParseArgs_ReverseFlagMissingFile`: assert `--reverse=nonexistent.txt` returns an error containing `"File not found"`.
     - `TestParseArgs_ReverseFlagBadFormat`: assert `--reverse` (no `=`) returns a usage error.
   - Run `go test ./internal/input/...` — confirm **FAIL**.

2. **GREEN (Write Code):**
   - Add `ReverseFile string` to `pkg/model/config.go`.
   - Add `--reverse` handling to `internal/input/flags.go`:
     - Parse with `strings.SplitN(arg, "=", 2)`.
     - If no `=` or empty value: return usage error.
     - Use `os.Stat` to check file existence; if not found: return file-not-found error.
     - Set `config.ReverseFile = parts[1]`.
     - In reverse mode, skip positional argument validation (no `[STRING]` required).
   - Run `go test ./internal/input/...` — confirm **PASS**.

---

### Cycle 4: Wiring (`cmd/ascii-art/main.go`)

1. **RED (Write Test):**
   - Add `TestGoldenReverse` to `test/integration_test.go`:
     - GT-19: `--reverse=test/golden/hello.txt` → stdout == `"hello\n"`.
     - GT-20: `--reverse=test/golden/multiline.txt` → stdout == `"Hello\nThere\n"`.
     - GT-21: `--reverse=test/golden/shadow_hello.txt shadow` → stdout == `"hello\n"`.
     - GT-22: `--reverse=nonexistent.txt` → exit code != 0, stderr contains `"File not found"`.
   - Run `go test ./test/...` — confirm **FAIL**.

2. **GREEN (Write Code):**
   - Update `Run` in `cmd/ascii-art/main.go`:
     - After `ParseArgs`, check `cfg.ReverseFile != ""`.
     - If set: load the banner, read the file, call `reverse.Reverse`, print result to stdout.
     - Otherwise: existing render path unchanged.
   - Run `go test ./test/...` — confirm **PASS**.
   - Run `go test ./...` — confirm full suite **PASS**.

---

## File Checklist

| File | Action |
|------|--------|
| `pkg/model/config.go` | Add `ReverseFile string` field |
| `internal/input/flags.go` | Add `--reverse` flag parsing with file-existence check |
| `internal/input/flags_test.go` | Add 3 new test cases for `--reverse` |
| `internal/reverse/reverse.go` | New — core reverse logic |
| `internal/reverse/reverse_test.go` | New — unit tests for reverse logic |
| `cmd/ascii-art/main.go` | Wire reverse path into `Run` |
| `test/integration_test.go` | Add `TestGoldenReverse` (GT-19..GT-22) |
| `docs/logs/ai.log` | Updated after each change |

---

## Acceptance Criteria

- [ ] `--reverse=<fileName>` reconstructs the original string from standard banner art.
- [ ] `--reverse=<fileName> <banner>` reconstructs using the specified banner.
- [ ] Multiline art (blocks separated by blank line) reconstructs with `\n` between words.
- [ ] If the file does not exist, a usage message with `"File not found: <fileName>"` is printed and exit code is 1.
- [ ] If the flag format is incorrect (no `=`), a usage message is printed and exit code is 1.
- [ ] All existing tests continue to pass (no regression).
- [ ] Unit tests cover: simple reverse, empty input, multiline, flag parsing.
