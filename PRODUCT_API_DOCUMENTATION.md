# Product & Category API Documentation

This document provides comprehensive information about the Product, Category, and Product Gallery APIs.

## Table of Contents
- [Authentication](#authentication)
- [Category APIs](#category-apis)
- [Product APIs](#product-apis)
- [Product Gallery APIs](#product-gallery-apis)
- [Google Drive Integration](#google-drive-integration)
- [Testing with Postman/cURL](#testing-with-postmancurl)

---

## Authentication

Most write operations (POST, PUT, DELETE) require authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

---

## Category APIs

### 1. Get All Categories
**Endpoint:** `GET /api/v1/categories`

**Description:** Retrieve all categories with optional filtering by active status.

**Query Parameters:**
- `is_active` (optional): Filter by active status (true/false)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/categories?is_active=true"
```

**Response:**
```json
{
  "categories": [
    {
      "id": 1,
      "category_name": "Electronics",
      "slug": "electronics",
      "icon": "📱",
      "image_path": "https://drive.google.com/uc?id=...",
      "is_active": true,
      "created_at": "2024-11-24T10:00:00Z"
    }
  ],
  "total": 1
}
```

---

### 2. Get Category by ID
**Endpoint:** `GET /api/v1/categories/:id`

**Description:** Retrieve a single category by its ID.

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/categories/1"
```

**Response:**
```json
{
  "category": {
    "id": 1,
    "category_name": "Electronics",
    "slug": "electronics",
    "icon": "📱",
    "is_active": true,
    "created_at": "2024-11-24T10:00:00Z"
  }
}
```

---

### 3. Create Category
**Endpoint:** `POST /api/v1/categories`

**Description:** Create a new category with optional image upload.

**Authentication:** Required

**Content-Type:** `multipart/form-data`

**Form Fields:**
- `category_name` (required): Category name
- `slug` (optional): URL-friendly slug (auto-generated if not provided)
- `icon` (optional): Icon emoji or text
- `is_active` (optional): Active status (default: true)
- `image` (optional): Category image file

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer <token>" \
  -F "category_name=Electronics" \
  -F "slug=electronics" \
  -F "icon=📱" \
  -F "is_active=true" \
  -F "image=@/path/to/image.jpg"
```

**Response:**
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

---

### 4. Update Category
**Endpoint:** `PUT /api/v1/categories/:id`

**Description:** Update an existing category.

**Authentication:** Required

**Content-Type:** `multipart/form-data`

**Form Fields:** (all optional)
- `category_name`: New category name
- `slug`: New slug
- `icon`: New icon
- `is_active`: New active status
- `image`: New category image (replaces old one)

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/api/v1/categories/1" \
  -H "Authorization: Bearer <token>" \
  -F "category_name=Consumer Electronics" \
  -F "is_active=true"
```

---

### 5. Delete Category
**Endpoint:** `DELETE /api/v1/categories/:id`

**Description:** Delete a category permanently.

**Authentication:** Required

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/categories/1" \
  -H "Authorization: Bearer <token>"
```

**Response:**
```json
{
  "message": "Category deleted successfully"
}
```

---

## Product APIs

### 1. Get All Products
**Endpoint:** `GET /api/v1/products`

**Description:** Retrieve all products with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)
- `is_published` (optional): Filter by published status

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/products?page=1&limit=10&is_published=true"
```

**Response:**
```json
{
  "products": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "user_id": "123e4567-e89b-12d3-a456-426614174001",
      "product_name": "Smartphone XYZ Pro",
      "slug": "smartphone-xyz-pro",
      "sku": "ELEC-001",
      "price": 7999000,
      "discount_percent": 10,
      "short_description": "Latest flagship smartphone",
      "weight_gram": 200,
      "stock_quantity": 50,
      "is_published": true,
      "created_at": "2024-11-24T10:00:00Z",
      "categories": [
        {
          "id": 1,
          "category_name": "Electronics",
          "slug": "electronics"
        }
      ],
      "galleries": [
        {
          "id": "...",
          "image_path": "https://drive.google.com/uc?id=...",
          "display_order": 0,
          "is_thumbnail": true
        }
      ]
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

---

### 2. Get Product by ID
**Endpoint:** `GET /api/v1/products/:id`

**Description:** Retrieve a single product with categories and galleries.

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/products/123e4567-e89b-12d3-a456-426614174000"
```

---

### 3. Get Product by Slug
**Endpoint:** `GET /api/v1/products/slug/:slug`

**Description:** Retrieve a product by its URL-friendly slug.

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/products/slug/smartphone-xyz-pro"
```

---

### 4. Get Products by Category
**Endpoint:** `GET /api/v1/products/category/:category_id`

**Description:** Get all published products in a specific category with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/products/category/1?page=1&limit=10"
```

**Response:**
```json
{
  "products": [...],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

---

### 5. Create Product
**Endpoint:** `POST /api/v1/products`

**Description:** Create a new product.

**Authentication:** Required

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "product_name": "Smartphone XYZ Pro",
  "slug": "smartphone-xyz-pro",
  "sku": "ELEC-001",
  "price": 7999000,
  "discount_percent": 10,
  "short_description": "Latest flagship smartphone",
  "description": "High-end smartphone with advanced features",
  "weight_gram": 200,
  "stock_quantity": 50,
  "is_published": true,
  "category_ids": [1, 2]
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/products" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Smartphone XYZ Pro",
    "slug": "smartphone-xyz-pro",
    "price": 7999000,
    "weight_gram": 200,
    "stock_quantity": 50,
    "category_ids": [1]
  }'
```

---

### 6. Update Product
**Endpoint:** `PUT /api/v1/products/:id`

**Description:** Update an existing product.

**Authentication:** Required

**Content-Type:** `application/json`

**Request Body:** (all fields optional)
```json
{
  "product_name": "Updated Product Name",
  "price": 8999000,
  "stock_quantity": 45,
  "is_published": false,
  "category_ids": [1, 3]
}
```

---

### 7. Delete Product
**Endpoint:** `DELETE /api/v1/products/:id`

**Description:** Soft delete a product (sets deleted_at timestamp).

**Authentication:** Required

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/products/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer <token>"
```

---

## Product Gallery APIs

### 1. Get Product Galleries
**Endpoint:** `GET /api/v1/products/:product_id/galleries`

**Description:** Get all gallery images for a product.

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/v1/products/123e4567-e89b-12d3-a456-426614174000/galleries"
```

**Response:**
```json
{
  "galleries": [
    {
      "id": "...",
      "product_id": "123e4567-e89b-12d3-a456-426614174000",
      "image_path": "https://drive.google.com/uc?id=...",
      "display_order": 0,
      "is_thumbnail": true,
      "created_at": "2024-11-24T10:00:00Z"
    }
  ],
  "total": 1
}
```

---

### 2. Add Product Gallery Image
**Endpoint:** `POST /api/v1/products/galleries`

**Description:** Upload and add an image to product gallery.

**Authentication:** Required

**Content-Type:** `multipart/form-data`

**Form Fields:**
- `product_id` (required): Product UUID
- `image` (required): Image file
- `display_order` (optional): Display order (default: 0)
- `is_thumbnail` (optional): Set as thumbnail (default: false)

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/products/galleries" \
  -H "Authorization: Bearer <token>" \
  -F "product_id=123e4567-e89b-12d3-a456-426614174000" \
  -F "image=@/path/to/image.jpg" \
  -F "display_order=0" \
  -F "is_thumbnail=true"
```

---

### 3. Update Product Gallery
**Endpoint:** `PUT /api/v1/products/galleries/:id`

**Description:** Update gallery image metadata (not the image itself).

**Authentication:** Required

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "display_order": 1,
  "is_thumbnail": false
}
```

---

### 4. Delete Product Gallery
**Endpoint:** `DELETE /api/v1/products/galleries/:id`

**Description:** Delete a gallery image permanently.

**Authentication:** Required

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/products/galleries/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer <token>"
```

---

## Google Drive Integration

### Setup Instructions

1. **Create Google Cloud Project:**
   - Go to [Google Cloud Console](https://console.cloud.google.com/)
   - Create a new project
   - Enable Google Drive API

2. **Create Service Account:**
   - Navigate to "IAM & Admin" > "Service Accounts"
   - Create a new service account
   - Download the JSON credentials file

3. **Create Google Drive Folder:**
   - Create a folder in your Google Drive
   - Share it with the service account email (found in credentials.json)
   - Give "Editor" permissions
   - Copy the folder ID from the URL

4. **Configure Environment:**
   ```env
   GOOGLE_DRIVE_CREDENTIALS_PATH=./credentials.json
   GOOGLE_DRIVE_FOLDER_ID=your-folder-id-here
   STORAGE_TYPE=google_drive
   ```

5. **Place Credentials:**
   - Save the downloaded JSON file as `credentials.json` in the backend root directory

### Local Storage Alternative

If you don't want to use Google Drive, set:
```env
STORAGE_TYPE=local
```

Note: Local storage implementation is basic and mainly for development. Consider implementing file system operations for production use.

---

## Testing with Postman/cURL

### 1. Login to Get Token
```bash
curl -X POST "http://localhost:8080/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "password123"
  }'
```

Save the token from the response.

### 2. Create a Category
```bash
curl -X POST "http://localhost:8080/api/v1/categories" \
  -H "Authorization: Bearer <your-token>" \
  -F "category_name=Electronics" \
  -F "icon=📱"
```

### 3. Create a Product
```bash
curl -X POST "http://localhost:8080/api/v1/products" \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Test Product",
    "slug": "test-product",
    "price": 100000,
    "weight_gram": 500,
    "stock_quantity": 10,
    "category_ids": [1]
  }'
```

### 4. Upload Product Image
```bash
curl -X POST "http://localhost:8080/api/v1/products/galleries" \
  -H "Authorization: Bearer <your-token>" \
  -F "product_id=<product-uuid>" \
  -F "image=@/path/to/image.jpg" \
  -F "is_thumbnail=true"
```

### 5. Get Products by Category
```bash
curl -X GET "http://localhost:8080/api/v1/products/category/1?page=1&limit=10"
```

---

## Error Responses

All endpoints return consistent error responses:

```json
{
  "error": "Error message description"
}
```

Common HTTP status codes:
- `400` - Bad Request (invalid input)
- `401` - Unauthorized (missing or invalid token)
- `404` - Not Found (resource doesn't exist)
- `500` - Internal Server Error

---

## Notes

1. **Image Upload:** Images are uploaded to Google Drive and public URLs are returned
2. **Pagination:** Use `page` and `limit` query parameters for pagination
3. **Soft Delete:** Products are soft-deleted (deleted_at is set) but not removed from database
4. **Categories:** Products can have multiple categories (many-to-many relationship)
5. **Galleries:** Each product can have multiple gallery images with customizable order

---

## Support

For issues or questions, please check:
- API logs: `backend.log`
- Swagger documentation: `http://localhost:8080/swagger/`
- Database migrations: `migrations/` folder
