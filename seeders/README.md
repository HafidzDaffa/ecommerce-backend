# Database Seeders

This directory contains SQL seed files to populate the database with initial data.

## What are Seeders?

Seeders are scripts that populate your database with initial or sample data. Unlike migrations (which create schema structure), seeders insert actual data into tables.

## Seeder Files

Seeders are executed in alphabetical order:

### Master Data Seeders
1. **001_roles.sql** - User roles (Customer, Seller, Admin)
2. **002_order_statuses.sql** - Order status types (Pending, Processing, etc.)
3. **003_categories.sql** - Product categories with images

## Running Seeders

### Using Makefile (Recommended)

```bash
# Run all seeders
make seed

# Or run seeder command directly
make seed-run
```

### Using Go Command

```bash
# Run from backend directory
go run cmd/seeder/main.go
```

## Seeder Structure

Each seeder file contains:
- SQL INSERT statements
- ON CONFLICT DO NOTHING (to allow re-running without errors)
- Sequence reset statements (for auto-increment IDs)

Example:
```sql
INSERT INTO roles (id, slug, name, created_at) VALUES
(1, 'customer', 'Customer', NOW())
ON CONFLICT (id) DO NOTHING;

SELECT setval('roles_id_seq', (SELECT MAX(id) FROM roles));
```

## Categories with Free Images

The categories seeder includes 20 pre-defined categories with high-quality images from [Unsplash](https://unsplash.com/) (free to use):

| ID | Category | Icon | Description |
|----|----------|------|-------------|
| 1 | Electronics | ⚡ | Gadgets, devices, tech products |
| 2 | Fashion | 👔 | Clothing, apparel, accessories |
| 3 | Home & Living | 🏠 | Furniture, decor, kitchenware |
| 4 | Beauty & Health | 💄 | Cosmetics, skincare, wellness |
| 5 | Sports & Outdoor | ⚽ | Fitness, camping, sports gear |
| 6 | Books & Stationery | 📚 | Books, notebooks, pens |
| 7 | Toys & Games | 🎮 | Kids toys, board games, video games |
| 8 | Food & Beverage | 🍔 | Groceries, snacks, drinks |
| 9 | Automotive | 🚗 | Car parts, accessories, tools |
| 10 | Baby & Kids | 👶 | Baby products, kids items |
| 11 | Pet Supplies | 🐾 | Pet food, toys, accessories |
| 12 | Office Supplies | 💼 | Desk items, office equipment |
| 13 | Garden & Outdoor | 🌱 | Plants, gardening tools |
| 14 | Musical Instruments | 🎸 | Guitars, keyboards, drums |
| 15 | Jewelry & Accessories | 💎 | Rings, necklaces, watches |
| 16 | Arts & Crafts | 🎨 | Art supplies, craft materials |
| 17 | Furniture | 🛋️ | Sofas, tables, chairs |
| 18 | Computer & Laptops | 💻 | PCs, laptops, components |
| 19 | Mobile Phones | 📱 | Smartphones, tablets |
| 20 | Cameras & Photography | 📷 | Cameras, lenses, accessories |

## Creating New Seeders

### 1. Create a new SQL file

Follow the naming convention: `{number}_{name}.sql`

```bash
# Example
touch seeders/004_sample_products.sql
```

### 2. Write your SQL

```sql
-- 004_sample_products.sql
INSERT INTO products (id, user_id, product_name, slug, price, stock_quantity) VALUES
(gen_random_uuid(), 'user-uuid-here', 'Sample Product', 'sample-product', 99.99, 100)
ON CONFLICT (id) DO NOTHING;
```

### 3. Run seeders

```bash
make seed
```

## Best Practices

1. **Use ON CONFLICT DO NOTHING** - Allows re-running seeders without errors
2. **Use fixed IDs for master data** - Makes it easier to reference in code
3. **Reset sequences** - Ensure auto-increment continues correctly
4. **Keep it idempotent** - Seeders should be safe to run multiple times
5. **Document your data** - Add comments explaining what data is being seeded

## When to Use Seeders

✅ **Use seeders for:**
- Master data (roles, statuses, categories)
- Sample/demo data for development
- Initial configuration data
- Test data

❌ **Don't use seeders for:**
- Schema changes (use migrations)
- Production user data (security risk)
- Sensitive information (passwords, keys)

## Troubleshooting

### Problem: "Duplicate key violation"
**Solution:** Use `ON CONFLICT DO NOTHING` or `ON CONFLICT (id) DO UPDATE`

### Problem: "Foreign key constraint violation"
**Solution:** Run seeders in correct order (dependencies first)

### Problem: Want to reset data
**Solution:** 
```bash
# Truncate tables and re-seed
make seed-reset
```

## Difference: Migrations vs Seeders

| Aspect | Migrations | Seeders |
|--------|-----------|---------|
| Purpose | Schema structure | Data population |
| When | Initial setup, schema changes | Initial data, sample data |
| Versioning | Sequential, tracked | Can be re-run |
| Production | Always run | Optional (master data only) |
| Location | `migrations/` | `seeders/` |

## Image Attribution

All category images are sourced from [Unsplash](https://unsplash.com/), a free stock photo platform. Images are free to use for commercial and non-commercial purposes under the [Unsplash License](https://unsplash.com/license).

If you want to use different images, you can:
1. Replace URLs in `003_categories.sql`
2. Use your own hosted images
3. Use other free image sources (Pexels, Pixabay, etc.)
