# Task 07: Feature Output (Developer 3)

**Objective:** Implement file writing capability in the main application controller.

## Steps

1.  **Update Main Logic (`cmd/ascii-art/main.go`)**
    *   Update `Run` to use the new `input.ParseArgs` (Task 05).
    *   Check `Config.OutputFile`.
    *   **Logic:**
        *   If `OutputFile` is empty, write to `os.Stdout` (existing behavior).
        *   If `OutputFile` is set, create/truncate the file and write the result there.

2.  **Integration Test Update**
    *   Add a test case in `test/integration_test.go` that runs with `--output` and verifies the file creation and content.

## Acceptance Criteria
*   [ ] Flag `--output=<file>` redirects output from stdout to file.
*   [ ] File is created if it doesn't exist.
*   [ ] File is overwritten if it exists.
*   [ ] Permissions are handled gracefully (return error if write fails).

## Note
*   This task depends on Task 05 being complete (or the Config struct being available).