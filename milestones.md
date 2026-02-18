# Project Milestones

## Milestone 1: Foundation
**Focus:** Project Setup & Shared Contracts
- [ ] Task 00: Core Infrastructure
    - Project initialized.
    - Directory structure created.
    - Shared `Banner` model defined.

## Milestone 2: Data & Input Handling
**Focus:** Parsing & Validation
- [ ] Task 01: Banner Management (Data Layer)
    - Banner loader implemented.
    - Unit tests for file reading and parsing passed.
- [ ] Task 02: Input Processing (Input Layer)
    - CLI argument parser implemented.
    - Escape sequence handling (`\n`) working.
    - ASCII validation working.

## Milestone 3: Core Logic
**Focus:** Rendering Engine
- [ ] Task 03: Rendering Engine (Logic Layer)
    - Single line rendering working.
    - Multi-line rendering working.
    - Edge cases (empty strings) handled.

## Milestone 4: Delivery
**Focus:** Integration & Quality Assurance
- [ ] Task 04: Integration & Golden Tests
    - Application wired in `main.go`.
    - Golden files generated and verified.
    - Integration test suite passing.
    - Final `go test ./...` check passed.