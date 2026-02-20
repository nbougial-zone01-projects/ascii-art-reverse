.PHONY: all build test clean run run-examples fmt

BINARY_NAME := ascii-art
SOURCE_DIR := ./cmd/ascii-art

all: build

build:
	go build -o $(BINARY_NAME) $(SOURCE_DIR)

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

# Run with default argument "Hello". Override with: make run ARGS="something else"
ARGS ?= "Hello"
run:
	go run $(SOURCE_DIR) $(ARGS)

run-examples:
	go run $(SOURCE_DIR) "Hello"
	go run $(SOURCE_DIR) "wfjeh"
	go run $(SOURCE_DIR) "Hello\nWorld"
	go run $(SOURCE_DIR) "12345"

fmt:
	go fmt ./...