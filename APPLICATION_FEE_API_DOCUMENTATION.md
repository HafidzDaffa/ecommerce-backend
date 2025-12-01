# Application Fee API Documentation

## Overview
API untuk mengelola konfigurasi biaya aplikasi (application fee) yang dapat diatur oleh admin. Application fee dapat berupa persentase atau nilai tetap yang akan ditambahkan pada setiap transaksi.

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
Endpoints yang memerlukan authentication menggunakan Bearer Token JWT di header:
```
Authorization: Bearer <your_jwt_token>
```

---

## Endpoints

### 1. Get All Application Fees
Mendapatkan daftar semua application fees dengan pagination.

**Endpoint:** `GET /application-fees`

**Auth Required:** No

**Query Parameters:**
- `is_active` (boolean, optional) - Filter berdasarkan status aktif (true/false)
- `page` (integer, optional, default: 1) - Nomor halaman
- `per_page` (integer, optional, default: 10) - Jumlah item per halaman

**Success Response:**
```json
{
  "fees": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "fee_name": "Platform Fee (Percentage)",
      "fee_type": "PERCENTAGE",
      "fee_value": 2.5,
      "description": "Platform service fee calculated as percentage of order amount",
      "is_active": true,
      "created_by": "550e8400-e29b-41d4-a716-446655440001",
      "created_at": "2024-11-26T10:00:00Z",
      "updated_at": null
    }
  ],
  "total": 4,
  "page": 1,
  "per_page": 10,
  "total_pages": 1
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/application-fees?is_active=true&page=1&per_page=10"
```

---

### 2. Get Application Fee by ID
Mendapatkan detail application fee berdasarkan ID.

**Endpoint:** `GET /application-fees/:id`

**Auth Required:** No

**URL Parameters:**
- `id` (UUID, required) - Application fee ID

**Success Response:**
```json
{
  "fee": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "fee_name": "Platform Fee (Percentage)",
    "fee_type": "PERCENTAGE",
    "fee_value": 2.5,
    "description": "Platform service fee calculated as percentage of order amount",
    "is_active": true,
    "created_by": "550e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-11-26T10:00:00Z",
    "updated_at": null
  }
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/application-fees/550e8400-e29b-41d4-a716-446655440000"
```

---

### 3. Get Active Application Fee by Type
Mendapatkan application fee aktif terbaru berdasarkan tipe.

**Endpoint:** `GET /application-fees/active`

**Auth Required:** No

**Query Parameters:**
- `fee_type` (string, required) - Tipe fee: "PERCENTAGE" atau "FIXED"

**Success Response:**
```json
{
  "fee": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "fee_name": "Platform Fee (Percentage)",
    "fee_type": "PERCENTAGE",
    "fee_value": 2.5,
    "description": "Platform service fee calculated as percentage of order amount",
    "is_active": true,
    "created_by": "550e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-11-26T10:00:00Z",
    "updated_at": null
  }
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/application-fees/active?fee_type=PERCENTAGE"
```

---

### 4. Calculate Fee
Menghitung jumlah fee berdasarkan application fee ID dan jumlah dasar.

**Endpoint:** `POST /application-fees/calculate`

**Auth Required:** No

**Request Body:**
```json
{
  "fee_id": "550e8400-e29b-41d4-a716-446655440000",
  "base_amount": 100000
}
```

**Success Response:**
```json
{
  "fee_amount": 2500,
  "base_amount": 100000,
  "total_amount": 102500
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/application-fees/calculate" \
  -H "Content-Type: application/json" \
  -d '{
    "fee_id": "550e8400-e29b-41d4-a716-446655440000",
    "base_amount": 100000
  }'
```

---

### 5. Create Application Fee (Admin Only)
Membuat konfigurasi application fee baru.

**Endpoint:** `POST /application-fees`

**Auth Required:** Yes (Admin)

**Request Body:**
```json
{
  "fee_name": "Platform Fee",
  "fee_type": "PERCENTAGE",
  "fee_value": 2.5,
  "description": "Platform service fee",
  "is_active": true
}
```

**Field Validations:**
- `fee_name` (string, required) - Min: 3, Max: 255 karakter
- `fee_type` (string, required) - Must be: "PERCENTAGE" atau "FIXED"
- `fee_value` (float64, required) - Must be greater than 0
  - Untuk PERCENTAGE: max 100
  - Untuk FIXED: jumlah dalam rupiah
- `description` (string, optional)
- `is_active` (boolean, optional, default: true)

**Success Response:**
```json
{
  "message": "Application fee created successfully",
  "fee": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "fee_name": "Platform Fee",
    "fee_type": "PERCENTAGE",
    "fee_value": 2.5,
    "description": "Platform service fee",
    "is_active": true,
    "created_by": "550e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-11-26T10:00:00Z",
    "updated_at": null
  }
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/application-fees" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "fee_name": "Platform Fee",
    "fee_type": "PERCENTAGE",
    "fee_value": 2.5,
    "description": "Platform service fee",
    "is_active": true
  }'
```

---

### 6. Update Application Fee (Admin Only)
Mengupdate konfigurasi application fee yang sudah ada.

**Endpoint:** `PUT /application-fees/:id`

**Auth Required:** Yes (Admin)

**URL Parameters:**
- `id` (UUID, required) - Application fee ID

