# 🔄 Cara Update Swagger Documentation

## ❓ Kenapa Swagger UI Tidak Menampilkan Semua Endpoint?

Jika Swagger UI hanya menampilkan endpoint Authentication saja, ikuti langkah-langkah berikut:

---

## ✅ Solusi 1: Regenerate Swagger Docs (Paling Mudah)

### Menggunakan Makefile:
```bash
make swagger
```

### Atau menggunakan script:
```bash
./scripts/regenerate-swagger.sh
```

### Atau manual:
```bash
~/go/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -d ./,./internal/adapters/primary/http
```

---

## ✅ Solusi 2: Hard Refresh Browser

Setelah regenerate swagger docs:

### Chrome / Edge / Firefox:
- **Windows/Linux:** `Ctrl + Shift + R`
- **Mac:** `Cmd + Shift + R`

### Atau Clear Cache:
1. Buka Developer Tools: `F12`
2. Right-click pada refresh button
3. Pilih "Empty Cache and Hard Reload"

---

## ✅ Solusi 3: Restart Server

```bash
# Stop server (Ctrl+C)
# Kemudian start lagi:
go run cmd/api/main.go

# Atau dengan Makefile:
make run
```

---

## 📋 Langkah-Langkah Lengkap

### 1. Install swag tool (jika belum):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Generate Swagger docs:
```bash
make swagger
```

Atau:
```bash
~/go/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -d ./,./internal/adapters/primary/http
```

**Output yang benar:**
```
✓ Generating ecommerce-backend_internal_core_domain.CreateProductRequest
✓ Generating ecommerce-backend_internal_core_domain.UpdateProductRequest
✓ warning: route POST /categories is declared multiple times
✓ warning: route GET /products is declared multiple times
✓ create docs.go at docs/docs.go
✓ create swagger.json at docs/swagger.json
✓ create swagger.yaml at docs/swagger.yaml
```

### 3. Verify endpoints:
```bash
jq '.paths | keys' docs/swagger.json
```

**Expected output:**
```json
[
  "/auth/login",
  "/auth/logout",
  "/auth/me",
  "/auth/register",
  "/categories",
  "/categories/{id}",
  "/products",
  "/products/category/{category_id}",
  "/products/galleries",
  "/products/galleries/{id}",
  "/products/slug/{slug}",
  "/products/{id}",
  "/products/{product_id}/galleries"
]
```

### 4. Start/Restart server:
```bash
go run cmd/api/main.go
```

### 5. Open Swagger UI:
```
http://localhost:8080/swagger/index.html
```

### 6. Hard refresh browser:
- **Windows/Linux:** `Ctrl + Shift + R`
- **Mac:** `Cmd + Shift + R`

---

## 🎯 Verifikasi Sukses

Setelah mengikuti langkah di atas, Swagger UI harus menampilkan:

### ✅ Tags (Sections):
- **Authentication** (4 endpoints)
- **Categories** (5 endpoints)
- **Products** (7 endpoints)
- **Product Galleries** (4 endpoints)

### ✅ Total: 20 endpoints

---

## 🔍 Troubleshooting

### Problem 1: "swag: command not found"
**Solution:**
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Problem 2: "No endpoints showing in Swagger UI"
**Solutions:**
1. Check if `docs/swagger.json` exists and has content
2. Check browser console for errors (F12)
3. Verify server is running
4. Try different browser
5. Clear browser cache completely

### Problem 3: "Old endpoints still showing"
**Solutions:**
1. Hard refresh: `Ctrl + Shift + R`
2. Restart server
3. Check timestamp of `docs/swagger.json`:
   ```bash
   ls -lh docs/swagger.json
   ```
4. Regenerate docs again

### Problem 4: "Categories/Products not in Swagger"
**Check if handlers have godoc comments:**
```bash
# Should show swagger annotations
grep -n "@Summary" internal/adapters/primary/http/category_handler.go
grep -n "@Summary" internal/adapters/primary/http/product_handler.go
```

