.PHONY: fmt lint test build run clean

# Formatting
fmt:
	@echo "Running gofmt..."
	@gofmt -w -s .
	@echo "Running goimports..."
	@goimports -w -local github.com/ys-1052/routinex/backend .

# Linting
lint:
	@echo "Running golangci-lint..."
	@golangci-lint run

# Testing
test:
	@echo "Running tests..."
	@go test -v ./...

# Build
build:
	@echo "Building application..."
	@go build -o bin/server cmd/server/main.go

# Run
run:
	@echo "Starting server..."
	@go run cmd/server/main.go

# Development workflow
dev: fmt lint test build

# Clean
clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@go clean