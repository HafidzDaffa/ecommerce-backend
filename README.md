# E-commerce Backend

Backend API for e-commerce application built with Go Fiber using Hexagonal Architecture.

## Tech Stack

- **Go**: 1.21+
- **Web Framework**: Go Fiber v2
- **Architecture**: Hexagonal (Ports and Adapters)
- **Database**: PostgreSQL
- **Migrations**: golang-migrate (File-based)
- **Seeders**: Custom seeder implementation
- **Container**: Docker & Docker Compose
- **API Documentation**: Swagger UI
- **Testing**: Go Test + Testify

## Project Structure

```
backend/
├── main.go                      # Application entry point
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksums
├── .env.example                 # Environment variables template
├── .env.migrate                 # Migration environment variables
├── .gitignore                   # Git ignore rules
├── ecommerce.dbml              # Database schema
├── Makefile                     # Build & run commands
├── migrate                      # Migration CLI binary
│
├── common/                      # Shared utilities and error handling
│   ├── errors/                  # Custom error definitions
│   ├── middleware/             # HTTP middleware
│   └── utils/                   # Utility functions
│
├── domain/                      # Business logic layer
│   ├── entities/                # Domain entities
│   │   └── user.go
│   └── repositories/            # Repository interfaces
│       └── repositories.go
│
├── application/                 # Application layer
│   ├── dtos/                    # Data transfer objects
│   │   └── user_dto.go
│   ├── ports/                   # Port interfaces
│   │   └── server.go
│   └── usecases/                # Use case implementations
│       ├── user_usecase.go
│       └── user_usecase_test.go
│
├── infrastructure/              # Infrastructure implementations
│   ├── config/                  # Configuration management
│   │   └── config.go
│   ├── database/                # Database setup
│   │   ├── database.go
│   │   └── repositories.go
│   └── controllers/             # Controllers (if needed)
│
├── interfaces/                  # Interface adapters
│   └── http/                    # HTTP layer
│       ├── router.go            # Route definitions
│       └── user_handler.go      # HTTP handlers
│
├── migrations/                  # SQL migration files
│   ├── 000001_create_tables.up.sql
│   ├── 000001_create_tables.down.sql
│   └── ...
│
├── seeders/                    # Database seeder files
│   ├── seeders.go
│   ├── role_seeder.go
│   └── ...
│
├── cmd/                        # Command line tools
│   └── migrate.go             # Migration & seeder CLI
│
├── scripts/                    # Helper scripts
│   └── create_migration.sh    # Create new migration
│
├── postman/                    # Postman collections
│   ├── Ecommerce-API.postman_collection.json
│   └── Ecommerce-API.postman_environment.json
│
├── docs/                       # Swagger documentation (generated)
└── README.md                   # This file
```

## Architecture Layers

### 1. Domain Layer
Contains pure business logic without any external dependencies.
- **Entities**: Core business objects
- **Repositories**: Interface definitions for data access

### 2. Application Layer
Orchestrates business logic and defines use cases.
- **Use Cases**: Business operations
- **DTOs**: Data transfer between layers
- **Ports**: Interface definitions for external interactions

### 3. Infrastructure Layer
Provides technical implementations.
- **Database**: ORM and database connections
- **Config**: Configuration management
- **External APIs**: Third-party service integrations

### 4. Interfaces Layer
Handles external communications.
- **HTTP**: REST API endpoints and handlers

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker & Docker Compose
- PostgreSQL (optional, for database)

### Installation

#### Using Docker (Recommended)

1. Clone the repository
2. Navigate to the backend directory
3. Copy environment variables:

```bash
cp .env.example .env
```

4. Build and run containers:

```bash
make build
make run
```

Or using docker compose directly:

```bash
docker compose up -d
```

Note: Requires Docker Compose v2. See `DOCKER_COMPOSE.md` for details.

5. View logs:

```bash
make logs
```

6. Connect to PostgreSQL shell (optional):

```bash
make db-shell
```

#### Manual Installation

1. Clone the repository
2. Navigate to the backend directory
3. Copy environment variables:

