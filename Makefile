.PHONY: help build run stop restart logs clean db-run db-stop db-shell swagger test test-coverage swagger-docs

help:
	@echo "Available commands:"
	@echo "  make build         - Build Docker images"
	@echo "  make run           - Run all containers (backend + postgres)"
	@echo "  make stop          - Stop all containers"
	@echo "  make restart       - Restart all containers"
	@echo "  make logs          - Show container logs"
	@echo "  make clean         - Stop and remove containers and volumes"
	@echo "  make db-run        - Run only PostgreSQL container"
	@echo "  make db-stop       - Stop PostgreSQL container"
	@echo "  make db-shell      - Connect to PostgreSQL shell"
	@echo "  make swagger-docs  - Generate Swagger documentation"
	@echo "  make test          - Run unit tests"
	@echo "  make test-coverage - Run unit tests with coverage"

build:
	docker-compose build

run:
	docker-compose up -d
	@echo "Waiting for services to be ready..."
	@sleep 5
	@docker-compose logs

stop:
	docker-compose down

restart:
	docker-compose restart

logs:
	docker-compose logs -f

clean:
	docker-compose down -v
	docker system prune -f

db-run:
	docker-compose -f docker-compose.dev.yml up -d
	@echo "PostgreSQL is running on localhost:5432"
	@echo "Use 'make db-shell' to connect"

db-stop:
	docker-compose -f docker-compose.dev.yml down

db-shell:
	docker exec -it ecommerce-postgres psql -U ecommerce -d ecommerce

swagger-docs:
	swag init -g main.go -o docs
	@echo "Swagger documentation generated successfully"

test:
	go test ./... -v

test-coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
