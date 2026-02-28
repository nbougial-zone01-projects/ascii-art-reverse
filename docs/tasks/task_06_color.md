# Task 06: Feature Color (Developer 2)

**Objective:** Update the rendering engine to support ANSI color codes.

## Steps

1.  **Update Render Signature**
    *   Update `Render` to accept `(*model.Config, model.Banner)`.

2.  **Implement Color Logic (`internal/render/color.go` or inside `renderer.go`)**
    *   Create a helper to map color names (red, blue) to ANSI codes (e.g., `\033[31m`).
    *   Logic:
        *   If `Config.Color` is empty, render normally.
        *   If `Config.Color` is set:
            *   Identify indices of characters to color (Whole string vs Substring).
            *   When building the line, prepend ANSI code before the character and `\033[0m` (reset) after it.

3.  **Unit Tests**
    *   Test that specific substrings are wrapped in ANSI codes.
    *   Test that the rest of the string remains untouched.

## Acceptance Criteria
*   [ ] Supports standard colors (red, green, blue, etc.).
*   [ ] Colors the whole string if no substring is provided.
*   [ ] Colors only the substring if provided.
*   [ ] Output contains correct ANSI escape sequences.

## Note
*   Coordinate with Developer 4 (Align) as you both modify `renderer.go`.