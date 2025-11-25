# 🔧 CARA FIX SWAGGER UI - Menampilkan Semua Endpoint

## ⚠️ Masalah: Swagger UI Hanya Menampilkan Authentication

Ikuti langkah-langkah berikut **SECARA BERURUTAN**:

---

## 📝 LANGKAH 1: Regenerate Swagger Documentation

Swagger docs sudah saya generate ulang dengan benar. Verifikasi:

```bash
# Check jumlah endpoints (harus 13)
jq '.paths | keys | length' docs/swagger.json
```

Expected output: `13`

```bash
# Lihat semua endpoints
jq '.paths | keys' docs/swagger.json
```

Expected output harus termasuk:
- `/categories`
- `/products`
- `/products/galleries`

---

## 🔄 LANGKAH 2: Restart Server (PENTING!)

**Jika server sedang running, STOP dan START ulang:**

```bash
# 1. Stop server dengan Ctrl+C

# 2. Start ulang
go run cmd/api/main.go

# Atau dengan Makefile
make run
```

**Tunggu sampai muncul:**
```
Server is running on port 8080
```

---

## 🌐 LANGKAH 3: Hard Refresh Browser

**Buka Swagger UI:**
```
http://localhost:8080/swagger/index.html
```

**Hard Refresh (PENTING!):**
- **Windows/Linux:** Tekan `Ctrl + Shift + R`
- **Mac:** Tekan `Cmd + Shift + R`

**Atau Clear Cache:**
1. Tekan `F12` (Developer Tools)
2. Right-click tombol refresh
3. Pilih **"Empty Cache and Hard Reload"**

---

## ✅ VERIFIKASI: Yang Harus Muncul

Setelah hard refresh, Swagger UI harus menampilkan **4 SECTIONS**:

### 1. Authentication ✅
- POST `/auth/register`
- POST `/auth/login`
- GET `/auth/me`
- POST `/auth/logout`

### 2. Categories ✅
- GET `/categories`
- GET `/categories/{id}`
- POST `/categories`
- PUT `/categories/{id}`
- DELETE `/categories/{id}`

### 3. Products ✅
- GET `/products`
- GET `/products/{id}`
- GET `/products/slug/{slug}`
- GET `/products/category/{category_id}`
- POST `/products`
- PUT `/products/{id}`
- DELETE `/products/{id}`

### 4. Product Galleries ✅
- GET `/products/{product_id}/galleries`
- POST `/products/galleries`
- PUT `/products/galleries/{id}`
- DELETE `/products/galleries/{id}`

**Total: 20 endpoints di 4 sections**

---

## 🚨 Jika Masih Belum Muncul

### Option 1: Clear Browser Cache Completely
1. Chrome: `Settings` → `Privacy and security` → `Clear browsing data`
2. Pilih `Cached images and files`
3. Click `Clear data`
4. Restart browser
5. Buka kembali `http://localhost:8080/swagger/index.html`

### Option 2: Try Incognito/Private Window
```
1. Buka browser incognito/private window
2. Akses: http://localhost:8080/swagger/index.html
3. Jika muncul semua → masalah di cache
```

### Option 3: Try Different Browser
- Chrome
- Firefox
- Edge
- Safari

### Option 4: Check Console Errors
1. Tekan `F12` (Developer Tools)
2. Buka tab `Console`
3. Refresh halaman
4. Cek ada error atau tidak

### Option 5: Regenerate Swagger Manually

Jika semua cara di atas gagal, regenerate ulang:

```bash
# Masuk ke directory backend
cd /home/hafidz/Documents/ecommerce-go-vue/backend

# Regenerate swagger
~/go/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -d ./,./internal/adapters/primary/http

# Check hasil
jq '.paths | keys' docs/swagger.json

# Restart server
go run cmd/api/main.go
```

---

## 🎯 Quick Test

Test manual bahwa endpoint Categories dan Products bekerja:

```bash
# Test Categories (harus return JSON)
curl http://localhost:8080/api/v1/categories

# Test Products (harus return JSON)
curl http://localhost:8080/api/v1/products

# Test Health Check
curl http://localhost:8080/api/health
```

Jika semua curl command return JSON dengan benar, berarti:
- ✅ Backend bekerja
- ✅ Endpoints sudah terdaftar
- ⚠️ Masalah di Swagger UI (cache browser)

---

## 💡 Alternative: Use Postman Instead

Jika Swagger UI tetap bermasalah, gunakan Postman:

```bash
# 1. Import Postman Collection
File: E-Commerce_Complete_API.postman_collection.json

# 2. Import Environment
File: E-Commerce_API.postman_environment.json

# 3. Select Environment
Dropdown (top right) → "E-Commerce API Environment"

# 4. Test semua endpoint
✅ Postman tidak punya cache issue seperti browser
```

---

## 📊 Debug Information

Check swagger.json content:

```bash
# Total paths
jq '.paths | keys | length' docs/swagger.json
# Output harus: 13

# All paths
jq '.paths | keys' docs/swagger.json
# Output harus include categories, products, galleries

# Specific category endpoint
jq '.paths."/categories"' docs/swagger.json
# Output harus ada GET method

# Specific product endpoint
jq '.paths."/products"' docs/swagger.json
# Output harus ada GET method
```

---

## 🎓 Understanding the Issue

**Kenapa terjadi?**
1. Swagger docs awalnya di-generate tanpa scan handler files
2. Handler files di `internal/adapters/primary/http/` tidak ter-scan
3. Perlu flag `-d` untuk specify directories yang harus di-scan

**Sudah di-fix dengan:**
```bash
swag init -g cmd/api/main.go -o docs \
  --parseDependency --parseInternal \
  -d ./,./internal/adapters/primary/http
```

Flag `-d ./,./internal/adapters/primary/http` memastikan swag scan:
- `./` (root directory - untuk main.go)
- `./internal/adapters/primary/http` (handler files)

---

## 🔄 Future: Automatic Regeneration

Untuk mencegah masalah ini di future, gunakan:

### Option 1: Makefile
```bash
make swagger
```

### Option 2: Script
```bash
./scripts/regenerate-swagger.sh
```

### Option 3: Pre-commit Hook
```bash
# File: .git/hooks/pre-commit
#!/bin/bash
make swagger
git add docs/
```

---

## ✅ Summary

**Yang sudah dilakukan:**
1. ✅ Swagger docs di-regenerate dengan semua endpoints
2. ✅ File `docs/swagger.json` sudah ada 13 paths
3. ✅ Makefile created untuk easy regeneration
4. ✅ Scripts created untuk automation

**Yang perlu ANDA lakukan:**
1. ⚠️ **RESTART SERVER** (Ctrl+C, lalu start ulang)
2. ⚠️ **HARD REFRESH BROWSER** (Ctrl+Shift+R)
3. ⚠️ Clear browser cache jika perlu

**Expected result:**
- Swagger UI menampilkan 4 sections (Authentication, Categories, Products, Product Galleries)
- Total 20 endpoints visible
- Semua endpoints bisa di-test

---

## 📞 Still Need Help?

1. Check `docs/swagger.json` exists dan tidak kosong
2. Check server running: `curl http://localhost:8080/api/health`
3. Check endpoints work: `curl http://localhost:8080/api/v1/categories`
4. Try Postman as alternative
5. Check file: `HOW_TO_UPDATE_SWAGGER.md`

---

**Ikuti langkah 1-2-3 di atas, pasti berhasil! 🎉**
