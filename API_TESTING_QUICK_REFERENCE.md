# API Testing Quick Reference

Quick reference guide untuk testing API E-Commerce menggunakan cURL, Postman, dan Swagger.

## 🚀 Quick Start Commands

### 1. Start Server
```bash
cd backend
go run cmd/api/main.go
```

Server akan berjalan di: `http://localhost:8080`

### 2. Access Documentation
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **Health Check:** http://localhost:8080/api/health

---

## 🔥 cURL Examples

### Authentication

#### Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "password123",
    "full_name": "Test Seller",
    "phone_number": "081234567890",
    "role_id": 2
  }'
```

#### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "password123"
  }'
```
**Save the token from response!**

#### Get Current User
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

### Categories

#### Get All Categories (Public)
```bash
curl -X GET "http://localhost:8080/api/v1/categories"
```

#### Get Active Categories Only
```bash
curl -X GET "http://localhost:8080/api/v1/categories?is_active=true"
```

#### Get Category by ID
```bash
curl -X GET "http://localhost:8080/api/v1/categories/1"
```

#### Create Category (Auth Required)
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "category_name=Electronics" \
  -F "slug=electronics" \
  -F "icon=📱" \
  -F "is_active=true"
```

#### Create Category with Image
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "category_name=Electronics" \
  -F "slug=electronics" \
  -F "icon=📱" \
  -F "image=@/path/to/image.jpg"
```

#### Update Category
```bash
curl -X PUT http://localhost:8080/api/v1/categories/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "category_name=Consumer Electronics" \
  -F "is_active=true"
```

#### Delete Category
```bash
curl -X DELETE http://localhost:8080/api/v1/categories/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

### Products

#### Get All Products (Public)
```bash
curl -X GET "http://localhost:8080/api/v1/products?page=1&limit=10"
```

#### Get Published Products Only
```bash
curl -X GET "http://localhost:8080/api/v1/products?is_published=true"
```

#### Get Product by ID
```bash
curl -X GET "http://localhost:8080/api/v1/products/PRODUCT_UUID"
```

#### Get Product by Slug
```bash
curl -X GET "http://localhost:8080/api/v1/products/slug/smartphone-xyz-pro"
```

#### Get Products by Category (Public)
```bash
curl -X GET "http://localhost:8080/api/v1/products/category/1?page=1&limit=10"
```

#### Create Product (Auth Required)
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "iPhone 15 Pro",
    "slug": "iphone-15-pro",
    "sku": "ELEC-001",
    "price": 15999000,
    "discount_percent": 10,
    "short_description": "Latest iPhone",
    "description": "Apple iPhone 15 Pro with A17 chip",
    "weight_gram": 200,
    "stock_quantity": 50,
    "is_published": true,
    "category_ids": [1]
  }'
```

#### Update Product
```bash
curl -X PUT http://localhost:8080/api/v1/products/PRODUCT_UUID \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 14999000,
    "stock_quantity": 45,
    "is_published": true
  }'
```

#### Delete Product (Soft Delete)
```bash
curl -X DELETE http://localhost:8080/api/v1/products/PRODUCT_UUID \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

### Product Galleries

#### Get Product Galleries
```bash
curl -X GET "http://localhost:8080/api/v1/products/PRODUCT_UUID/galleries"
```

#### Upload Product Image
```bash
curl -X POST http://localhost:8080/api/v1/products/galleries \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "product_id=PRODUCT_UUID" \
  -F "image=@/path/to/product-image.jpg" \
  -F "display_order=0" \
  -F "is_thumbnail=true"
```

#### Upload Multiple Images
```bash
# Image 1 (thumbnail)
curl -X POST http://localhost:8080/api/v1/products/galleries \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "product_id=PRODUCT_UUID" \
  -F "image=@/path/to/image1.jpg" \
  -F "display_order=0" \
  -F "is_thumbnail=true"

# Image 2
curl -X POST http://localhost:8080/api/v1/products/galleries \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "product_id=PRODUCT_UUID" \
  -F "image=@/path/to/image2.jpg" \
  -F "display_order=1" \
  -F "is_thumbnail=false"
