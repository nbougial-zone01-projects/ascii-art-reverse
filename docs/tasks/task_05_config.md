# Task 05: Input Configuration (Developer 1)

**Objective:** Update the shared model and implement robust flag parsing to support the new features.

## Steps

1.  **Define Config Model (`pkg/model/config.go`)**
    *   Create a struct `Config` containing:
        *   `Input` (string)
        *   `BannerFile` (string, default "standard")
        *   `Color` (string)
        *   `ColorSubstr` (string)
        *   `OutputFile` (string)
        *   `Align` (string, default "left")

2.  **Implement Flag Parser (`internal/input/parser.go`)**
    *   Refactor `ParseInput` to `ParseArgs(args []string) (*model.Config, error)`.
    *   Logic:
        *   Loop through args to find flags starting with `--`.
        *   Extract values for `--color`, `--output`, `--align`.
        *   Identify positional arguments:
            *   If 1 arg remaining: `[STRING]` (Banner = standard).
            *   If 2 args remaining: `[STRING] [BANNER]`.
            *   Handle special case for Color: `[SUBSTRING] [STRING]`.
    *   **Validation:** Ensure flags follow the exact format defined in PRD. Return specific usage errors if malformed.

3.  **Unit Tests (`internal/input/parser_test.go`)**
    *   Test flag extraction (`--output=test.txt`).
    *   Test positional arg logic (1 vs 2 args).
    *   Test the specific usage error messages.

## Acceptance Criteria
*   [ ] `pkg/model/config.go` exists.
*   [ ] `ParseArgs` returns a populated `Config` object.
*   [ ] Correctly identifies `[STRING]` vs `[BANNER]`.
*   [ ] Correctly parses `--color`, `--output`, `--align`.
*   [ ] Returns specific error messages for bad flags as per PRD.