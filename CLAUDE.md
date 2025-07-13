# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Architecture

This is a full-stack web application template with:
- **Frontend**: Next.js 15 with TypeScript, TailwindCSS, and React Query
- **Backend**: Go with Gin framework, GORM, PostgreSQL, and Redis
- **MCP Server**: Python-based League of Legends MCP server using fast-mcp

The project follows clean architecture patterns with clear separation between controllers, services, and models in the backend.

## Essential Commands

### Development Environment
```bash
make dev                    # Start full development environment with Docker Compose
make dev-frontend           # Frontend only: cd frontend && npm run dev
make dev-backend           # Backend only: cd backend && go run main.go
make db-up                 # Start only PostgreSQL and Redis databases
```

### Dependencies
```bash
make install               # Install all dependencies (npm + go mod tidy)
```

### Testing
```bash
make test                  # Run all tests (frontend + backend)
make test-frontend         # cd frontend && npm test
make test-backend          # cd backend && go test ./...
make test-backend-unit     # Go unit tests with coverage
make test-backend-integration  # Go integration tests
make test-coverage         # Generate HTML coverage report
```

### Code Quality
```bash
npm run lint               # Frontend ESLint (from frontend/ directory)
```

### Production Build
```bash
make build-prod            # Build both frontend and backend for production
```

### MCP Servers
```bash
make mcp-setup             # Setup GitHub MCP server configuration
make lol-mcp-setup         # Setup League of Legends MCP server
make lol-mcp-dev           # Start LoL MCP server in development mode
```

## Project Structure

### Backend (`backend/`)
- `main.go` - Application entry point
- `config/` - Database and Redis configuration
- `controllers/` - HTTP request handlers
- `models/` - GORM database models
- `routes/` - API route definitions
- `services/` - Business logic layer
- `utils/` - Utility functions

### Frontend (`frontend/`)
- `src/app/` - Next.js App Router pages and layouts
- Uses TypeScript, TailwindCSS, and Radix UI components
- React Query for server state management

### MCP Servers (`mcp-servers/`)
- `lol-mcp-server/` - Python-based League of Legends API integration
- Uses fast-mcp framework for Model Context Protocol

## Database

The application uses PostgreSQL as the main database and Redis for caching. Database configuration is in `backend/config/database.go`. GORM models are defined in `backend/models/`.

## Environment Setup

1. Copy `env.example` to `.env` and configure
2. Run `make setup` for initial environment setup
3. The Docker Compose setup includes PostgreSQL and Redis services

## Key Technologies

- **Go 1.24+** with Gin web framework
- **Next.js 15** with App Router
- **PostgreSQL 15** and **Redis 7**
- **GORM** for database ORM
- **TypeScript** and **TailwindCSS**
- **React Query** for data fetching
- **fast-mcp** for MCP server implementation

## Development Notes

- The Go module is named `reverse-engineering-backend`
- PostgreSQL runs on port 5432, Redis on 6379, backend on 8080, frontend on 3000
- All services run in Docker containers with volume mounts for hot reloading
- The project includes AI-driven issue creation and project management systems