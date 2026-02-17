# Task 01: Banner Management (Developer 1)

**Objective:** Implement the Data Layer responsible for reading `ascii_library.txt` and parsing it into the `Banner` map.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: File Loading
1.  **RED (Write Test):**
    *   Create `banner_test.go`.
    *   Write `TestLoadBanner_MissingFile`. Call `LoadBanner("nonexistent.txt")` and assert that it returns an error.
2.  **GREEN (Write Code):**
    *   In `banner.go`, implement `LoadBanner` using `os.ReadFile`.
    *   Return an error if the read fails.
3.  **REFACTOR:**
    *   Ensure error messages are clear.

### Cycle 2: Parsing Logic
1.  **RED (Write Test):**
    *   Write `TestLoadBanner_ValidParsing`.
    *   Create a temporary test file (or use a mock string) containing a single character in the standard format (8 lines of art).
    *   Call `LoadBanner` on this file.
    *   Assert that the returned `Banner` map contains the expected key (e.g., ' ') and that the value is a slice of 8 strings.
2.  **GREEN (Write Code):**
    *   Implement the parsing logic in `LoadBanner`.
    *   **Logic:**
        *   Split the file content by newlines (`\n`).
        *   Iterate through the lines.
        *   The file structure is usually: 8 lines of character data, followed by 1 empty line separator.
        *   Map the chunks to ASCII characters starting from 32 (Space) to 126 (~).
3.  **REFACTOR:**
    *   Handle potential Windows line endings (`\r\n`) by sanitizing the input.
    *   Ensure the loop handles the end of the file correctly.

## Deliverables
*   `banner.go` with fully implemented `LoadBanner`.
*   `banner_test.go` with passing tests.