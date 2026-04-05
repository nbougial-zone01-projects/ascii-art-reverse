# 📄 PRD – ASCII Art Generator (Go)

## 1. Overview

The ASCII Art Generator is a feature-rich command-line application written in Go. It converts input strings into graphical ASCII-art using predefined banner files, with support for text colorization, output alignment, and file export.

The program must render letters, numbers, spaces, special characters, and newline sequences (`\n`) according to the format defined in the provided banner files.

The program also supports the **reverse** operation: given a file containing ASCII art produced by this tool, it reconstructs and prints the original input string.

---

## 2. Objectives

- Convert input text into ASCII art format.
- Support banner templates:
  - `standard`
  - `shadow`
  - `thinkertoy`
- **Colorize output** based on user input (whole string or substrings).
- **Export output** to a file.
- **Align text** (left, center, right, justify) relative to terminal size.
- **Reverse ASCII art** back to the original text string from a file.
- Ensure output strictly matches expected formatting examples.
- Maintain clean, modular, and testable Go code.

---

## 3. Scope

### In Scope

- Parsing CLI input.
- Parsing flags (`--color`, `--output`, `--align`, `--reverse`).
- Interpreting literal `\n` as line breaks.
- Loading banner files from the filesystem.
- Rendering ASCII characters (ASCII 32–126).
- Applying ANSI color codes to output.
- Writing output to files.
- Detecting terminal window size for alignment.
- Supporting:
  - Uppercase letters
  - Lowercase letters
  - Numbers
  - Spaces
  - Special characters
- Handling multiple consecutive newline sequences.
- Unit testing core logic.

### Out of Scope

- Generating ASCII art algorithmically.
- Editing or modifying banner files.
- Supporting characters outside ASCII 32–126.
- GUI or web interface.
- Reversing ASCII art that was produced with color or alignment flags.

---

## 4. Functional Requirements

### 4.1 General Input Handling

- The program must accept flags and positional arguments.
- The input may contain:
  - Letters
  - Numbers
  - Spaces
  - Special characters
  - Literal `\n` sequences
- The program must:
  - Interpret `\n` as a new line
  - Handle multiple consecutive `\n`
  - Handle empty string input

### 4.2 Color Feature

