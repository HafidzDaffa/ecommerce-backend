#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./scripts/create_migration.sh migration_name"
    echo "Example: ./scripts/create_migration.sh create_products_table"
    echo ""
    echo "Tips:"
    echo "  - Use 'create_' prefix for new tables"
    echo "  - Use 'add_' prefix for new columns"
    echo "  - Use 'alter_' prefix for table modifications"
    echo "  - Use 'drop_' prefix for dropping tables/columns"
    exit 1
fi

MIGRATION_NAME=$1

if [ -z "$2" ]; then
    # Auto-generate next migration number
    COUNT=$(ls migrations/*.up.sql 2>/dev/null | wc -l)
    NEXT_NUM=$((COUNT + 1))
    VERSION=$(printf "%06d" $NEXT_NUM)
else
    VERSION=$2
fi

FILENAME="${VERSION}_${MIGRATION_NAME}"

mkdir -p migrations

cat > "migrations/${FILENAME}.up.sql" <<EOF
-- +migrate Up
-- Write your UP migration SQL here
-- Example:
-- CREATE TABLE table_name (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL
-- );

EOF

cat > "migrations/${FILENAME}.down.sql" <<EOF
-- +migrate Down
-- Write your DOWN migration SQL here
-- Example:
-- DROP TABLE IF EXISTS table_name;

EOF

echo "✓ Migration files created:"
echo "  migrations/${FILENAME}.up.sql"
echo "  migrations/${FILENAME}.down.sql"
echo ""
echo "Next migration number: $(($(ls migrations/*.up.sql 2>/dev/null | wc -l) + 1))"
EOF
