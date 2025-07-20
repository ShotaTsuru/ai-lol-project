.PHONY: help build up down logs clean dev-frontend dev-backend test

# Help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Docker commands
build: ## Build all Docker images
	docker-compose build

up: ## Start all services
	docker-compose up -d

down: ## Stop all services
	docker-compose down

logs: ## Show logs for all services
	docker-compose logs -f

clean: ## Clean up Docker resources
	docker-compose down -v --remove-orphans
	docker system prune -f

# Development commands
dev: ## Start development environment
	docker-compose up

dev-frontend: ## Start frontend development server
	cd frontend && npm run dev

dev-backend: ## Start backend development server
	cd backend && go run main.go

# Database commands
db-up: ## Start only database services
	docker-compose up -d postgres redis

db-down: ## Stop database services
	docker-compose stop postgres redis

# Testing
test-frontend: ## Run frontend tests
	cd frontend && npm test

test-frontend-e2e: ## Run frontend E2E tests
	cd frontend && npm run test:e2e

test-backend: ## Run backend tests
	cd backend && go test ./...

test-backend-unit: ## Run backend unit tests
	cd backend && go test -v -race -coverprofile=coverage.out ./...

test-backend-integration: ## Run backend integration tests
	cd backend && go test -v -tags=integration ./tests/integration/...

test-coverage: ## Show test coverage
	cd backend && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: backend/coverage.html"

test: test-frontend test-backend ## Run all tests

test-all: test-frontend test-frontend-e2e test-backend-unit test-backend-integration ## Run all tests including E2E

# Installation
install: ## Install dependencies
	cd frontend && npm install
	cd backend && go mod tidy

# OpenAPI code generation
generate-openapi: ## Generate DTOs and API code from OpenAPI specification
	@echo "🚀 Generating OpenAPI code..."
	@chmod +x scripts/generate-dto.sh
	@./scripts/generate-dto.sh

generate-openapi-types: ## Generate only DTO types from OpenAPI specification
	@echo "📝 Generating OpenAPI types..."
	@if ! command -v oapi-codegen &> /dev/null; then \
		echo "📦 Installing oapi-codegen..."; \
		go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest; \
	fi
	@mkdir -p backend/controllers/dto
	@oapi-codegen -package dto -generate types openapi.yaml > backend/controllers/dto/generated_types.go
	@echo "✅ OpenAPI types generated"

generate-openapi-server: ## Generate server code from OpenAPI specification
	@echo "🖥️  Generating OpenAPI server code..."
	@if ! command -v oapi-codegen &> /dev/null; then \
		echo "📦 Installing oapi-codegen..."; \
		go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest; \
	fi
	@mkdir -p backend/controllers/dto
	@oapi-codegen -package dto -generate server openapi.yaml > backend/controllers/dto/generated_server.go
	@echo "✅ OpenAPI server code generated"

generate-openapi-spec: ## Generate spec validation from OpenAPI specification
	@echo "🔍 Generating OpenAPI spec validation..."
	@if ! command -v oapi-codegen &> /dev/null; then \
		echo "📦 Installing oapi-codegen..."; \
		go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest; \
	fi
	@mkdir -p backend/controllers/dto
	@oapi-codegen -package dto -generate spec openapi.yaml > backend/controllers/dto/generated_spec.go
	@echo "✅ OpenAPI spec validation generated"

validate-openapi: ## Validate OpenAPI specification
	@echo "🔍 Validating OpenAPI specification..."
	@if ! command -v oapi-codegen &> /dev/null; then \
		echo "📦 Installing oapi-codegen..."; \
		go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest; \
	fi
	@oapi-codegen -package dto -generate types openapi.yaml > /dev/null
	@echo "✅ OpenAPI specification is valid"

openapi-docs: ## Generate OpenAPI documentation
	@echo "📚 Generating OpenAPI documentation..."
	@if ! command -v swag &> /dev/null; then \
		echo "📦 Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@mkdir -p docs
	@swag init -g backend/main.go -o docs
	@echo "✅ OpenAPI documentation generated"

