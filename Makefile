.PHONY: all build clean test lint fmt vet tidy help

# Build variables
BINARY_NAME=beamlight
BUILD_DIR=./build/bin
GO=go
LDFLAGS=-ldflags "-X main.GitCommit=`git rev-parse HEAD` -X main.BuildDate=`date -u +%Y-%m-%dT%H:%M:%SZ`"

all: clean build

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build          Build the beamlight binary"
	@echo "  clean          Remove build artifacts"
	@echo "  test           Run tests"
	@echo "  lint           Run linters"
	@echo "  fmt            Run go fmt"
	@echo "  vet            Run go vet"
	@echo "  tidy           Run go mod tidy"
	@echo "  all            Clean and build"
	@echo "  help           Show this help"

build:
	@echo "Building beamlight..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/beamlight
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

test:
	@echo "Running tests..."
	$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "Tests complete"

lint:
	@echo "Running linters..."
	golangci-lint run ./...
	@echo "Linting complete"

fmt:
	@echo "Running go fmt..."
	$(GO) fmt ./...
	@echo "Formatting complete"

vet:
	@echo "Running go vet..."
	$(GO) vet ./...
	@echo "Vetting complete"

tidy:
	@echo "Running go mod tidy..."
	$(GO) mod tidy
	@echo "Tidy complete" 