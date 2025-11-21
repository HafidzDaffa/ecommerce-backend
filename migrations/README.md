# Database Migrations

This directory contains all database migration files for the e-commerce application.

## Migration Tool

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Install

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Migration Files

Each migration consists of two files:
- `{version}_{name}.up.sql` - Apply the migration
- `{version}_{name}.down.sql` - Rollback the migration

### Current Migrations

1. **000001_create_roles_table** - Create roles table for user roles
2. **000002_create_users_table** - Create users table
3. **000003_create_user_addresses_table** - Create user addresses table
4. **000004_create_categories_table** - Create product categories table
5. **000005_create_products_table** - Create products table
6. **000006_create_product_categories_table** - Create many-to-many relationship table
7. **000007_create_product_galleries_table** - Create product galleries table
8. **000008_create_carts_table** - Create shopping carts table
9. **000009_create_order_statuses_table** - Create order statuses table
10. **000010_create_orders_table** - Create orders table
11. **000011_create_order_items_table** - Create order items table
12. **000012_create_product_ratings_table** - Create product ratings table
13. **000013_seed_roles_and_order_statuses** - Seed master data

## Running Migrations

### Using Makefile (Recommended)

```bash
# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Check current version
make migrate-version

# Create new migration
make migrate-create NAME=add_new_feature

# Force set version (use with caution)
make migrate-force VERSION=1
```

### Using migrate CLI directly

```bash
# Set database URL
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/ecommerce_db?sslmode=disable"

# Run migrations
migrate -path migrations -database $DATABASE_URL up

# Rollback
migrate -path migrations -database $DATABASE_URL down 1

# Go to specific version
migrate -path migrations -database $DATABASE_URL goto 5

# Force version (if dirty)
migrate -path migrations -database $DATABASE_URL force 1
```

## Creating New Migrations

### Using Makefile

```bash
make migrate-create NAME=add_email_verification_token
```

### Using migrate CLI

```bash
migrate create -ext sql -dir migrations -seq add_email_verification_token
```

This will create two files:
- `000014_add_email_verification_token.up.sql`
- `000014_add_email_verification_token.down.sql`

## Migration Best Practices

1. **Always test migrations** - Test both up and down migrations in development
2. **Keep migrations small** - One logical change per migration
3. **Write reversible migrations** - Always write corresponding down migrations
4. **Use transactions** - PostgreSQL migrations are transactional by default
5. **Don't modify old migrations** - Create new migrations to fix issues
6. **Backup before migrating** - Always backup production database before running migrations

## Troubleshooting

### Dirty Database State

If migration fails midway, database might be in "dirty" state:

```bash
# Check version
make migrate-version

# Force to last known good version
make migrate-force VERSION=5
```

### Reset All Migrations

⚠️ **WARNING: This will delete all data!**

```bash
# Drop all migrations
migrate -path migrations -database $DATABASE_URL drop

# Run all migrations again
make migrate-up
```

## Schema Diagram

See `dbml.dbml` file in the backend root directory for complete database schema diagram.
