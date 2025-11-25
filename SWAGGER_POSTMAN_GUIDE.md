# Swagger & Postman Documentation Guide

Dokumentasi lengkap untuk menggunakan Swagger UI dan Postman Collection untuk E-Commerce API.

## 📚 Table of Contents
- [Swagger Documentation](#swagger-documentation)
- [Postman Collection](#postman-collection)
- [Quick Start Guide](#quick-start-guide)
- [API Testing Workflow](#api-testing-workflow)
- [Tips & Best Practices](#tips--best-practices)

---

## 📖 Swagger Documentation

### Accessing Swagger UI

1. **Start the server:**
   ```bash
   go run cmd/api/main.go
   # or
   make run
   ```

2. **Open Swagger UI in your browser:**
   ```
   http://localhost:8080/swagger/index.html
   ```

3. **You will see:**
   - Complete API documentation
   - All endpoints organized by tags
   - Request/response schemas
   - Try it out feature

### Using Swagger UI

#### 1. **Explore Endpoints**
- Navigate through different sections: Authentication, Categories, Products, Product Galleries
- Click on any endpoint to see details
- View request parameters, body schema, and response examples

#### 2. **Authenticate**
To test protected endpoints:

1. First, login via `/auth/login` endpoint:
   - Click on "POST /auth/login"
   - Click "Try it out"
   - Enter credentials:
     ```json
     {
       "email": "seller@example.com",
       "password": "password123"
     }
     ```
   - Click "Execute"
   - Copy the `token` from the response

2. Authorize Swagger:
   - Click the "Authorize" button at the top right
   - Enter: `Bearer <your-token>`
   - Click "Authorize"
   - Click "Close"

3. Now you can test all protected endpoints!

#### 3. **Try Endpoints**
- Click on any endpoint
- Click "Try it out"
- Fill in parameters/body
- Click "Execute"
- See the response below

### Regenerating Swagger Docs

After modifying API handlers or adding new endpoints:

```bash
# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
~/go/bin/swag init -g cmd/api/main.go -o docs

# Or use make command if available
make swagger
```

---

## 📮 Postman Collection

### Importing the Collection

1. **Open Postman**

2. **Import Collection:**
   - Click "Import" button
   - Select file: `E-Commerce_Complete_API.postman_collection.json`
   - Click "Import"

3. **Import Environment:**
   - Click gear icon (⚙️) in top right
   - Click "Import"
   - Select file: `E-Commerce_API.postman_environment.json`
   - Click "Import"

4. **Select Environment:**
   - In top right, select "E-Commerce API Environment" from dropdown

### Collection Structure

The collection is organized into folders:

```
E-Commerce Complete API
├── Authentication
│   ├── Register User
│   ├── Login (auto-saves token)
│   ├── Get Current User
│   └── Logout
├── Categories
│   ├── Get All Categories
│   ├── Get Category by ID
│   ├── Create Category (auto-saves category_id)
│   ├── Update Category
│   └── Delete Category
├── Products
│   ├── Get All Products
│   ├── Get Product by ID
│   ├── Get Product by Slug
│   ├── Get Products by Category
│   ├── Create Product (auto-saves product_id)
│   ├── Update Product
│   └── Delete Product
├── Product Galleries
│   ├── Get Product Galleries
│   ├── Upload Product Image (auto-saves gallery_id)
│   ├── Update Product Gallery
│   └── Delete Product Gallery
└── Health Check
```

### Environment Variables

The environment includes:

| Variable | Description | Auto-populated |
|----------|-------------|----------------|
| `base_url` | API base URL | No (default: localhost:8080/api/v1) |
| `token` | JWT token | Yes (from Login) |
| `user_id` | Current user ID | Yes (from Login) |
| `category_id` | Last created category | Yes (from Create Category) |
| `product_id` | Last created product | Yes (from Create Product) |
| `gallery_id` | Last created gallery | Yes (from Upload Image) |

---

## 🚀 Quick Start Guide

### Step 1: Setup

1. Make sure server is running:
   ```bash
   go run cmd/api/main.go
   ```

2. Import Postman collection and environment

3. Select the environment

### Step 2: Authentication

1. **Register a User** (if needed):
   - Open "Authentication" → "Register User"
   - Modify the body with your details
   - Click "Send"
   - Note: role_id values: 1=Admin, 2=Seller, 3=Customer

2. **Login:**
   - Open "Authentication" → "Login"
   - Enter credentials
   - Click "Send"
   - **Token is automatically saved!** Check environment variables

3. **Verify Authentication:**
   - Open "Authentication" → "Get Current User"
   - Click "Send"
   - You should see your user info

### Step 3: Create a Category

1. Open "Categories" → "Create Category"
2. In the Body tab (form-data), fill:
   - `category_name`: Electronics
   - `slug`: electronics
   - `icon`: 📱
   - `is_active`: true
   - `image`: (optional) Select a file
3. Click "Send"
4. **Category ID is automatically saved!**

### Step 4: Create a Product

1. Open "Products" → "Create Product"
2. Modify the JSON body:
   ```json
   {
     "product_name": "iPhone 15 Pro",
     "slug": "iphone-15-pro",
     "sku": "ELEC-001",
     "price": 15999000,
     "discount_percent": 0,
     "short_description": "Latest iPhone",
     "description": "Apple iPhone 15 Pro with A17 chip",
     "weight_gram": 200,
     "stock_quantity": 50,
     "is_published": true,
     "category_ids": [{{category_id}}]
   }
   ```
3. Click "Send"
4. **Product ID is automatically saved!**

### Step 5: Upload Product Images

1. Open "Product Galleries" → "Upload Product Image"
2. In the Body tab (form-data):
   - `product_id`: {{product_id}} (already filled)
   - `image`: Select an image file
   - `display_order`: 0
   - `is_thumbnail`: true
3. Click "Send"
4. Image is uploaded to Google Drive (or local storage)
5. **Gallery ID is automatically saved!**

### Step 6: Get Products by Category

1. Open "Products" → "Get Products by Category"
2. The URL uses {{category_id}} automatically
3. Click "Send"
4. See all products in that category

---

## 🧪 API Testing Workflow

### Complete Product Creation Flow

```
1. Login
   ↓ (saves token)
2. Create Category
   ↓ (saves category_id)
3. Create Product (with category_ids)
   ↓ (saves product_id)
4. Upload Product Images (multiple)
   ↓ (saves gallery_id)
5. Get Product by ID
   ↓ (shows product with categories and galleries)
6. Get Products by Category
   ↓ (shows all products in category)
```

### Testing Different Scenarios

#### 1. **Public Access (No Auth)**
Test these without authentication:
- Get All Categories
- Get Category by ID
- Get All Products
- Get Product by ID
- Get Product by Slug
- Get Products by Category
- Get Product Galleries

#### 2. **Protected Operations (With Auth)**
These require authentication:
- Create/Update/Delete Category
- Create/Update/Delete Product
- Upload/Update/Delete Product Gallery

#### 3. **Pagination Testing**
Test products endpoints with different page/limit:
```
GET /products?page=1&limit=5
GET /products?page=2&limit=10
GET /products/category/1?page=1&limit=20
```

#### 4. **Filter Testing**
Test with filters:
```
GET /categories?is_active=true
GET /products?is_published=true
GET /products?is_published=false
```

---

## 💡 Tips & Best Practices

### Postman Tips

1. **Use Environment Variables:**
   - Don't hardcode IDs, use `{{variable_name}}`
   - Environment variables are shared across all requests

2. **Save Example Responses:**
   - After successful requests, save responses as examples
   - Helps document expected behavior

3. **Use Pre-request Scripts:**
   - Generate dynamic data
   - Example: Generate random SKU
   ```javascript
   pm.variables.set("random_sku", "PROD-" + Math.floor(Math.random() * 10000));
   ```

4. **Use Test Scripts:**
   - Already included for auto-saving IDs
   - Add custom validations:
   ```javascript
   pm.test("Status code is 200", function () {
       pm.response.to.have.status(200);
   });
   
   pm.test("Response has products array", function () {
       var jsonData = pm.response.json();
       pm.expect(jsonData.products).to.be.an('array');
   });
   ```

5. **Organize with Folders:**
   - Keep related requests together
   - Use descriptive names

### Swagger Tips

1. **Use Models:**
   - Swagger shows request/response schemas
   - Copy example values for testing

2. **Download OpenAPI Spec:**
   - Available at: `http://localhost:8080/swagger/doc.json`
   - Import to other tools (Postman, Insomnia, etc.)

3. **Customize Swagger UI:**
   - Modify `docs/swagger.json` for custom descriptions
   - Regenerate with `swag init`

### File Upload Tips

1. **Image Requirements:**
   - Supported formats: JPG, PNG, GIF, WebP
   - Recommended max size: 5MB
   - Use optimized images for better performance

2. **Google Drive Setup:**
   - Follow setup in `PRODUCT_API_DOCUMENTATION.md`
   - Or use `STORAGE_TYPE=local` for development

3. **Testing File Uploads:**
   - In Postman: Select file in form-data
   - In Swagger: Click "Choose File" button
   - Verify upload by checking returned `image_path`

### Debugging Tips

1. **Check Server Logs:**
   ```bash
   tail -f backend.log
   ```

2. **Common Errors:**
   - `401 Unauthorized`: Token expired or invalid
   - `404 Not Found`: Wrong ID or resource doesn't exist
   - `400 Bad Request`: Invalid input data

3. **Verify Database:**
   ```sql
   SELECT * FROM categories;
   SELECT * FROM products;
   SELECT * FROM product_galleries;
   ```

---

## 📊 API Endpoints Summary

### Authentication
| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| POST | /auth/register | No | Register new user |
| POST | /auth/login | No | Login user |
| GET | /auth/me | Yes | Get current user |
| POST | /auth/logout | No | Logout user |

### Categories
| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | /categories | No | Get all categories |
| GET | /categories/:id | No | Get category by ID |
| POST | /categories | Yes | Create category |
| PUT | /categories/:id | Yes | Update category |
| DELETE | /categories/:id | Yes | Delete category |

### Products
| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | /products | No | Get all products (paginated) |
| GET | /products/:id | No | Get product by ID |
| GET | /products/slug/:slug | No | Get product by slug |
| GET | /products/category/:id | No | Get products by category |
| POST | /products | Yes | Create product |
| PUT | /products/:id | Yes | Update product |
| DELETE | /products/:id | Yes | Delete product |

### Product Galleries
| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | /products/:id/galleries | No | Get product galleries |
| POST | /products/galleries | Yes | Upload product image |
| PUT | /products/galleries/:id | Yes | Update gallery |
| DELETE | /products/galleries/:id | Yes | Delete gallery |

---

## 🎓 Learning Resources

### Swagger/OpenAPI
- [Swagger Documentation](https://swagger.io/docs/)
- [OpenAPI Specification](https://swagger.io/specification/)
- [Swaggo GitHub](https://github.com/swaggo/swag)

### Postman
- [Postman Learning Center](https://learning.postman.com/)
- [Postman Collections](https://learning.postman.com/docs/sending-requests/intro-to-collections/)
- [Postman Environments](https://learning.postman.com/docs/sending-requests/managing-environments/)

---

## 🆘 Troubleshooting

### Swagger Not Loading

1. Check if server is running
2. Verify URL: `http://localhost:8080/swagger/index.html`
3. Regenerate docs: `~/go/bin/swag init -g cmd/api/main.go -o docs`
4. Check for import errors in main.go: `_ "ecommerce-backend/docs"`

### Postman Token Issues

1. Check if token is in environment variables
2. Verify token format: `Bearer <token>` (note the space)
3. Re-login to get new token
4. Check token expiration (default: 24 hours)

### File Upload Errors

1. Verify `STORAGE_TYPE` in .env
2. For Google Drive:
   - Check credentials.json exists
   - Verify folder permissions
   - Check GOOGLE_DRIVE_FOLDER_ID
3. For local storage:
   - Ensure uploads directory exists
   - Check file permissions

---

## 📝 Notes

1. **Token Expiration:**
   - JWT tokens expire after 24 hours (configurable in .env)
   - Re-login when token expires

2. **Database State:**
   - Use seeders to populate test data
   - Reset database with migrations if needed

3. **Google Drive Storage:**
   - Files are public by default
   - Consider implementing access control for production

4. **Rate Limiting:**
   - Not implemented yet
   - Consider adding for production

---

## ✅ Checklist

Before testing:
- [ ] Server is running
- [ ] Database migrations applied
- [ ] Seeders executed (optional)
- [ ] Google Drive configured (if using)
- [ ] Postman collection imported
- [ ] Environment selected in Postman
- [ ] Swagger UI accessible

---

## 🎉 Summary

You now have:
- ✅ Complete Swagger documentation at `/swagger/index.html`
- ✅ Comprehensive Postman collection with auto-save features
- ✅ Environment variables for easy testing
- ✅ Organized API structure
- ✅ Testing workflows and best practices

Happy API testing! 🚀
