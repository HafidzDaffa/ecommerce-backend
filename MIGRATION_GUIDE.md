# Quick Migration Guide

## Setup (One-time)

1. **Install golang-migrate CLI:**
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

2. **Verify installation:**
   ```bash
   migrate -version
   ```

## Daily Usage

### 1. Start Database (Docker)
```bash
make docker-up
```

### 2. Run Migrations
```bash
make migrate-up
```

### 3. Run Seeders
```bash
make seed
```

This will populate:
- 3 roles (Customer, Seller, Admin)
- 6 order statuses
- 20 product categories with images

### 4. Verify Migration
```bash
make migrate-version
```

Expected output: `12` (latest version)

### 5. Start Application
```bash
make run
```

## Common Commands

| Command | Description |
|---------|-------------|
| `make migrate-up` | Run all pending migrations |
| `make migrate-down` | Rollback last migration |
| `make migrate-version` | Check current version |
| `make migrate-create NAME=xxx` | Create new migration |
| `make seed` | Run all database seeders |

## Database Connection

Default connection (when using Docker):
```
Host: localhost
Port: 5432
Database: ecommerce_db
User: postgres
Password: postgres
```

## What Gets Created?

### After `make migrate-up` - Tables (12 total):
1. ✅ `roles` - User roles
2. ✅ `users` - User accounts  
3. ✅ `user_addresses` - Delivery addresses
4. ✅ `categories` - Product categories
5. ✅ `products` - Product catalog
6. ✅ `product_categories` - Product-category mapping
7. ✅ `product_galleries` - Product images
8. ✅ `carts` - Shopping carts
9. ✅ `order_statuses` - Order status types
10. ✅ `orders` - Customer orders
11. ✅ `order_items` - Order line items
12. ✅ `product_ratings` - Product reviews

### After `make seed` - Initial Data:

**Roles (3 records):**
- Customer (default for new users)
- Seller (for vendors)
- Admin (for administrators)

**Order Statuses (6 records):**
- Pending
- Processing
- Shipped
- Delivered
- Cancelled
- Refunded

**Categories (20 records):**
- Electronics, Fashion, Home & Living, Beauty, Sports, Books, Toys, Food, etc.
- All with emoji icons and free images from Unsplash

## Troubleshooting

### Problem: "migrate: command not found"
**Solution:** Install migrate CLI tool
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Problem: "error: Dirty database version"
**Solution:** Force to last known good version
```bash
make migrate-force VERSION=13
```

### Problem: "connection refused"
**Solution:** Start PostgreSQL
```bash
make docker-up
# Wait 10 seconds for PostgreSQL to start
make migrate-up
```

### Problem: Want to start fresh
**Solution:** Drop database and recreate
```bash
make docker-down
make docker-up
make migrate-up
```

## Quick Test

After migration, verify with psql:

```bash
# Connect to database
docker exec -it ecommerce_postgres psql -U postgres -d ecommerce_db

# Check tables
\dt

# Check roles
SELECT * FROM roles;

# Check order statuses
SELECT * FROM order_statuses;

# Exit
\q
```

Expected: You should see 12 tables + 1 schema_migrations table.

## Next Steps

After migrations are successful:
1. ✅ Tables are ready
2. 🔄 Create domain entities in `internal/core/domain/`
3. 🔄 Create repositories in `internal/adapters/secondary/repository/`
4. 🔄 Create services in `internal/core/services/`
5. 🔄 Create HTTP handlers in `internal/adapters/primary/http/`

Happy coding! 🚀
