# E-commerce Backend

Backend API for e-commerce application built with Go Fiber using Hexagonal Architecture.

## Tech Stack

- **Go**: 1.21+
- **Web Framework**: Go Fiber v2
- **Architecture**: Hexagonal (Ports and Adapters)
- **Database**: PostgreSQL
- **ORM**: GORM
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
├── .gitignore                   # Git ignore rules
├── ecommerce.dbml              # Database schema
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
│       └── user_usecase.go
│
├── infrastructure/              # Infrastructure implementations
│   ├── config/                  # Configuration management
│   │   └── config.go
│   ├── database/                # Database setup and migrations
│   │   ├── database.go
│   │   ├── repositories.go
│   │   └── migration.go
│   └── controllers/             # Controllers (if needed)
│
└── interfaces/                  # Interface adapters
    └── http/                    # HTTP layer
        ├── router.go            # Route definitions
        └── user_handler.go      # HTTP handlers
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

Or using docker-compose directly:

```bash
docker-compose up -d
```

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

5. Run the application:

```bash
go run main.go
```

Or build and run:

```bash
go build -o bin/main
./bin/main
```

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
Access the Swagger UI at: http://localhost:8080/swagger/index.html

## API Testing

### Postman Collection
Import the Postman collection from `postman/Ecommerce-API.postman_collection.json` into Postman.

**Features:**
- Pre-configured requests for all endpoints
- Environment variables for easy configuration
- Test scripts for response validation

### Swagger UI
Interactive API documentation available at:
- Production: http://localhost:8080/swagger/index.html
- Try out endpoints directly from the browser
- View request/response schemas

## Environment Variables

| Variable       | Description                      | Default           |
|----------------|----------------------------------|-------------------|
| APP_ENV        | Application environment          | development       |
| PORT           | Server port                      | 8080              |
| DB_HOST        | Database host                    | localhost         |
| DB_PORT        | Database port                    | 5432              |
| DB_USER        | Database username                | postgres          |
| DB_PASSWORD    | Database password                | -                 |
| DB_NAME        | Database name                    | ecommerce         |
| JWT_SECRET     | JWT signing secret               | -                 |
| JWT_EXPIRY     | JWT token expiry time            | 24h               |

## Development

### Docker Commands

```bash
make build         - Build Docker images
make run           - Run all containers (backend + postgres)
make stop          - Stop all containers
make restart       - Restart all containers
make logs          - Show container logs
make clean         - Stop and remove containers and volumes
make db-run        - Run only PostgreSQL container
make db-stop       - Stop PostgreSQL container
make db-shell      - Connect to PostgreSQL shell
```

### Swagger & Testing Commands

```bash
make swagger-docs   - Generate Swagger documentation
make test           - Run unit tests
make test-coverage  - Run unit tests with coverage report
```

### Build

```bash
go test ./...
```

### Format Code

```bash
go fmt ./...
```

## License

MIT