- **Flag:** `--color=<color>`
- **Syntax:** `go run . --color=<color> <substring> "string"`
- **Behavior:**
  - Colors the specified `<substring>` within the string.
  - If `<substring>` is not provided (or matches the main string), the whole string is colored.
  - Supports standard color names (red, blue, etc.), Hex (#RRGGBB), RGB (rgb(r,g,b)), and HSL (hsl(h,s,l)).
  - **Ambiguity Handling:** When 2 arguments are provided (`go run . --color=red arg1 arg2`), the program checks if `arg2` is a valid banner.
    - If valid banner: `Input=arg1`, `Banner=arg2`.
    - If not: `ColorSubstr=arg1`, `Input=arg2`.
 - **Usage Error:** If the flag format is incorrect or color is invalid:
  ```text
  Usage: go run . --color=<color> <substring> "string"

  Supported formats: ANSI standard colors (red, green, blue...), Hex (#RRGGBB), RGB (rgb(r,g,b)), HSL (hsl(h,s,l))
  ```

### 4.3 Output File Feature

- **Flag:** `--output=<fileName.txt>`
- **Syntax:** `go run . --output=<fileName.txt> [STRING] [BANNER]`
- **Behavior:**
  - Writes the resulting ASCII art to `<fileName.txt>` instead of stdout.
  - Supports optional `[BANNER]` argument.
- **Usage Error:** If the flag format is incorrect:
  ```text
  Usage: go run . [OPTION] [STRING] [BANNER]
  
  EX: go run . --output=<fileName.txt> something standard
  ```

### 4.4 Alignment Feature

- **Flag:** `--align=<type>`
- **Syntax:** `go run . --align=<type> [STRING] [BANNER]`
- **Types:**
  - `left` (default behavior)
  - `center`
  - `right`
  - `justify`
- **Behavior:**
  - Adapts the graphical representation to the current terminal size.
  - Only text that fits the terminal size will be tested.
- **Usage Error:** If the flag format is incorrect:
  ```text
  Usage: go run . [OPTION] [STRING] [BANNER]
  
  Example: go run . --align=right something standard
  ```

### 4.5 Reverse Feature

- **Flag:** `--reverse=<fileName>`
- **Syntax:** `go run . --reverse=<fileName>`
- **Behavior:**
  - Reads the specified file containing ASCII art produced by this tool.
  - Reconstructs and prints the original input string to stdout.
  - Uses the `standard` banner by default for glyph matching. An optional `[BANNER]` argument may be provided to specify which banner was used to produce the art.
  - Each line of ASCII art (8 rows) is matched against the banner glyph map to identify the original character.
  - Empty 8-row blocks (all blank lines) are treated as spaces.
  - Multiple lines of ASCII art separated by a blank line are reconstructed with `\n` between them.
- **File Not Found Error:** If the specified file does not exist:
  ```text
  Usage: go run . [OPTION]

  EX: go run . --reverse=<fileName>

  File not found: <fileName>
  ```
- **Usage Error:** If the flag format is incorrect:
  ```text
  Usage: go run . [OPTION]

  EX: go run . --reverse=<fileName>
  ```

### 4.6 Banner Selection

Banner files are preformatted ASCII templates.
The user can specify a banner as the last argument.

- **Syntax:** `go run . [STRING] [BANNER]`
- **Defaults:** If not specified, use `standard`.
- **Usage Error:** If the format is incorrect:
  ```text
  Usage: go run . [STRING] [BANNER]
  
  EX: go run . something standard
  ```

---

### 4.7 Rendering Rules

Characters must be rendered horizontally.

Each output block must contain exactly 8 lines per input line.

Rendering must preserve:

- Case sensitivity
- Spacing

Empty input lines must result in blank output lines.

Output must match provided examples exactly.

---

### 4.8 Error Handling

The program must handle:

- Missing banner file.
- Invalid ASCII characters (outside 32–126).
- Incorrect CLI usage.
- Missing or non-existent file for `--reverse`.

Errors must:

- Not crash the program unexpectedly.
- Provide clear, readable error messages.
- Return the specific usage messages defined in sections 4.2, 4.3, 4.4, 4.5, and 4.6 depending on the context.

### 5. Non-Functional Requirements

- Must use only standard Go packages.
- Code must follow Go best practices.
- Must be modular and maintainable.
- Must include unit tests.
- Must avoid hardcoded ASCII representations inside source code.

---

### 6. Constraints

- Banner height is fixed at 8 lines.
- Characters are separated by newline in banner files.
- ASCII range supported: 32–126.
- Banner files are read-only.
- Implementation must be written in Go.

---

### 7. Acceptance Criteria

The project is considered complete when:

- Output matches all provided examples exactly.
- Multiple newline inputs behave correctly.
- Case sensitivity is preserved.
- **Color flag works for substrings and full strings.**
- **Output flag correctly writes to files.**
- **Alignment flag correctly positions text based on terminal width.**
- **Banner argument correctly switches fonts.**
- **Reverse flag correctly reconstructs the original string from an ASCII art file.**
- **Reverse flag returns a usage message with file-not-found info when the file does not exist.**
- Unit tests pass.
- Code is clean and modular.
- No banner data is hardcoded.

---

### 8. Success Metrics

- Zero formatting mismatches in expected output.
- Full support for ASCII 32–126.
- Clear separation of responsibilities in code structure.
- All team members can understand and extend the codebase.

---

### 9. Risks & Mitigation

| Risk | Mitigation |
|------|------------|
| Incorrect banner indexing | Write unit tests for specific characters |
| Improper newline handling | Add dedicated newline test cases |
| Merge conflicts between team members | Define API contracts early |
| Output formatting mismatch | Use golden file tests |
| Terminal size detection failure | Fallback to default width (e.g., 80 chars) |
| Ambiguous glyph matching in reverse | Use exact 8-line block comparison against banner map |

---

### 10. Future Improvements (Optional)

- Add performance optimizations.
- Add additional ASCII fonts.
- Add integration tests.

### 11. Misc.
When Unsure ask for confirmation
Use only standard packages
Do not hallunicate
Stay strict to what is mentioned in the documentation
Do not assume things that are not stated
In every new update made in the repo update ai.log
When updating ai.log keep the format of the file untouched
