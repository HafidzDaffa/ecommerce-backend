# Database Migrations & Seeders

## Overview

Backend ini menggunakan **file-based migrations** dan **seeders** yang dapat dijalankan secara manual. Setiap perubahan skema database dibuat sebagai file migration terpisah.

## Project Structure

```
backend/
├── migrations/                    # SQL migration files
│   ├── 000001_create_tables.up.sql
│   ├── 000001_create_tables.down.sql
│   └── ...
├── seeders/                      # Database seeder files
│   ├── seeders.go                # Main seeder runner
│   ├── role_seeder.go            # Role seeder
│   └── ...
├── cmd/
│   └── migrate.go                # Migration & seeder CLI
└── scripts/
    └── create_migration.sh        # Helper script to create new migration
```

## Migration File Format

Each migration consists of two files:
- `VERSION_migration_name.up.sql` - SQL for applying migration
- `VERSION_migration_name.down.sql` - SQL for rolling back migration

Each table should have its own migration file for better tracking.

Example:
```
migrations/
├── 000001_create_roles_table.up.sql
├── 000001_create_roles_table.down.sql
├── 000002_create_users_table.up.sql
└── 000002_create_users_table.down.sql
```

**Migration Numbers:**
- Sequential numbering (001, 002, 003, ...)
- Each table/feature gets its own number
- Run order is determined by number (001 first, then 002, etc.)
migrations/
├── 000001_create_tables.up.sql
└── 000001_create_tables.down.sql
```

## Available Commands

### Migrate Up (Apply migrations)

```bash
# Run all pending migrations
make migrate

# Or run directly
go run cmd/migrate.go migrate
```

### Migrate Down (Rollback)

```bash
# Rollback last migration
make migrate-down

# Or run directly
migrate -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" -path migrations down 1
```

### Refresh (Rollback all and migrate again)

```bash
make migrate-refresh
```

### Seed Database

```bash
make seed

# Or run directly
go run cmd/migrate.go seed
```

## Creating New Migration

### Option 1: Using Helper Script

```bash
./scripts/create_migration.sh add_products_table
```

This creates:
```
migrations/
├── 20240106100000_add_products_table.up.sql
└── 20240106100000_add_products_table.down.sql
```

### Option 2: Manual Creation

1. Create two files in `migrations/` directory:
   - `VERSION_migration_name.up.sql`
   - `VERSION_migration_name.down.sql`

2. Find the next migration number:
   - Count existing `.up.sql` files in `migrations/`
   - Add 1 to get the next number

3. Write SQL in UP file:
```sql
-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

3. Write rollback SQL in DOWN file:
```sql
-- +migrate Down
DROP TABLE products;
```

## Seeders

### Available Seeders

- **RoleSeeder**: Seeds default roles (customer, toko, admin)

### Creating New Seeder

1. Create new file in `seeders/` directory:
```go
package seeders

import "github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"

type YourSeeder struct{}

func (s *YourSeeder) Seed() error {
    // Your seeder logic here
    return nil
}
```

2. Add to `seeders/seeders.go`:
```go
func RunSeeders() error {
    seeders := []Seeder{
        &RoleSeeder{},
        &YourSeeder{},  // Add new seeder here
    }

    // ... rest of the code
}
```

## Workflow Example

### 1. Create Roles Table (Already Done)

```bash
# Migration files created manually:
migrations/
├── 000001_create_roles_table.up.sql
├── 000001_create_roles_table.down.sql
```

### 2. Create Users Table (Already Done)

```bash
# Migration files created manually:
migrations/
├── 000002_create_users_table.up.sql
├── 000002_create_users_table.down.sql
```

### 3. Create New Table

```bash
# Create migration file using helper script
./scripts/create_migration.sh create_products_table

# Edit migrations/000003_create_products_table.up.sql
# Add your CREATE TABLE SQL

# Edit migrations/000003_create_products_table.down.sql
# Add your DROP TABLE SQL

# Run migration
make migrate

# Check status
make migrate-status
```

### 2. Add Initial Data

```bash
# Create seeder file
# Add seeder logic

# Run seeders
make seed
```

### 3. Add Column to Existing Table

```bash
# Create migration for column addition
./scripts/create_migration.sh add_status_column_to_users

# Edit migrations/000003_add_status_column_to_users.up.sql
# ALTER TABLE users ADD COLUMN status VARCHAR(50);

# Edit migrations/000003_add_status_column_to_users.down.sql
# ALTER TABLE users DROP COLUMN IF EXISTS status;

# Run migration
make migrate

# Check status
make migrate-status
```

### 4. Development Flow

```bash
# Start PostgreSQL
make db-run

# Apply migrations
make migrate

# Run seeders
make seed

# Make changes, test, repeat

# Check migration status
make migrate-status

# If needed, rollback last migration
make migrate-down

# Or refresh everything
make migrate-refresh
```

## Best Practices

1. **One Table Per Migration**: Each table should have its own migration file
   - Example: `000001_create_roles_table`, `000002_create_users_table`
   - Makes tracking changes easier
   - Enables selective rollbacks

2. **Migration Numbers**: Use sequential numbering (001, 002, 003, ...)
   - Determines execution order
   - Easier to track dependencies

3. **Naming**: Use descriptive names
   - `create_table_name` - for new tables
   - `add_column_name_to_table` - for new columns
   - `alter_table` - for table modifications
   - `drop_table_name` - for dropping tables

4. **Dependencies**: Order migrations correctly
   - Create parent tables first (e.g., roles before users)
   - Create indexes after tables

5. **Rollback**: Always write a DOWN file
   - Ensure you can rollback changes
   - Test both UP and DOWN migrations

6. **Idempotent**: Seeders should check if data already exists
7. **Testing**: Test migrations on development database first
8. **Backups**: Backup production database before running migrations

## Role Seeder

Default roles seeded:
- **ID 1**: customer - Customer dengan akses belanja standar
- **ID 2**: toko - Toko dengan akses manajemen produk
- **ID 3**: admin - Administrator dengan akses penuh ke sistem

## Troubleshooting

### Check Migration Status

Always check migration status before running new migrations:
```bash
make migrate-status
```

This shows:
- Current migration version applied
- All available migration files
- Which migrations are pending

### Migration fails

1. Check database connection in `.env`
2. Verify migration file syntax
3. Check if previous migrations are complete (`make migrate-status`)
4. Use `make db-shell` to inspect database

### Seeder fails

1. Ensure migrations have been run first (`make migrate`)
2. Check if data already exists (seeders are idempotent)
3. Verify foreign key constraints

### Migration state

Check current migration version:
```sql
SELECT * FROM schema_migrations;
```

Check all tables:
```bash
make db-shell
\dt
```

## Environment Variables

For migrations and seeders, ensure these are set in `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=ecommerce
DB_PASSWORD=ecommerce123
DB_NAME=ecommerce
```
