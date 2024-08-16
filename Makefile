# Simple Makefile for a Go project

# Build the application
all: build

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run test clean
		
