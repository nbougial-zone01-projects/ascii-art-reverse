# Golden File Testing Strategy

## Objective
To ensure regression testing by comparing the program's actual output against pre-verified "golden" output files. This guarantees that refactoring or new features do not break existing functionality.

## Test Suite

The following test cases define the input and the expected behavior.

| Case ID | Input String | Description | Expected Output File |
| :--- | :--- | :--- | :--- |
| **GT-01** | `"hello"` | Basic lowercase word | `test/golden/hello.txt` |
| **GT-02** | `"HELLO"` | Basic uppercase word | `test/golden/HELLO.txt` |
| **GT-03** | `"HeLlo WoRlD"` | Mixed case with spaces | `test/golden/mixed_case.txt` |
| **GT-04** | `"1234567890"` | Numbers | `test/golden/numbers.txt` |
| **GT-05** | `"!@#$%^&*()"` | Special characters | `test/golden/special_chars.txt` |
| **GT-06** | `"Hello\nThere"` | Embedded newline | `test/golden/multiline.txt` |
| **GT-07** | `"\n"` | Single newline (should print empty line) | `test/golden/newline_only.txt` |
| **GT-08** | `""` | Empty string (should be handled gracefully) | `test/golden/empty.txt` |
| **GT-09** | `"Hello\n\nWorld"` | Multiple consecutive newlines | `test/golden/double_newline.txt` |
| **GT-10** | `"ABCDEFGHIJKLMNOPQRSTUVWXYZ"` | Full uppercase alphabet | `test/golden/all_upper.txt` |

## Implementation Plan

1.  **Generate Golden Files:**
    *   Run a verified version of the code.
    *   Redirect output: `go run . "hello" > test/golden/hello.txt`.
    *   Manually inspect `hello.txt` to confirm it is correct.

2.  **Automated Test Runner (`integration_test.go`):**
    *   Iterate through the table of test cases.
    *   For each case:
        1.  Capture `stdout` of the `main()` function or `Render()` function.
        2.  Read the content of the corresponding `.txt` file.
        3.  Compare `actual_output` vs `expected_output`.
        4.  Fail the test if they differ.

## Maintenance
*   If the rendering logic changes intentionally (e.g., fixing a bug in the font), regenerate the golden files and verify them manually.