```bash
cp .env.example .env
```

4. Install dependencies:

```bash
go mod download
```

5. Run application:

```bash
go run main.go
```

Or build and run:

```bash
go build -o bin/main .
./bin/main
```

## Database Migrations & Seeders

This project uses **file-based migrations** with **separate migration files** for each table. See `MIGRATIONS.md` for detailed documentation.

### Quick Start

```bash
# Start PostgreSQL
make db-run

# Run migrations
make migrate

# Check migration status
make migrate-status

# Run seeders
make seed
```

### Commands

```bash
make migrate            # Apply pending migrations
make migrate-down      # Rollback last migration
make migrate-status     # Show migration version and available migrations
make migrate-refresh   # Rollback all and migrate again
make seed             # Run database seeders
```

### Creating New Migration

```bash
# Create migration for new table
./scripts/create_migration.sh create_products_table

# This creates:
# migrations/000003_create_products_table.up.sql
# migrations/000003_create_products_table.down.sql
```

This creates two files:
- `migrations/TIMESTAMP_add_products_table.up.sql`
- `migrations/TIMESTAMP_add_products_table.down.sql`

## API Endpoints

### Health Check
```
GET /api/v1/health
```

### Users
```
POST   /api/v1/users/register  - Register new user
POST   /api/v1/users/login     - User login
GET    /api/v1/users/:id       - Get user by ID
PUT    /api/v1/users/:id       - Update user
DELETE /api/v1/users/:id       - Delete user
GET    /api/v1/users/          - List all users
```

### API Documentation (Swagger)
```
GET /swagger/*
```
Access Swagger UI at: http://localhost:8080/swagger/index.html

## API Testing

### Postman Collection
Import Postman collection from `postman/Ecommerce-API.postman_collection.json` into Postman.

**Features:**
- Pre-configured requests for all endpoints
- Environment variables for easy configuration
- Test scripts for response validation

### Swagger UI
Interactive API documentation available at:
- Production: http://localhost:8080/swagger/index.html
- Try out endpoints directly from browser
- View request/response schemas

## Environment Variables

| Variable       | Description                      | Default           |
|----------------|----------------------------------|-------------------|
| APP_ENV        | Application environment          | development       |
| PORT           | Server port                      | 8080              |
| DB_HOST        | Database host                    | localhost         |
| DB_PORT        | Database port                    | 5432              |
| DB_USER        | Database username                | ecommerce         |
| DB_PASSWORD    | Database password                | ecommerce123      |
| DB_NAME        | Database name                    | ecommerce         |
| JWT_SECRET     | JWT signing secret               | -                 |
| JWT_EXPIRY     | JWT token expiry time            | 24h               |

## Default Roles

Three roles are seeded by default:
- **ID 1: customer** - Customer dengan akses belanja standar
- **ID 2: toko** - Toko dengan akses manajemen produk
- **ID 3: admin** - Administrator dengan akses penuh ke sistem

## Current Migrations

| Version | Migration                  | Description             |
|---------|---------------------------|-----------------------|
| 001     | create_roles_table       | Create roles table     |
| 002     | create_users_table       | Create users table     |

Check migration status: `make migrate-status`

## Development

### Docker Commands

```bash
make build         # Build Docker images
make run           # Run all containers (backend + postgres)
make stop          # Stop all containers
make restart       # Restart all containers
make logs          # Show container logs
make clean         # Stop and remove containers and volumes
make db-run        # Run only PostgreSQL container
make db-stop       # Stop PostgreSQL container
make db-shell      # Connect to PostgreSQL shell
```

### Database Commands

```bash
make migrate            # Run database migrations
make migrate-down      # Rollback last migration
make migrate-refresh   # Rollback all and migrate again
make seed             # Run database seeders
```

### Swagger & Testing Commands

```bash
make swagger-docs   # Generate Swagger documentation
make test           # Run unit tests
make test-coverage  # Run unit tests with coverage report
```

### Build

```bash
go build -o bin/main .
```

### Run Tests

```bash
go test ./...
```

### Format Code

```bash
go fmt ./...
```

## License

MIT
