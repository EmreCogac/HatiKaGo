.PHONY: help build run test clean migrate docker-build docker-up docker-down

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building application..."
	go build -o bin/api cmd/api/main.go
	@echo "Build complete: bin/api"

run: ## Run the application
	@echo "Running application..."
	go run cmd/api/main.go

test: ## Run tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	@echo "Clean complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies ready"

migrate: ## Run database migrations
	@echo "Running migrations..."
	go run cmd/api/main.go migrate
	@echo "Migrations complete"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t hatikago-api:latest .
	@echo "Docker image built"

docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Containers started"
 
swagger: ## Generate swagger documentation
	@echo "Generating Swagger documentation..."
	swag init -g cmd/api/main.go
	@echo "Swagger docs generated at docs/"
	@echo "Access at: http://localhost:8080/swagger/index.html"
	
docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	docker-compose down
	@echo "Containers stopped"

docker-logs: ## Show Docker logs
	docker-compose logs -f

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete"

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run
	@echo "Lint complete"

dev: ## Run in development mode with hot reload
	@echo "Starting development mode..."
	air

.DEFAULT_GOAL := help
