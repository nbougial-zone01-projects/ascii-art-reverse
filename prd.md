# 📄 PRD – ASCII Art Generator (Go)

## 1. Overview

The ASCII Art Generator is a command-line application written in Go that receives a string as input and outputs a graphical ASCII-art representation of that string using predefined banner files.

The program must render letters, numbers, spaces, special characters, and newline sequences (`\n`) according to the format defined in the provided banner files.

---

## 2. Objectives

- Convert input text into ASCII art format.
- Support banner templates:
  - `standard`
  - `shadow`
  - `thinkertoy`
- Ensure output strictly matches expected formatting examples.
- Maintain clean, modular, and testable Go code.

---

## 3. Scope

### In Scope

- Parsing CLI input.
- Interpreting literal `\n` as line breaks.
- Loading banner files from the filesystem.
- Rendering ASCII characters (ASCII 32–126).
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

---

## 4. Functional Requirements

### 4.1 Input Handling

- The program must accept exactly one string argument.
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

Example:

```bash
go run . "Hello\nThere"
```

### 4.2 Banner Handling

Banner files are preformatted ASCII templates.

Each character:

- Has a height of 8 lines.
- Is separated by a newline.
- Characters are ordered by ASCII value (starting at ASCII 32: space).

Banner files must not be modified.

The system must load banner files dynamically from disk.

---

### 4.3 Rendering Rules

Characters must be rendered horizontally.

Each output block must contain exactly 8 lines per input line.

Rendering must preserve:

- Case sensitivity
- Spacing

Empty input lines must result in blank output lines.

Output must match provided examples exactly.

---

### 4.4 Error Handling

The program must handle:

- Missing banner file.
- Invalid ASCII characters (outside 32–126).
- Incorrect CLI usage.

Errors must:

- Not crash the program unexpectedly.
- Provide clear, readable error messages.

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

---

### 10. Future Improvements (Optional)

- Support banner selection via CLI argument.
- Add performance optimizations.
- Add additional ASCII fonts.
- Add integration tests.

