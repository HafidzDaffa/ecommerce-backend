# 🔧 Panduan Login dengan Postman

## ✅ Langkah-Langkah yang BENAR

### 1. Setup Request Postman
```
Method: POST
URL: http://localhost:8080/api/v1/auth/login
```

### 2. Set Headers
**Penting!** Tambahkan header ini:
```
Key: Content-Type
Value: application/json
```

### 3. Set Body
Pilih tab **Body** → Pilih **raw** → Pilih **JSON** dari dropdown

Masukkan JSON ini (COPY PASTE persis seperti ini):
```json
{
  "email": "seller@ecommerce.com",
  "password": "password123"
}
```

### 4. Klik Send

## ✅ Response yang BENAR (Status 200)
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "b0000000-0000-0000-0000-000000000002",
    "email": "seller@ecommerce.com",
    "full_name": "Seller User",
    "phone_number": "081234567891",
    "role_id": 2,
    "is_email_verified": true,
    "is_active": true,
    "last_login_at": "2025-11-25T03:57:56.387029Z",
    "created_at": "2025-11-21T12:07:42.234865Z"
  }
}
```

## ❌ Kesalahan Umum

### 1. Email Salah
**SALAH:**
```json
{
  "email": "seller@example.com",  // ❌ Domain salah!
  "password": "password123"
}
```

**BENAR:**
```json
{
  "email": "seller@ecommerce.com",  // ✅ Harus @ecommerce.com
  "password": "password123"
}
```

### 2. Tidak Set Content-Type Header
Jika Anda tidak menambahkan header `Content-Type: application/json`, server tidak bisa parse JSON body.

### 3. Format JSON Salah
- Jangan ada koma di akhir
- Gunakan double quotes ("), bukan single quotes (')
- Pastikan tidak ada whitespace/enter tambahan di email

### 4. Body Type Salah di Postman
Pastikan memilih:
- ✅ **raw** (bukan form-data atau x-www-form-urlencoded)
- ✅ **JSON** dari dropdown (bukan Text)

## 📋 User yang Tersedia

```
1. Admin:
   Email: admin@ecommerce.com
   Password: password123
   Role: 3 (Admin)

2. Seller:
   Email: seller@ecommerce.com
   Password: password123
   Role: 2 (Seller)

3. Customer:
   Email: customer@ecommerce.com
   Password: password123
   Role: 1 (Customer)
```

## 🐛 Troubleshooting

### Jika masih error "invalid email or password":

1. **Cek log backend** (terminal tempat server berjalan)
   - Akan muncul log: `Login attempt - Email: ...`
   - Periksa apakah email yang terlog sama persis dengan yang Anda ketik

2. **Cek whitespace tersembunyi**
   - Copy email dari sini: `seller@ecommerce.com`
   - Jangan ketik manual, paste langsung

3. **Test dengan curl dulu**
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{"email":"seller@ecommerce.com","password":"password123"}'
   ```
   
   Jika curl berhasil tapi Postman gagal, berarti masalah di setup Postman Anda.

4. **Restart Postman**
   Kadang Postman cache request lama.

## 📸 Screenshot Postman yang Benar

```
┌─────────────────────────────────────────┐
│ POST  http://localhost:8080/api/v1/auth/login  [Send]
├─────────────────────────────────────────┤
│ Params | Authorization | Headers | Body | ...
│                                         │
│ Headers (1)                             │
│ Key            | Value                  │
│ Content-Type   | application/json       │
│                                         │
│ Body                                    │
│ ◉ none  ○ form-data  ◉ raw  ○ binary  │
│                          [JSON ▼]       │
│                                         │
│ {                                       │
│   "email": "seller@ecommerce.com",      │
│   "password": "password123"             │
│ }                                       │
└─────────────────────────────────────────┘
```
