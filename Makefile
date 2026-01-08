include .env.migrate

.PHONY: help build run stop restart logs clean db-run db-stop db-shell swagger test test-coverage swagger-docs migrate migrate-down seed migrate-status migrate-refresh

help:
	@echo "Available commands:"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make build              - Build Docker images"
	@echo "  make run                - Run all containers (backend + postgres)"
	@echo "  make stop               - Stop and remove all containers"
	@echo "  make pause              - Stop containers (without removing)"
	@echo "  make restart            - Restart all containers"
	@echo "  make logs               - Show container logs"
	@echo "  make clean              - Stop and remove containers and volumes"
	@echo "  make db-run             - Run only PostgreSQL container"
	@echo "  make db-stop            - Stop PostgreSQL container"
	@echo "  make db-shell           - Connect to PostgreSQL shell"
	@echo ""
	@echo "Database Commands (Local):"
	@echo "  make migrate            - Run database migrations"
	@echo "  make migrate-down        - Rollback last migration"
	@echo "  make migrate-status     - Show migration version"
	@echo "  make migrate-refresh     - Rollback all and migrate again"
	@echo "  make seed               - Run database seeders"
	@echo ""
	@echo "Development Commands:"
	@echo "  make swagger-docs       - Generate Swagger documentation"
	@echo "  make test               - Run unit tests"
	@echo "  make test-coverage      - Run unit tests with coverage"

build:
	docker compose build

run:
	docker compose up -d
	@echo "Waiting for services to be ready..."
	@sleep 5
	@docker compose logs

stop:
	docker compose down

pause:
	docker compose stop

restart:
	docker compose restart

logs:
	docker compose logs -f

clean:
	docker compose down -v
	docker system prune -f

db-run:
	docker compose -f docker-compose.dev.yml up -d
	@echo "PostgreSQL is running on localhost:5432"
	@echo "Use 'make db-shell' to connect"

db-stop:
	docker compose -f docker-compose.dev.yml down

db-shell:
	docker exec -it ecommerce-postgres psql -U ecommerce -d ecommerce

migrate:
	@echo "Running database migrations..."
	./migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations up

migrate-down:
	@echo "Rolling back last migration..."
	./migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations down 1

migrate-refresh:
	@echo "Rolling back all migrations..."
	./migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations down -all
	@echo "Running migrations again..."
	@make migrate

migrate-status:
	@echo "=== Migration Status ==="
	@echo ""
	@echo "Current version:"
	@./migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations version || echo "  No migrations applied yet"
	@echo ""
	@echo "Available migrations:"
	@ls -1 migrations/*.up.sql 2>/dev/null | sed 's/migrations\///g' | sed 's/.up.sql//g' | while read file; do echo "  - $$file"; done
	@echo ""

seed:
	@echo "Running database seeders..."
	go run cmd/migrate.go seed

swagger-docs:
	swag init -g main.go -o docs
	@echo "Swagger documentation generated successfully"

test:
	go test ./... -v

test-coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
