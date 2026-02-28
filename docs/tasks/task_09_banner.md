# Task 09: Banner Selection (Developer 1 or Shared)

**Objective:** Ensure the correct banner file is loaded based on user input.

## Steps

1.  **Map Banner Names to Paths**
    *   In `internal/banner` or `cmd/ascii-art`, create a mapping:
        *   `shadow` -> `assets/banners/shadow.txt`
        *   `standard` -> `assets/banners/ascii_library.txt`
        *   `thinkertoy` -> `assets/banners/thinkertoy.txt`

2.  **Update Loader Call**
    *   In `cmd/ascii-art/main.go`, use `Config.BannerFile` (from Task 05) to determine which file to load.
    *   Pass the resolved path to `banner.LoadBanner`.

3.  **Verification**
    *   Ensure `go run . "hello" shadow` loads the shadow font.
    *   Ensure `go run . "hello"` defaults to standard.

## Acceptance Criteria
*   [ ] Supports `standard`, `shadow`, `thinkertoy`.
*   [ ] Returns error if banner is unknown or file missing.
*   [ ] Defaults to `standard` if not specified.