**Request Body:**
```json
{
  "fee_name": "Updated Platform Fee",
  "fee_type": "PERCENTAGE",
  "fee_value": 3.0,
  "description": "Updated platform service fee",
  "is_active": false
}
```

**Note:** Semua field adalah optional, hanya field yang dikirim yang akan diupdate.

**Success Response:**
```json
{
  "message": "Application fee updated successfully",
  "fee": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "fee_name": "Updated Platform Fee",
    "fee_type": "PERCENTAGE",
    "fee_value": 3.0,
    "description": "Updated platform service fee",
    "is_active": false,
    "created_by": "550e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-11-26T10:00:00Z",
    "updated_at": "2024-11-26T11:00:00Z"
  }
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/api/v1/application-fees/550e8400-e29b-41d4-a716-446655440000" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "fee_value": 3.0,
    "is_active": false
  }'
```

---

### 7. Delete Application Fee (Admin Only)
Menghapus (soft delete) application fee.

**Endpoint:** `DELETE /application-fees/:id`

**Auth Required:** Yes (Admin)

**URL Parameters:**
- `id` (UUID, required) - Application fee ID

**Success Response:**
```json
{
  "message": "Application fee deleted successfully"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/application-fees/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer <your_jwt_token>"
```

---

## Fee Types

### PERCENTAGE
Fee dihitung sebagai persentase dari jumlah dasar (base amount).

**Formula:** `fee_amount = base_amount × (fee_value / 100)`

**Example:**
- Base Amount: Rp 100,000
- Fee Value: 2.5 (%)
- Fee Amount: Rp 100,000 × 0.025 = Rp 2,500
- Total: Rp 102,500

### FIXED
Fee adalah nilai tetap yang ditambahkan ke jumlah dasar.

**Formula:** `fee_amount = fee_value`

**Example:**
- Base Amount: Rp 100,000
- Fee Value: Rp 5,000
- Fee Amount: Rp 5,000
- Total: Rp 105,000

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized"
}
```

### 404 Not Found
```json
{
  "error": "application fee not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error message"
}
```

---

## Integration with Payment Transactions

Application fees dapat diintegrasikan dengan payment transactions:

1. Saat membuat payment transaction, sistem dapat otomatis menghitung dan menambahkan application fee
2. Field `application_fee_id` di payment_transactions table menyimpan referensi ke application fee yang digunakan
3. Field `application_fee_amount` menyimpan jumlah fee yang dihitung

**Example Flow:**
```
1. Get active application fee
   GET /application-fees/active?fee_type=PERCENTAGE

2. Calculate fee
   POST /application-fees/calculate
   {
     "fee_id": "<fee_id>",
     "base_amount": 100000
   }

3. Create payment with application fee
   POST /payments
   {
     "order_id": "<order_id>",
     "payment_method": "credit_card",
     "application_fee_id": "<fee_id>",
     "application_fee_amount": 2500
   }
```

---

## Running Migrations & Seeders

### Run Migrations
```bash
# Navigate to backend directory
cd backend

# Run all migrations up
make migrate-up

# Or use migration tool directly
migrate -path migrations -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" up
```

### Run Seeders
```bash
# Run all seeders in order
psql -U user -d dbname -f seeders/001_roles.sql
psql -U user -d dbname -f seeders/002_order_statuses.sql
psql -U user -d dbname -f seeders/003_categories.sql
psql -U user -d dbname -f seeders/004_users.sql
psql -U user -d dbname -f seeders/005_categories.sql
psql -U user -d dbname -f seeders/006_products.sql
psql -U user -d dbname -f seeders/007_application_fees.sql
```

---

## Testing

### Test dengan cURL

1. **Get all application fees:**
```bash
curl http://localhost:8080/api/v1/application-fees
```

2. **Get active fee:**
```bash
curl "http://localhost:8080/api/v1/application-fees/active?fee_type=PERCENTAGE"
```

3. **Calculate fee:**
```bash
curl -X POST http://localhost:8080/api/v1/application-fees/calculate \
  -H "Content-Type: application/json" \
  -d '{"fee_id":"<your-fee-id>","base_amount":100000}'
```

4. **Create fee (Admin):**
```bash
curl -X POST http://localhost:8080/api/v1/application-fees \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{
    "fee_name": "Test Fee",
    "fee_type": "PERCENTAGE",
    "fee_value": 2.5,
    "is_active": true
  }'
```

---

## Database Schema

```sql
CREATE TABLE application_fees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fee_name VARCHAR(255) NOT NULL,
    fee_type VARCHAR(50) NOT NULL CHECK (fee_type IN ('PERCENTAGE', 'FIXED')),
    fee_value DECIMAL(15,2) NOT NULL CHECK (fee_value >= 0),
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

---

## Notes

1. **Admin Access**: Endpoints untuk create, update, dan delete memerlukan autentikasi admin
2. **Soft Delete**: Delete operation menggunakan soft delete (set deleted_at timestamp)
3. **Active Fee**: Hanya satu application fee yang sebaiknya aktif per tipe pada satu waktu
4. **Validation**: Fee value untuk PERCENTAGE tidak boleh melebihi 100
5. **Integration**: Application fee dapat diintegrasikan dengan payment transactions

---

## Support

Untuk pertanyaan atau bantuan, silakan hubungi tim development.
