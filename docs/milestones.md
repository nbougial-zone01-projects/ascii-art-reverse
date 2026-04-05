# Project Milestones

## Milestone 1: Foundation
**Focus:** Project Setup & Shared Contracts
- [x] Task 00: Core Infrastructure
    - Project initialized.
    - Directory structure created.
    - Shared `Banner` model defined.

## Milestone 2: Data & Input Handling
**Focus:** Parsing & Validation
- [x] Task 01: Banner Management (Data Layer)
    - Banner loader implemented.
    - Unit tests for file reading and parsing passed.
- [x] Task 02: Input Processing (Input Layer)
    - CLI argument parser implemented.
    - Escape sequence handling (`\n`) working.
    - ASCII validation working.

## Milestone 3: Core Logic
**Focus:** Rendering Engine
- [x] Task 03: Rendering Engine (Logic Layer)
    - Single line rendering working.
    - Multi-line rendering working.
    - Edge cases (empty strings) handled.

## Milestone 4: Delivery
**Focus:** Integration & Quality Assurance
- [x] Task 04: Integration & Golden Tests
    - Application wired in `main.go`.
    - Golden files generated and verified.
    - Integration test suite passing.
    - Final `go test ./...` check passed.

## Milestone 5: Extended Features (Color, Output, Align)
**Focus:** Advanced Functionality
- [x] Task 05: Advanced Input Parsing
    - Implement flag parsing (`--color`, `--output`, `--align`).
    - Update `Config` model.
- [x] Task 06: Enhanced Rendering
    - Implement ANSI color application.
    - Implement text alignment (terminal size detection).
- [x] Task 07: File Output & Integration
    - Implement file writing in `main`.
    - Update integration tests for new features.

## Milestone 6: Reverse Feature
**Focus:** ASCII Art Reversal
- [ ] Task 10: Reverse Feature
    - Implement `--reverse=<fileName>` flag parsing with file-existence check.
    - Implement `internal/reverse` package with core reverse logic.
    - Wire reverse path into `cmd/ascii-art/main.go`.
    - Add integration tests GT-19..GT-22.
    - Full test suite passing.