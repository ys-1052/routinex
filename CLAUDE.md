# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Communication Language

**IMPORTANT**: Always respond in Japanese when working with this repository. All explanations, comments, and communications should be in Japanese.

## Project Overview

RoutineX is a routine management application with a Go backend using the Echo web framework. The project follows a monorepo structure with separate backend and frontend directories.

## Backend Architecture

The backend uses a standard Go project layout:

- `cmd/server/main.go` - Main application entry point using Echo v4
- `internal/` - Private application code
  - `handlers/` - HTTP request handlers  
  - `middleware/` - Custom middleware
  - `models/` - Data models
  - `repository/` - Data access layer
  - `service/` - Business logic layer
- `pkg/` - Public library code
  - `config/` - Configuration management
  - `utils/` - Utility functions

The main server runs on port 8080 and includes logging and recovery middleware by default.

## Development Commands

### Backend Development (run from `/backend` directory)

**IMPORTANT**: 
- Always run formatting and linting before committing code
- .gitignore should ONLY contain files/directories that actually exist in the project
- NEVER add future or speculative entries to .gitignore - only add what currently exists
- When new files/directories are created, update .gitignore at that time, not before

```bash
# Format code (REQUIRED before commits)
make fmt
# OR manually:
gofmt -w -s .
goimports -w -local github.com/ys-1052/routinex/backend .

# Lint code (REQUIRED before commits)
make lint
# OR manually:
golangci-lint run

# Full development workflow (format + lint + test + build)
make dev

# Run the server
make run
# OR manually:
go run cmd/server/main.go

# Build the application  
make build
# OR manually:
go build -o bin/server cmd/server/main.go

# Run tests
make test
# OR manually:
go test -v ./...

# Update dependencies
go mod tidy

# Clean build artifacts
make clean
```

### Docker

```bash
# Build Docker image
docker build -t routinex-backend .

# Run container
docker run -p 8080:8080 routinex-backend
```

## Key Endpoints

- `GET /` - API status endpoint
- `GET /health` - Health check endpoint

## Module Information

- Go version: 1.24.4
- Main framework: Echo v4.13.4
- Module path: `github.com/ys-1052/routinex/backend`