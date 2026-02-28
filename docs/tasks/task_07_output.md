# Task 07: Feature Output (Developer 3)

**Objective:** Implement file writing capability in the main application controller.

## Steps

1.  **Implement File Writer (`cmd/ascii-art/file_writer.go`)**
    *   Create `cmd/ascii-art/file_writer.go`.
    *   Implement a function `WriteOutput(filename string, content string) error`.
    *   Logic: Create/Truncate file and write content.

2.  **Unit Tests (`cmd/ascii-art/file_writer_test.go`)**
    *   Create `cmd/ascii-art/file_writer_test.go`.
    *   Test file creation and content writing (use temporary files).

3.  **Update Main Logic (`cmd/ascii-art/main.go`)**
    *   Update `Run` to use the new `input.ParseArgs` (Task 05).
    *   If `Config.OutputFile` is set, call `WriteOutput`.
    *   If empty, write to `os.Stdout`.
    *   Add a test case in `test/integration_test.go` that runs with `--output` and verifies the file creation and content.

## Acceptance Criteria
*   [ ] Flag `--output=<file>` redirects output from stdout to file.
*   [ ] File is created if it doesn't exist.
*   [ ] File is overwritten if it exists.
*   [ ] Permissions are handled gracefully (return error if write fails).

## Note
*   This task depends on Task 05 being complete (or the Config struct being available).