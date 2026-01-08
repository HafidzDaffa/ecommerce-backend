# Migration Structure

## Current Migrations

```
migrations/
├── 000001_create_roles_table.up.sql     # Migration 001: Create roles table
├── 000001_create_roles_table.down.sql   # Migration 001: Rollback roles
├── 000002_create_users_table.up.sql     # Migration 002: Create users table
└── 000002_create_users_table.down.sql   # Migration 002: Rollback users
```

## Migration Execution Order

Migrations are executed in order by version number:

```
001. create_roles_table (runs first)
     ↓
002. create_users_table (runs second, depends on roles)
     ↓
003. (next migration)
     ↓
...
```

## Version Tracking

The `schema_migrations` table tracks which migrations have been applied:

```sql
SELECT * FROM schema_migrations;
```

Output:
```
 version | dirty 
---------+-------
       2 | f
```

- **version**: Latest migration number applied
- **dirty**: `f` (false) if migrations completed successfully

## Migration State Examples

### Fresh Database
```
version: 0
dirty: f
Available: 001, 002, 003
Applied: none
Pending: 001, 002, 003
```

### After Running Migrations
```
version: 3
dirty: f
Available: 001, 002, 003
Applied: 001, 002, 003
Pending: none
```

### After Partial Rollback
```
version: 2
dirty: f
Available: 001, 002, 003
Applied: 001, 002
Pending: 003
Rolled back: 003
```

## Why Separate Migration Files?

### ✅ Benefits

1. **Granular Tracking**: Know exactly when each table was created
2. **Selective Rollback**: Rollback specific changes without affecting others
3. **Dependency Management**: Clear order of table creation
4. **Team Collaboration**: Easier to review and merge changes
5. **Debugging**: Identify which migration caused issues

### ❌ vs Single Large Migration File

**Before (Single File):**
```
migrations/
└── 000001_create_tables.up.sql    # All tables in one file
```

**Problems:**
- Cannot track individual table creation
- Hard to rollback specific table
- Confusing merge conflicts
- Difficult to debug which table failed

**After (Separate Files):**
```
migrations/
├── 000001_create_roles_table.up.sql    # Table 001
├── 000002_create_users_table.up.sql    # Table 002
└── 000003_create_products_table.up.sql # Table 003
```

**Benefits:**
- Clear history of each table
- Can rollback individual migrations
- Easier to understand changes
- Better team collaboration
