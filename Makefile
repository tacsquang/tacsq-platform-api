# Makefile for Go API Backend

# Define the Go binary name
BINARY_NAME=go-api-backend

# Define the Go source files
SRC=$(shell find . -name '*.go')

# Define the build directory
BUILD_DIR=bin

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC)

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/*

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

# Run migrations
.PHONY: migrate
migrate:
	@echo "Running migrations..."
	sh scripts/migrate.sh

# Help message
.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make build     - Build the application"
	@echo "  make run       - Run the application"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make test      - Run tests"
	@echo "  make migrate    - Run database migrations"
	@echo "  make help      - Show this help message"