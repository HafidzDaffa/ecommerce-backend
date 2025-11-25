.PHONY: help run build test swagger clean migrate-up migrate-down seed docker-up docker-down docker-build docker-logs docker-restart docker-clean

# Variables
APP_NAME=ecommerce-backend
MAIN_PATH=./cmd/api/main.go
BUILD_DIR=./bin
SWAGGER_DIR=./docs

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@echo "Starting server..."
	@go run $(MAIN_PATH)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/server $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/server"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@~/go/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -d ./,./internal/adapters/primary/http
	@echo "✅ Swagger docs generated successfully!"
	@echo "📖 Access at: http://localhost:8080/swagger/index.html"

clean: ## Clean build files
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

migrate-up: ## Run database migrations up
	@echo "Running migrations up..."
	@migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce_db?sslmode=disable" up

migrate-down: ## Run database migrations down
	@echo "Running migrations down..."
	@migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce_db?sslmode=disable" down

seed: ## Run database seeders
	@echo "Running seeders..."
	@psql -d ecommerce_db -f seeders/004_users.sql
	@psql -d ecommerce_db -f seeders/005_categories.sql
	@psql -d ecommerce_db -f seeders/006_products.sql
	@echo "✅ Seeders executed successfully!"

install-tools: ## Install required tools (swag, migrate)
	@echo "Installing tools..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "✅ Tools installed!"

dev: ## Run in development mode with hot reload (requires air)
	@air

docker-up: ## Start all docker containers
	@echo "Starting docker containers..."
	@docker compose up -d
	@echo "✅ Docker containers started!"

docker-down: ## Stop all docker containers
	@echo "Stopping docker containers..."
	@docker compose down
	@echo "✅ Docker containers stopped!"

docker-build: ## Build and start docker containers
	@echo "Building and starting docker containers..."
	@docker compose up -d --build
	@echo "✅ Docker containers built and started!"

docker-logs: ## Show docker logs
	@docker compose logs -f

docker-restart: ## Restart docker containers
	@echo "Restarting docker containers..."
	@docker compose restart
	@echo "✅ Docker containers restarted!"

docker-clean: ## Stop and remove all docker containers, volumes, and networks
	@echo "Cleaning docker resources..."
	@docker compose down -v
	@echo "✅ Docker resources cleaned!"

.DEFAULT_GOAL := help