```

#### Update Gallery Metadata
```bash
curl -X PUT http://localhost:8080/api/v1/products/galleries/GALLERY_UUID \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "display_order": 2,
    "is_thumbnail": false
  }'
```

#### Delete Gallery Image
```bash
curl -X DELETE http://localhost:8080/api/v1/products/galleries/GALLERY_UUID \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 📮 Postman Collection

### Import Instructions

1. **Open Postman**
2. Click **Import** button
3. Select files:
   - `E-Commerce_Complete_API.postman_collection.json`
   - `E-Commerce_API.postman_environment.json`
4. Click **Import**
5. Select environment: **E-Commerce API Environment**

### Collection Features

✅ **Organized Folders:**
- Authentication
- Categories  
- Products
- Product Galleries

✅ **Auto-Save Variables:**
- Login → saves `token` and `user_id`
- Create Category → saves `category_id`
- Create Product → saves `product_id`
- Upload Image → saves `gallery_id`

✅ **Pre-configured Headers:**
- Authorization automatically uses `{{token}}`
- Content-Type already set

✅ **Example Data:**
- All requests include sample data
- Easy to modify for your needs

### Quick Workflow

1. **Login** → Token saved automatically
2. **Create Category** → Category ID saved
3. **Create Product** → Product ID saved
4. **Upload Images** → Gallery ID saved
5. **Get Products by Category** → Uses saved IDs

---

## 🎯 Swagger UI

### Access
Open browser: `http://localhost:8080/swagger/index.html`

### Features

✅ **Interactive Documentation:**
- All endpoints with descriptions
- Request/response schemas
- Try it out feature

✅ **Authentication:**
1. Click **Authorize** button
2. Enter: `Bearer YOUR_TOKEN`
3. Click **Authorize**
4. Test protected endpoints

✅ **Try Endpoints:**
1. Select endpoint
2. Click **Try it out**
3. Fill parameters/body
4. Click **Execute**
5. See response

### Sections in Swagger

- **Authentication** - User management
- **Categories** - Category CRUD
- **Products** - Product management
- **Product Galleries** - Image uploads

---

## 🧪 Testing Scenarios

### Scenario 1: Complete Product Setup

```bash
# 1. Login
TOKEN=$(curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"seller@example.com","password":"password123"}' \
  | jq -r '.token')

# 2. Create Category
CATEGORY_ID=$(curl -X POST http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer $TOKEN" \
  -F "category_name=Electronics" \
  -F "slug=electronics" \
  | jq -r '.category.id')

# 3. Create Product
PRODUCT_ID=$(curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"product_name\":\"iPhone 15\",\"slug\":\"iphone-15\",\"price\":15000000,\"weight_gram\":200,\"stock_quantity\":50,\"category_ids\":[$CATEGORY_ID]}" \
  | jq -r '.product.id')

# 4. Upload Image
curl -X POST http://localhost:8080/api/v1/products/galleries \
  -H "Authorization: Bearer $TOKEN" \
  -F "product_id=$PRODUCT_ID" \
  -F "image=@image.jpg" \
  -F "is_thumbnail=true"

# 5. Get Product
curl -X GET "http://localhost:8080/api/v1/products/$PRODUCT_ID"
```

### Scenario 2: Browse Products by Category

```bash
# Get all categories
curl -X GET "http://localhost:8080/api/v1/categories"

# Get products in category 1
curl -X GET "http://localhost:8080/api/v1/products/category/1?page=1&limit=10"

# Get specific product
curl -X GET "http://localhost:8080/api/v1/products/slug/iphone-15"
```

### Scenario 3: Product Management

```bash
# Update stock
curl -X PUT http://localhost:8080/api/v1/products/$PRODUCT_ID \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"stock_quantity": 45}'

# Unpublish product
curl -X PUT http://localhost:8080/api/v1/products/$PRODUCT_ID \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"is_published": false}'

# Delete product
curl -X DELETE http://localhost:8080/api/v1/products/$PRODUCT_ID \
  -H "Authorization: Bearer $TOKEN"
```

