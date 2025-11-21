# E-Commerce Backend

Backend API untuk aplikasi e-commerce menggunakan **Hexagonal Architecture** (Ports & Adapters Pattern).

## Tech Stack

- **Go 1.21** - Programming Language
- **Fiber v2** - Web Framework
- **PostgreSQL 16** - Database
- **Docker & Docker Compose** - Containerization
- **sqlx** - SQL Extensions for Go

## Architecture

Proyek ini menggunakan **Hexagonal Architecture** dengan struktur sebagai berikut:

```
backend/
├── cmd/
│   └── api/
│       └── main.go                    # Entry point aplikasi
├── internal/
│   ├── core/                          # Business Logic Layer
│   │   ├── domain/                    # Domain entities
│   │   ├── ports/                     # Interfaces (input & output ports)
│   │   └── services/                  # Business logic implementation
│   ├── adapters/
│   │   ├── primary/                   # Input adapters
│   │   │   └── http/                  # HTTP handlers (Fiber)
│   │   └── secondary/                 # Output adapters
│   │       └── repository/            # Database repositories
│   └── infrastructure/                # Infrastructure Layer
│       ├── config/                    # Configuration management
│       ├── database/                  # Database connection
│       └── server/                    # Server setup
├── pkg/                               # Shared utilities
├── migrations/                        # Database migrations
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── go.mod
```

### Hexagonal Architecture Layers

1. **Domain Layer** (`internal/core/domain/`)
   - Entities dan business rules inti
   - Tidak bergantung pada framework atau infrastruktur

2. **Ports Layer** (`internal/core/ports/`)
   - Input Ports: Interface untuk use cases
   - Output Ports: Interface untuk repositories

3. **Services Layer** (`internal/core/services/`)
   - Implementasi business logic
   - Orchestrasi operasi domain

4. **Adapters Layer** (`internal/adapters/`)
   - **Primary (Input)**: HTTP handlers, gRPC, etc.
   - **Secondary (Output)**: Database repositories, external APIs

5. **Infrastructure Layer** (`internal/infrastructure/`)
   - Configuration, database connection, server setup

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Make (optional)
- Go 1.21+ (jika ingin run tanpa Docker)

### Quick Start dengan Docker

1. **Clone repository**
   ```bash
   cd backend
   ```

2. **Buat file .env**
   ```bash
   make env
   # atau
   cp .env.example .env
   ```

3. **Start semua services**
   ```bash
   make docker-up
   ```

4. **Install migrate CLI tool**
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

5. **Run database migrations**
   ```bash
   make migrate-up
   ```

6. **Run database seeders**
   ```bash
   make seed
   ```

7. **Cek status**
   ```bash
   curl http://localhost:8080/api/health
   ```

### Available Make Commands

**Application Commands:**
```bash
make help              # Lihat semua commands
make run               # Run locally (without Docker)
make build             # Build binary
make test              # Run tests
make clean             # Clean up
make deps              # Download dependencies
make tidy              # Tidy go.mod
make env               # Create .env from .env.example
```

**Docker Commands:**
```bash
make docker-up         # Start all services
make docker-down       # Stop all services
make docker-restart    # Restart all services
make docker-rebuild    # Rebuild and restart
make docker-logs       # Show all logs
make docker-logs-be    # Show backend logs
make docker-logs-db    # Show database logs
```

**Migration Commands:**
```bash
make migrate-up        # Run all pending migrations
make migrate-down      # Rollback last migration
make migrate-version   # Show current migration version
make migrate-create NAME=create_something  # Create new migration
make migrate-force VERSION=1               # Force set migration version
```

**Seeder Commands:**
```bash
make seed              # Run all database seeders
make seed-run          # Alternative command to run seeders
```

### Run Tanpa Docker

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Setup PostgreSQL** (pastikan sudah running)

3. **Setup .env file**
   ```bash
   cp .env.example .env
   # Edit .env sesuai kebutuhan
   ```

4. **Install migrate CLI tool**
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

