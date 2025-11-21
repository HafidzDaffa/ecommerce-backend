.PHONY: help run build docker-up docker-down docker-restart docker-logs docker-status backend-logs clean test deps migrate-up migrate-down migrate-create migrate-force seed seed-run

# Database configuration
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= ecommerce_db
DB_SSL_MODE ?= disable
MIGRATIONS_PATH = migrations
DATABASE_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# Default target
help:
	@echo "Available commands:"
	@echo "  make run            - Run the application locally (foreground)"
	@echo "  make build          - Build the application binary"
	@echo "  make docker-up      - Start PostgreSQL + backend server"
	@echo "  make docker-down    - Stop backend server + PostgreSQL"
	@echo "  make docker-status  - Show status of all services"
	@echo "  make docker-restart - Restart all services"
	@echo "  make docker-logs    - Show PostgreSQL logs"
	@echo "  make backend-logs   - Show backend logs"
	@echo "  make clean          - Remove binary and clean up"
	@echo "  make test           - Run tests"
	@echo "  make deps           - Download dependencies"
	@echo "  make tidy           - Tidy go.mod"
	@echo ""
	@echo "Migration commands:"
	@echo "  make migrate-up     - Run all pending migrations"
	@echo "  make migrate-down   - Rollback last migration"
	@echo "  make migrate-create - Create new migration (usage: make migrate-create NAME=create_table_name)"
	@echo "  make migrate-force  - Force migration version (usage: make migrate-force VERSION=1)"
	@echo "  make migrate-version - Show current migration version"
	@echo ""
	@echo "Seeder commands:"
	@echo "  make seed           - Run all database seeders"
	@echo "  make seed-run       - Alternative command to run seeders"

# Run application locally
run:
	@echo "Running application..."
	go run cmd/api/main.go

# Build application
build:
	@echo "Building application..."
	go build -o bin/main cmd/api/main.go

# Start PostgreSQL + Backend
docker-up:
	@echo "🚀 Starting all services..."
	@echo "1️⃣  Starting PostgreSQL..."
	@docker compose up -d postgres
	@echo "⏳ Waiting for PostgreSQL to be ready..."
	@sleep 5
	@echo "2️⃣  Starting backend server..."
	@./scripts/start-backend.sh
	@echo ""
	@echo "✅ All services are running!"
	@echo "📊 PostgreSQL: localhost:5432"
	@echo "🌐 Backend API: http://localhost:8080"
	@echo ""
	@echo "💡 Commands:"
	@echo "   make docker-status  - Check status"
	@echo "   make backend-logs   - View backend logs"
	@echo "   make docker-down    - Stop all services"

# Stop Backend + PostgreSQL
docker-down:
	@echo "🛑 Stopping all services..."
	@./scripts/stop-backend.sh
	@docker compose down
	@echo "✅ All services stopped"

# Show status of all services
docker-status:
	@echo "📊 Service Status:"
	@echo "==================="
	@echo ""
	@echo "PostgreSQL:"
	@docker compose ps postgres || echo "  ❌ Not running"
	@echo ""
	@echo "Backend:"
	@if [ -f .backend.pid ]; then \
		PID=$$(cat .backend.pid); \
		if ps -p $$PID > /dev/null 2>&1; then \
			echo "  ✅ Running (PID: $$PID)"; \
		else \
			echo "  ❌ Not running (stale PID file)"; \
		fi \
	else \
		echo "  ❌ Not running"; \
	fi

# Restart all services
docker-restart:
	@echo "🔄 Restarting all services..."
	@$(MAKE) docker-down
	@sleep 2
	@$(MAKE) docker-up

# Show PostgreSQL logs
docker-logs:
	@echo "📝 PostgreSQL logs:"
	@docker compose logs -f postgres

# Show backend logs
backend-logs:
	@if [ -f backend.log ]; then \
		tail -f backend.log; \
	else \
		echo "❌ Backend log file not found. Is backend running?"; \
	fi

# Clean up
clean:
	@echo "Cleaning up..."
	rm -rf bin/
	go clean

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download

# Tidy go.mod
tidy:
	@echo "Tidying go.mod..."
	go mod tidy

# Create .env from example
env:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from .env.example"; \
	else \
		echo ".env file already exists"; \
	fi

# Migration commands
migrate-up:
	@echo "Running migrations..."
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up; \
	else \
		echo "Error: migrate tool not found. Install it with:"; \
		echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-down:
	@echo "Rolling back migration..."
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1; \
	else \
		echo "Error: migrate tool not found. Install it with:"; \
		echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=create_table_name"; \
		exit 1; \
	fi
	@if command -v migrate > /dev/null; then \
		migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME); \
	else \
		echo "Error: migrate tool not found. Install it with:"; \
		echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(VERSION); \
	else \
		echo "Error: migrate tool not found. Install it with:"; \
		echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-version:
	@if command -v migrate > /dev/null; then \
		migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version; \
	else \
		echo "Error: migrate tool not found. Install it with:"; \
		echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

# Seeder commands
seed:
	@echo "Running database seeders..."
	@go run cmd/seeder/main.go

seed-run:
	@echo "Running database seeders..."
	@go run cmd/seeder/main.go
