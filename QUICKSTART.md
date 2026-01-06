# Quick Start Guide

## Option 1: Using Docker (Recommended for Production)

### Start Everything (Backend + PostgreSQL)

```bash
make build
make run
```

Access the API at: http://localhost:8080

### Start Only PostgreSQL (for Local Development)

```bash
make db-run
```

Then run the backend locally:

```bash
cp .env.example .env
go mod download
go run main.go
```

## Option 2: Local Development Without Docker

### 1. Install PostgreSQL Locally

```bash
# On Ubuntu/Debian
sudo apt-get install postgresql postgresql-contrib

# On macOS with Homebrew
brew install postgresql
brew services start postgresql
```

### 2. Create Database

```bash
sudo -u postgres psql
CREATE DATABASE ecommerce;
CREATE USER ecommerce WITH PASSWORD 'ecommerce123';
GRANT ALL PRIVILEGES ON DATABASE ecommerce TO ecommerce;
\q
```

### 3. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` if needed:

```env
APP_ENV=development
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=ecommerce
DB_PASSWORD=ecommerce123
DB_NAME=ecommerce
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
```

### 4. Run the Application

```bash
go mod download
go run main.go
```

Or build and run:

```bash
go build -o bin/main .
./bin/main
```

## Test the API

```bash
# Health Check
curl http://localhost:8080/api/v1/health

# Register User
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

## Database Management

### Connect to PostgreSQL (Docker)

```bash
make db-shell
```

### Connect to PostgreSQL (Local)

```bash
psql -U ecommerce -d ecommerce
```

### View Tables

```sql
\dt
```

### Query Users

```sql
SELECT * FROM users;
```

### Query Roles

```sql
SELECT * FROM roles;
```

## Troubleshooting

### Port Already in Use

Change the port in `.env`:

```env
PORT=8081
```

### Database Connection Failed

1. Check if PostgreSQL is running:
   ```bash
   # Docker
   docker ps | grep postgres

   # Local
   sudo systemctl status postgresql
   ```

2. Verify connection settings in `.env`

3. Test connection:
   ```bash
   psql -h localhost -U ecommerce -d ecommerce
   ```

### Build Errors

```bash
# Clean and rebuild
go clean -modcache
go mod download
go build -o bin/main .
```

### Docker Issues

```bash
# Reset everything
make clean
make build
make run
```
