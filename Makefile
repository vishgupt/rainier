.PHONY: build run test clean help proto deps setup fmt lint

help:
	@echo "Rainier Vector Database - Go Migration"
	@echo "Available commands:"
	@echo "  make setup    - Setup all dependencies (protoc, Go plugins, etc.)"
	@echo "  make deps     - Download Go module dependencies"
	@echo "  make proto    - Generate protobuf files"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make fmt      - Format code"
	@echo "  make lint     - Run linter"

setup: install-protoc install-go-plugins deps
	@echo "✓ All dependencies installed successfully"

install-protoc:
	@command -v protoc >/dev/null 2>&1 && echo "✓ protoc already installed" || { \
		echo "Installing protobuf compiler..."; \
		if command -v brew >/dev/null 2>&1; then \
			brew install protobuf; \
		elif command -v apt-get >/dev/null 2>&1; then \
			sudo apt-get update && sudo apt-get install -y protobuf-compiler; \
		else \
			echo "Please install protobuf compiler manually from https://github.com/protocolbuffers/protobuf/releases"; \
			exit 1; \
		fi; \
	}

install-go-plugins:
	@echo "Installing Go protobuf plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "✓ Go protobuf plugins installed"

proto:
	@command -v protoc >/dev/null 2>&1 || { echo "protoc is not installed. Run 'make setup' first."; exit 1; }
	@export PATH="$(shell go env GOPATH)/bin:$$PATH" && \
	protoc --go_out=. --go-grpc_out=. src/proto/vector_database.proto
	@echo "✓ Protobuf files generated successfully"

deps:
	@echo "Downloading Go module dependencies..."
	go mod download
	go mod tidy
	@echo "✓ Dependencies downloaded"

build: proto
	go build -o bin/rainier ./src/cmd

run: build
	./bin/rainier

test:
	go test -v ./src/...

clean:
	rm -rf bin/
	go clean

fmt:
	go fmt ./src/...

lint:
	golangci-lint run ./src/...