---

## 📊 Response Examples

### Success Response (Category)
```json
{
  "message": "Category created successfully",
  "category": {
    "id": 1,
    "category_name": "Electronics",
    "slug": "electronics",
    "icon": "📱",
    "image_path": "https://drive.google.com/uc?id=...",
    "is_active": true,
    "created_at": "2024-11-24T10:00:00Z"
  }
}
```

### Success Response (Product)
```json
{
  "message": "Product created successfully",
  "product": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "product_name": "iPhone 15 Pro",
    "slug": "iphone-15-pro",
    "price": 15999000,
    "stock_quantity": 50,
    "categories": [
      {
        "id": 1,
        "category_name": "Electronics",
        "slug": "electronics"
      }
    ],
    "galleries": []
  }
}
```

### Pagination Response
```json
{
  "products": [...],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

### Error Response
```json
{
  "error": "product not found"
}
```

---

## 🔍 Common Issues & Solutions

### 1. Token Issues
**Problem:** `401 Unauthorized`

**Solutions:**
- Check token format: `Bearer TOKEN` (with space)
- Re-login to get fresh token
- Verify token expiration (24 hours)

### 2. File Upload Issues
**Problem:** Image upload fails

**Solutions:**
- Check `STORAGE_TYPE` in .env
- Verify Google Drive credentials
- Check file size (max 5MB recommended)
- Ensure file path is correct

### 3. Category/Product Not Found
**Problem:** `404 Not Found`

**Solutions:**
- Verify ID is correct (UUID for products, integer for categories)
- Check if resource exists: `GET /products/{id}`
- Ensure not deleted (products use soft delete)

### 4. Validation Errors
**Problem:** `400 Bad Request`

**Solutions:**
- Check required fields are present
- Verify data types (price as number, not string)
- Check field length limits
- Validate email format

---

## 💡 Pro Tips

### 1. **Save Responses**
```bash
# Save token to file
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"seller@example.com","password":"password123"}' \
  | jq -r '.token' > token.txt

# Use saved token
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer $(cat token.txt)"
```

### 2. **Pretty Print JSON**
```bash
curl -X GET "http://localhost:8080/api/v1/products" | jq '.'
```

### 3. **Silent Mode**
```bash
curl -s -X GET "http://localhost:8080/api/v1/categories" | jq '.categories[].category_name'
```

### 4. **Check Response Headers**
```bash
curl -i -X GET "http://localhost:8080/api/v1/products"
```

### 5. **Benchmark API**
```bash
# Install Apache Bench
sudo apt install apache2-utils

# Test endpoint
ab -n 1000 -c 10 http://localhost:8080/api/v1/categories
```

---

## 📝 Environment Variables Reference

```env
# Server
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce_db

# Google Drive (optional)
GOOGLE_DRIVE_CREDENTIALS_PATH=./credentials.json
GOOGLE_DRIVE_FOLDER_ID=your-folder-id
STORAGE_TYPE=google_drive  # or "local"

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h
```

---

## ✅ Testing Checklist

- [ ] Server running on port 8080
- [ ] Database connected
- [ ] Migrations applied
- [ ] Can access Swagger UI
- [ ] Can login and get token
- [ ] Can create category
- [ ] Can create product
- [ ] Can upload image
- [ ] Can get products by category
- [ ] Postman collection imported
- [ ] Environment variables working

---

## 🎓 Additional Resources

- **Full Documentation:** `PRODUCT_API_DOCUMENTATION.md`
- **Swagger Guide:** `SWAGGER_POSTMAN_GUIDE.md`
- **Implementation Summary:** `PRODUCT_IMPLEMENTATION_SUMMARY.md`
- **Swagger UI:** http://localhost:8080/swagger/index.html

---

## 🆘 Need Help?

1. Check server logs: `tail -f backend.log`
2. Verify database: `psql -d ecommerce_db`
3. Test health endpoint: `curl http://localhost:8080/api/health`
4. Review error messages in response
5. Check environment variables: `cat .env`

---

**Happy Testing! 🚀**
