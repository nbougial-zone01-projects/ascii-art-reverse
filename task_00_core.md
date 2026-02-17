# Task 00: Core Infrastructure (Joint Team Task)

**Objective:** Initialize the project and define shared data structures (contracts) so all developers can work independently without blocking each other.

## Steps

1.  **Initialize Module**
    *   Run `go mod init ascii-art`.

2.  **Define Shared Types (`types.go`)**
    *   Create a file `types.go` to hold the shared structs.
    *   Define the Banner structure:
        ```go
        // Banner represents the font map. Key is the rune, Value is the 8 lines of ASCII art.
        type Banner map[rune][]string
        ```

3.  **Define Function Signatures (`main.go` or `interfaces.go`)**
    *   Create empty functions (stubs) to establish the API contract:
        *   `func LoadBanner(filename string) (Banner, error)`
        *   `func ParseInput(args []string) (string, error)`
        *   `func Render(input string, b Banner) string`
    *   *Note:* Do not implement the logic yet. Just ensure the code compiles.