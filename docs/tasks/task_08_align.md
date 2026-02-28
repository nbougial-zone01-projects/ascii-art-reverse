# Task 08: Feature Justify/Align (Developer 4)

**Objective:** Implement text alignment (left, center, right, justify) based on terminal size.

## Steps

1.  **Terminal Size Detection**
    *   Use `golang.org/x/term` or `syscall` to get terminal width.
    *   Fallback to a default (e.g., 80 columns) if detection fails or for testing.

2.  **Implement Alignment Logic (`internal/render/align.go`)**
    *   Calculate the width of the ASCII art block (8 lines).
    *   **Left (Default):** `padding = 0`. Return lines as is (Backward Compatibility).
    *   **Center:** `padding = (termWidth - artWidth) / 2`.
    *   **Right:** `padding = termWidth - artWidth`.
    *   **Justify:** Distribute spaces between words.
    *   Apply padding (spaces) to the left of every line in the 8-line block.

3.  **Unit Tests (`internal/render/align_test.go`)**
    *   Create `internal/render/align_test.go`.
    *   Mock the terminal width in tests to ensure deterministic results.
    *   Verify padding calculation for Center and Right.
    *   Verify that `align=left` adds zero padding.

## Acceptance Criteria
*   [ ] `--align=center` centers the art.
*   [ ] `--align=right` aligns to right edge.
*   [ ] `--align=justify` spreads words across the line.
*   [ ] Adapts to window size (if running in real terminal).

## Note
*   Coordinate with Developer 2 (Color) as you both modify `renderer.go`.