5. **Run database migrations**
   ```bash
   make migrate-up
   ```

6. **Run database seeders**
   ```bash
   make seed
   ```

7. **Run application**
   ```bash
   make run
   # atau
   go run cmd/api/main.go
   ```

## API Endpoints

### Health Check
```
GET /api/health
```

### API Version 1
```
GET /api/v1/
```

## Environment Variables

Lihat file `.env.example` untuk daftar lengkap environment variables.

Key variables:
- `APP_PORT` - Server port (default: 8080)
- `DB_HOST` - Database host
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password

## Development

### Adding New Features

1. **Define Domain Entity** di `internal/core/domain/`
2. **Define Ports** (interfaces) di `internal/core/ports/`
3. **Implement Service** di `internal/core/services/`
4. **Implement Repository** di `internal/adapters/secondary/repository/`
5. **Create HTTP Handler** di `internal/adapters/primary/http/`
6. **Register Routes** di `internal/infrastructure/server/fiber.go`

## Database Schema

Lihat file `dbml.dbml` untuk detail skema database.

### Database Migrations

Project ini menggunakan [golang-migrate](https://github.com/golang-migrate/migrate) untuk database migrations.

#### Install Migrate CLI

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

#### Migration Files

Migration files terletak di folder `migrations/` dengan format:
```
000001_create_roles_table.up.sql
000001_create_roles_table.down.sql
```

#### Running Migrations

**Run semua pending migrations:**
```bash
make migrate-up
```

**Rollback last migration:**
```bash
make migrate-down
```

**Check migration version:**
```bash
make migrate-version
```

**Create new migration:**
```bash
make migrate-create NAME=add_new_column_to_users
```

**Force set migration version (jika ada error):**
```bash
make migrate-force VERSION=1
```

#### Migration Tables

Migrations akan membuat tabel-tabel berikut secara berurutan:
1. `roles` - User roles (Customer, Seller, Admin)
2. `users` - User accounts
3. `user_addresses` - Delivery addresses
4. `categories` - Product categories
5. `products` - Product catalog
6. `product_categories` - Many-to-many relationship
7. `product_galleries` - Product images
8. `carts` - Shopping carts
9. `order_statuses` - Order status types
10. `orders` - Customer orders
11. `order_items` - Order details
12. `product_ratings` - Product reviews

## Database Seeders

Project ini memisahkan **migrations** (struktur database) dan **seeders** (data awal). Seeders terletak di folder `seeders/`.

### Running Seeders

**Run all seeders:**
```bash
make seed
```

**What gets seeded:**
1. **Roles** - 3 user roles (Customer, Seller, Admin)
2. **Order Statuses** - 6 status types (Pending, Processing, Shipped, Delivered, Cancelled, Refunded)
3. **Categories** - 20 product categories dengan gambar dari Unsplash

### Available Seeders

| File | Description | Records |
|------|-------------|---------|
| `001_roles.sql` | User roles | 3 roles |
| `002_order_statuses.sql` | Order statuses | 6 statuses |
| `003_categories.sql` | Product categories | 20 categories |

### Categories with Images

Seeder categories sudah include 20 kategori dengan gambar gratis dari [Unsplash](https://unsplash.com/):

- ⚡ Electronics
- 👔 Fashion  
- 🏠 Home & Living
- 💄 Beauty & Health
- ⚽ Sports & Outdoor
- 📚 Books & Stationery
- 🎮 Toys & Games
- 🍔 Food & Beverage
- 🚗 Automotive
- 👶 Baby & Kids
- 🐾 Pet Supplies
- 💼 Office Supplies
- 🌱 Garden & Outdoor
- 🎸 Musical Instruments
- 💎 Jewelry & Accessories
- 🎨 Arts & Crafts
- 🛋️ Furniture
- 💻 Computer & Laptops
- 📱 Mobile Phones
- 📷 Cameras & Photography

Lihat `seeders/README.md` untuk detail lengkap.

## License

MIT
