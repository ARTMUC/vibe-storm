# VibeStorm - Modern Web Application Makefile

.PHONY: help build run test clean migrate-up migrate-down dev-setup docker-build docker-run

# Default target
help: ## Show this help message
	@echo "VibeStorm - Modern Web Application"
	@echo ""
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development commands
dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	@cp .env.example .env
	@echo "✅ Environment file created. Please update .env with your configuration."
	@echo "Installing Go dependencies..."
	@go mod tidy
	@echo "Installing Node.js dependencies..."
	@cd web && npm install
	@echo "✅ Development environment setup complete!"

build: ## Build the Go application
	@echo "Building Go application..."
	@go build -o bin/server cmd/server/main.go
	@echo "✅ Application built successfully!"

build-frontend: ## Build frontend assets
	@echo "Building frontend assets..."
	@cd web && npm run build
	@echo "✅ Frontend assets built successfully!"

generate-templ: ## Generate Go code from Templ templates
	@echo "Generating Go code from Templ templates..."
	@which templ >/dev/null || (echo "Installing templ..." && go install github.com/a-h/templ/cmd/templ@latest)
	@cd web/templates && templ generate
	@echo "✅ Templ templates generated successfully!"

run: ## Run the application
	@echo "Starting VibeStorm server..."
	@go run cmd/server/main.go

run-dev: ## Run in development mode with hot reload
	@echo "Starting VibeStorm server in development mode..."
	@go run cmd/server/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Database commands
migrate-up: ## Run database migrations
	@echo "Running database migrations..."
	@migrate -path migrations -database "mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)" up

migrate-down: ## Rollback database migrations
	@echo "Rolling back database migrations..."
	@migrate -path migrations -database "mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)" down

migrate-create: ## Create a new migration file
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $${name}

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t vibe-storm:latest .

docker-run: ## Run application in Docker
	@echo "Running application in Docker..."
	@docker run -p 8080:8080 --env-file .env vibe-storm:latest

# Cleanup commands
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf web/static/dist/
	@rm -f coverage.out coverage.html
	@echo "✅ Cleanup complete!"

clean-docker: ## Clean Docker artifacts
	@echo "Cleaning Docker artifacts..."
	@docker system prune -f
	@docker image rm vibe-storm:latest 2>/dev/null || true
	@echo "✅ Docker cleanup complete!"

# Development workflow
dev: dev-setup build ## Full development setup
	@echo "🎉 Development environment ready!"
	@echo "Run 'make run' to start the server"

# Production commands
prod-build: ## Build for production
	@echo "Building for production..."
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/server/main.go
	@cd web && npm run build
	@echo "✅ Production build complete!"

# API Documentation
swagger-gen: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	@which swag >/dev/null || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	@swag init -g main.go -o api/docs
	@echo "✅ Swagger documentation generated!"

swagger-serve: ## Serve Swagger UI (requires running application)
	@echo "Opening Swagger UI..."
	@open http://localhost:$(SERVER_PORT)/swagger/index.html || echo "Start the application first with 'make run'"

# Health check
health: ## Check application health
	@curl -f http://localhost:8080/api/v1/health || echo "❌ Application is not running"
	@curl -f http://localhost:8080/ || echo "❌ Frontend is not accessible"

# Load environment variables from .env file (if it exists)
-include .env
export