clean-openapi: ## Clean generated OpenAPI files
	@echo "🧹 Cleaning generated OpenAPI files..."
	@rm -f backend/controllers/dto/generated_*.go
	@echo "✅ Generated OpenAPI files cleaned"

watch-openapi: ## Watch OpenAPI file changes and auto-generate code
	@echo "👀 Starting OpenAPI file watcher..."
	@chmod +x scripts/openapi-watch.sh
	@./scripts/openapi-watch.sh

# Environment setup
setup: ## Set up development environment
	cp env.example .env
	@echo "Please edit .env file with your configuration"

# Build for production
build-prod: ## Build for production
	cd frontend && npm run build
	cd backend && go build -o main main.go

# Development workflow
dev-setup: install generate-openapi ## Setup development environment with OpenAPI
	@echo "✅ Development environment setup complete"

dev-backend-with-openapi: generate-openapi ## Start backend with OpenAPI generation
	@echo "🔄 Starting backend with latest OpenAPI code..."
	cd backend && go run main.go

# GitHub MCP setup
mcp-setup: ## Setup GitHub MCP server configuration
	@echo "🔧 Setting up GitHub MCP server..."
	@if [ ! -f ".vscode/mcp.json" ]; then \
		echo "📋 Creating MCP configuration from sample..."; \
		mkdir -p .vscode; \
		cp .vscode/mcp.json.sample .vscode/mcp.json; \
		echo "✅ MCP configuration file created"; \
	else \
		echo "✅ MCP configuration file already exists"; \
	fi
	@echo ""
	@echo "📚 Please follow the instructions in MCP_SETUP_GUIDE.md"
	@echo "1. Create a GitHub Personal Access Token"
	@echo "2. Open this project in Visual Studio Code"
	@echo "3. Click the 'Start' button in .vscode/mcp.json"
	@echo "4. Enter your token when prompted"

mcp-check: ## Check MCP configuration
	@echo "🔍 Checking MCP configuration..."
	@if [ -f ".vscode/mcp.json" ]; then \
		echo "✅ MCP configuration file found"; \
	else \
		echo "❌ MCP configuration file not found"; \
		echo "💡 Run 'make mcp-setup' to create it"; \
		exit 1; \
	fi

mcp-clean: ## Clean MCP configuration
	@echo "🧹 Cleaning MCP configuration..."
	@rm -f .vscode/mcp.json
	@echo "MCP configuration removed"
	@echo "💡 Run 'make mcp-setup' to recreate it"

mcp-reset: ## Reset MCP configuration from sample
	@echo "🔄 Resetting MCP configuration..."
	@rm -f .vscode/mcp.json
	@cp .vscode/mcp.json.sample .vscode/mcp.json
	@echo "✅ MCP configuration reset from sample"

# League of Legends MCP Server
lol-mcp-setup: ## Setup League of Legends MCP server
	@echo "🎮 Setting up League of Legends MCP server..."
	@chmod +x scripts/setup-lol-mcp.sh
	@./scripts/setup-lol-mcp.sh

lol-mcp-build: ## Build League of Legends MCP server
	@echo "🔨 Building League of Legends MCP server..."
	@cd mcp-servers/lol-mcp-server && npm run build

lol-mcp-start: ## Start League of Legends MCP server
	@echo "🚀 Starting League of Legends MCP server..."
	@cd mcp-servers/lol-mcp-server && npm start

lol-mcp-dev: ## Start League of Legends MCP server in development mode
	@echo "🛠️  Starting League of Legends MCP server in development mode..."
	@cd mcp-servers/lol-mcp-server && npm run dev

lol-mcp-clean: ## Clean League of Legends MCP server build
	@echo "🧹 Cleaning League of Legends MCP server..."
	@cd mcp-servers/lol-mcp-server && rm -rf node_modules dist .env
	@echo "✅ League of Legends MCP server cleaned"