If no output, handlers missing swagger annotations.

---

## 📝 Best Practices

### 1. Always regenerate after:
- Adding new endpoints
- Modifying handler signatures
- Changing request/response structures
- Updating API documentation

### 2. Use Makefile commands:
```bash
make swagger    # Generate docs
make run        # Start server
make build      # Build application
```

### 3. Verify before commit:
```bash
# Check docs were generated
ls -lh docs/

# Verify endpoint count
jq '.paths | keys | length' docs/swagger.json

# Should show: 13 (for current API)
```

### 4. Add to git pre-commit (optional):
```bash
# Create .git/hooks/pre-commit
#!/bin/bash
make swagger
git add docs/
```

---

## 🔄 Automatic Regeneration

### Option 1: Watch mode (requires air)
```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Option 2: Pre-commit hook
```bash
# Create/edit .git/hooks/pre-commit
#!/bin/bash
echo "Regenerating Swagger docs..."
make swagger
git add docs/
```

---

## 📚 Quick Commands Reference

```bash
# Regenerate Swagger
make swagger

# Run server
make run

# Build app
make build

# View endpoints
jq '.paths | keys' docs/swagger.json

# Count endpoints
jq '.paths | keys | length' docs/swagger.json

# View specific endpoint
jq '.paths."/categories"' docs/swagger.json

# Check server
curl http://localhost:8080/api/health

# Open Swagger UI
open http://localhost:8080/swagger/index.html
```

---

## 🎓 Understanding Swagger Generation

### What `swag init` does:
1. Scans `cmd/api/main.go` for general API info
2. Scans specified directories for godoc comments
3. Parses `@Summary`, `@Description`, `@Tags`, etc.
4. Generates OpenAPI spec in `docs/swagger.json`
5. Creates Go bindings in `docs/docs.go`

### Flags explanation:
- `-g cmd/api/main.go` - General API info location
- `-o docs` - Output directory
- `--parseDependency` - Parse vendor dependencies
- `--parseInternal` - Parse internal packages
- `-d ./,./internal/adapters/primary/http` - Directories to scan

### Why need `-d` flag:
- Default: only scans `cmd/api/`
- Handlers are in `internal/adapters/primary/http/`
- Need explicit path to find all godoc comments

---

## ✅ Current API Endpoints (Should All Appear)

### Authentication (4)
- POST `/auth/register`
- POST `/auth/login`
- GET `/auth/me`
- POST `/auth/logout`

### Categories (5)
- GET `/categories`
- GET `/categories/{id}`
- POST `/categories`
- PUT `/categories/{id}`
- DELETE `/categories/{id}`

### Products (7)
- GET `/products`
- GET `/products/{id}`
- GET `/products/slug/{slug}`
- GET `/products/category/{category_id}`
- POST `/products`
- PUT `/products/{id}`
- DELETE `/products/{id}`

### Product Galleries (4)
- GET `/products/{product_id}/galleries`
- POST `/products/galleries`
- PUT `/products/galleries/{id}`
- DELETE `/products/galleries/{id}`

**Total: 20 endpoints**

---

## 🆘 Still Not Working?

1. **Check server logs:**
   ```bash
   tail -f backend.log
   ```

2. **Test endpoint directly:**
   ```bash
   curl http://localhost:8080/api/v1/categories
   ```

3. **Verify swagger.json:**
   ```bash
   cat docs/swagger.json | jq '.paths | keys'
   ```

4. **Check file permissions:**
   ```bash
   ls -la docs/
   ```

5. **Try different port:**
   ```bash
   APP_PORT=8081 go run cmd/api/main.go
   # Access: http://localhost:8081/swagger/index.html
   ```

---

## 📞 Support

Jika masih ada masalah:
1. Check documentation: `SWAGGER_POSTMAN_GUIDE.md`
2. Verify files exist: `ls -la docs/`
3. Test with Postman: `E-Commerce_Complete_API.postman_collection.json`
4. Check GitHub issues

---

**Happy Documentation! 📚**
