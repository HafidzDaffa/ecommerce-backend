# Migration Summary

## ✅ Changes Made

Migration files have been separated into individual files for each table to enable better tracking.

## 📁 New Migration Structure

```
migrations/
├── 000001_create_roles_table.up.sql      # Migration 001: Create roles table
├── 000001_create_roles_table.down.sql    # Migration 001: Rollback roles
├── 000002_create_users_table.up.sql      # Migration 002: Create users table
└── 000002_create_users_table.down.sql    # Migration 002: Rollback users
```

## 🔄 Migration Order

1. **001 - create_roles_table** (runs first)
   - Creates `roles` table
   - Creates index on `name` column

2. **002 - create_users_table** (runs second)
   - Creates `gender_enum` type
   - Creates `users` table with foreign key to `roles`
   - Creates indexes on users table

## 📊 Migration Tracking

### Check Status
```bash
make migrate-status
```

Shows:
- Current migration version applied
- All available migration files

### Apply Migrations
```bash
make migrate
```

Output:
```
1/u create_roles_table (23ms)
2/u create_users_table (62ms)
```

### Rollback
```bash
make migrate-down
```

Rolls back the last migration (e.g., users table)

### Refresh
```bash
make migrate-refresh
```

Rolls back all migrations and runs them again

## 🎯 Benefits of Separate Migration Files

### Granular Control
- Track each table independently
- Rollback specific migrations without affecting others
- Clear history of changes

### Better Collaboration
- Easier to merge changes
- Reduce conflicts in team development
- Clear ownership of each migration

### Improved Debugging
- Quickly identify which migration failed
- Test migrations individually
- Easier to understand impact

## 📋 Available Commands

```bash
make migrate            # Apply pending migrations
make migrate-down      # Rollback last migration
make migrate-status     # Show migration version and available files
make migrate-refresh   # Rollback all and migrate again
make seed             # Run database seeders
```

## 📖 Documentation

- `MIGRATIONS.md` - Complete migration & seeder guide
- `MIGRATION_STRUCTURE.md` - Visual migration structure and benefits
- `README.md` - Project overview with migration commands

## 🧪 Example: Create New Migration

```bash
# Create migration for products table
./scripts/create_migration.sh create_products_table

# This creates:
# migrations/000003_create_products_table.up.sql
# migrations/000003_create_products_table.down.sql

# Edit the UP file
# migrations/000003_create_products_table.up.sql:
-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

# Edit the DOWN file
# migrations/000003_create_products_table.down.sql:
-- +migrate Down
DROP TABLE IF EXISTS products;

# Run the migration
make migrate

# Check status
make migrate-status
```

## 🧪 Example: Add Column to Table

```bash
# Create migration for column addition
./scripts/create_migration.sh add_status_to_users

# This creates:
# migrations/000003_add_status_to_users.up.sql
# migrations/000003_add_status_to_users.down.sql

# Edit the UP file
-- +migrate Up
ALTER TABLE users ADD COLUMN status VARCHAR(50) DEFAULT 'active';

# Edit the DOWN file
-- +migrate Down
ALTER TABLE users DROP COLUMN IF EXISTS status;

# Run the migration
make migrate
```

## 🔍 Troubleshooting

### Check Which Migrations Are Applied

```bash
make db-shell
```

```sql
SELECT * FROM schema_migrations;
```

### Check All Tables

```bash
make db-shell
```

```sql
\dt
```

### Rollback to Specific Version

```bash
# Rollback to version 1 (only roles table)
./migrate -database "postgres://ecommerce:ecommerce123@localhost:5432/ecommerce?sslmode=disable" -path migrations force 1
```

## ✅ Testing

All migrations have been tested successfully:

```bash
✓ Fresh database setup
✓ Migration 001: create_roles_table
✓ Migration 002: create_users_table
✓ Seeder: roles (customer, toko, admin)
✓ Rollback: migration 002
✓ Re-migrate: all migrations
```

## 📝 Best Practices

1. **One Table Per Migration**: Create separate migration for each table
2. **Sequential Numbers**: Use 001, 002, 003 for versioning
3. **Dependencies**: Create parent tables first (roles before users)
4. **Descriptive Names**: Use `create_table_name` for clarity
5. **Always Rollback**: Test DOWN migrations to ensure they work
6. **Check Status**: Always run `make migrate-status` before creating new migrations

## 🎉 Success!

Migration files are now separate and can be tracked individually! Each table change has its own migration file with clear versioning.
