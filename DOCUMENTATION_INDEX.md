# 📚 E-Commerce API Documentation Index

Panduan lengkap untuk semua dokumentasi yang tersedia di project E-Commerce Backend.

---

## 🗂️ Quick Navigation

### 🚀 Getting Started
| Document | Description |
|----------|-------------|
| [README.md](README.md) | Project overview dan setup instructions |
| [QUICK_START.md](QUICK_START.md) | Quick start guide untuk development |
| [MIGRATION_GUIDE.md](MIGRATION_GUIDE.md) | Database migration guide |

---

### 🔐 Authentication
| Document | Description |
|----------|-------------|
| [AUTH_API_DOCUMENTATION.md](AUTH_API_DOCUMENTATION.md) | Complete authentication API docs |
| [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md) | Auth implementation summary |

---

### 📦 Products & Categories  
| Document | Description |
|----------|-------------|
| [PRODUCT_API_DOCUMENTATION.md](PRODUCT_API_DOCUMENTATION.md) | **Complete Product, Category & Gallery API docs** |
| [PRODUCT_IMPLEMENTATION_SUMMARY.md](PRODUCT_IMPLEMENTATION_SUMMARY.md) | Product API implementation summary |

---

### 🧪 API Testing
| Document | Description |
|----------|-------------|
| [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md) | **📖 Complete guide untuk Swagger & Postman** |
| [API_TESTING_QUICK_REFERENCE.md](API_TESTING_QUICK_REFERENCE.md) | **⚡ Quick reference dengan cURL examples** |
| [SWAGGER_POSTMAN_SUMMARY.md](SWAGGER_POSTMAN_SUMMARY.md) | Summary of Swagger & Postman documentation |

---

### 📄 API Documentation Tools
| File | Type | Description |
|------|------|-------------|
| `docs/swagger.json` | Swagger | OpenAPI specification |
| `docs/swagger.yaml` | Swagger | YAML format OpenAPI spec |
| `E-Commerce_Complete_API.postman_collection.json` | Postman | **Complete API collection (28 requests)** |
| `E-Commerce_API.postman_environment.json` | Postman | Environment variables |

---

### 🌱 Database Seeders
| File | Description |
|------|-------------|
| [SEEDER_SUMMARY.md](SEEDER_SUMMARY.md) | Seeder documentation |
| `seeders/004_users.sql` | Sample users |
| `seeders/005_categories.sql` | Sample categories |
| `seeders/006_products.sql` | Sample products |

---

### 📊 Other Documentation
| Document | Description |
|----------|-------------|
| [API_DOCUMENTATION_GUIDE.md](API_DOCUMENTATION_GUIDE.md) | General API documentation guide |
| [VISUAL_GUIDE.md](VISUAL_GUIDE.md) | Visual guide for API |

---

## 🎯 Recommended Reading Order

### For New Developers
1. **README.md** - Understand project structure
2. **QUICK_START.md** - Setup development environment
3. **PRODUCT_API_DOCUMENTATION.md** - Learn API endpoints
4. **SWAGGER_POSTMAN_GUIDE.md** - Learn testing tools
5. **API_TESTING_QUICK_REFERENCE.md** - Start testing

### For Frontend Developers
1. **PRODUCT_API_DOCUMENTATION.md** - API reference
2. **SWAGGER_POSTMAN_GUIDE.md** - Interactive testing
3. **API_TESTING_QUICK_REFERENCE.md** - Quick examples
4. Open Swagger UI at `http://localhost:8080/swagger/index.html`

### For QA/Testers
1. **SWAGGER_POSTMAN_GUIDE.md** - Testing tools setup
2. **API_TESTING_QUICK_REFERENCE.md** - Testing commands
3. Import Postman collection: `E-Commerce_Complete_API.postman_collection.json`
4. Use Swagger UI for exploratory testing

### For DevOps/Deployment
1. **README.md** - Deployment requirements
2. **MIGRATION_GUIDE.md** - Database setup
3. **QUICK_START.md** - Environment configuration
4. **SEEDER_SUMMARY.md** - Sample data setup

---

## 📖 Documentation by Topic

### Authentication & Authorization
- **Complete Guide:** [AUTH_API_DOCUMENTATION.md](AUTH_API_DOCUMENTATION.md)
- **Endpoints:**
  - POST `/auth/register` - Register user
  - POST `/auth/login` - Login
  - GET `/auth/me` - Get current user
  - POST `/auth/logout` - Logout

### Categories
- **Complete Guide:** [PRODUCT_API_DOCUMENTATION.md](PRODUCT_API_DOCUMENTATION.md#category-apis)
- **Endpoints:**
  - GET `/categories` - Get all categories
  - GET `/categories/:id` - Get category by ID
  - POST `/categories` - Create category
  - PUT `/categories/:id` - Update category
  - DELETE `/categories/:id` - Delete category

### Products
- **Complete Guide:** [PRODUCT_API_DOCUMENTATION.md](PRODUCT_API_DOCUMENTATION.md#product-apis)
- **Endpoints:**
  - GET `/products` - Get all products (paginated)
  - GET `/products/:id` - Get product by ID
  - GET `/products/slug/:slug` - Get product by slug
  - GET `/products/category/:id` - **Get products by category**
  - POST `/products` - Create product
  - PUT `/products/:id` - Update product
  - DELETE `/products/:id` - Delete product

### Product Galleries (Image Upload)
- **Complete Guide:** [PRODUCT_API_DOCUMENTATION.md](PRODUCT_API_DOCUMENTATION.md#product-gallery-apis)
- **Endpoints:**
  - GET `/products/:id/galleries` - Get product galleries
  - POST `/products/galleries` - Upload image
  - PUT `/products/galleries/:id` - Update gallery
  - DELETE `/products/galleries/:id` - Delete gallery

---

## 🛠️ Tools & Resources

### Swagger UI
- **URL:** http://localhost:8080/swagger/index.html
- **Guide:** [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md#swagger-documentation)
- **Features:**
  - Interactive API testing
  - Request/response schemas
  - Authentication support
  - Try it out functionality

### Postman Collection
- **Collection:** `E-Commerce_Complete_API.postman_collection.json`
- **Environment:** `E-Commerce_API.postman_environment.json`
- **Guide:** [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md#postman-collection)
- **Features:**
  - 28 pre-configured requests
  - Auto-save variables (token, IDs)
  - Organized folders
  - Test scripts included

### cURL Examples
- **Guide:** [API_TESTING_QUICK_REFERENCE.md](API_TESTING_QUICK_REFERENCE.md)
- **Features:**
  - Copy-paste ready commands
  - Complete workflows
  - Testing scenarios
  - Error handling examples

---

## 🚀 Quick Start Commands

### Start Server
```bash
cd backend
go run cmd/api/main.go
```

### Access Documentation
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **Health Check:** http://localhost:8080/api/health
- **API Base:** http://localhost:8080/api/v1

### Import Postman
1. Open Postman
2. Import `E-Commerce_Complete_API.postman_collection.json`
3. Import `E-Commerce_API.postman_environment.json`
4. Select environment
5. Run "Login" request
6. Start testing!

### Test with cURL
```bash
# Get categories (public)
curl http://localhost:8080/api/v1/categories

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"seller@example.com","password":"password123"}'

# Get products by category (public)
curl http://localhost:8080/api/v1/products/category/1?page=1&limit=10
```

---

## 📊 API Endpoints Summary

### Total Endpoints: **21**

| Category | Endpoints | Public | Protected |
|----------|-----------|--------|-----------|
| Authentication | 4 | 3 | 1 |
| Categories | 5 | 2 | 3 |
| Products | 7 | 4 | 3 |
| Product Galleries | 4 | 1 | 3 |
| Health Check | 1 | 1 | 0 |

### Features
- ✅ JWT Authentication
- ✅ File Upload (Google Drive / Local)
- ✅ Pagination
- ✅ Filtering
- ✅ Many-to-Many Relationships
- ✅ Soft Delete
- ✅ Image Management

---

## 💡 Key Features

### Google Drive Integration
- Upload images to Google Drive
- Public URL generation
- Auto-delete on update/delete
- Alternative: Local storage

### Pagination
- Supported on product list endpoints
- Default: page=1, limit=10
- Max limit: 100

### Auto-Save in Postman
- Login → saves token & user_id
- Create Category → saves category_id
- Create Product → saves product_id
- Upload Image → saves gallery_id

### Interactive Swagger
- Try out any endpoint
- See request/response schemas
- Authentication support
- Model definitions

---

## 📝 Testing Workflows

### Complete Product Setup
```
1. Login (Postman/Swagger/cURL)
   ↓
2. Create Category
   ↓
3. Create Product (with category)
   ↓
4. Upload Product Images
   ↓
5. Get Products by Category
   ↓
6. View Product Details
```

### Browse Products
```
1. Get All Categories
   ↓
2. Select Category
   ↓
3. Get Products by Category
   ↓
4. View Product Details with Images
```

---

## 🔍 Search Documentation

### Find by Topic
- **Authentication:** Search for "auth" or "login"
- **Categories:** Search for "category"
- **Products:** Search for "product"
- **Images:** Search for "gallery" or "upload"
- **Testing:** Search for "postman" or "swagger"

### Find by Tool
- **Swagger:** See [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md)
- **Postman:** See [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md)
- **cURL:** See [API_TESTING_QUICK_REFERENCE.md](API_TESTING_QUICK_REFERENCE.md)

### Find by Role
- **Developer:** Start with README.md
- **Tester:** Start with SWAGGER_POSTMAN_GUIDE.md
- **Frontend:** Start with PRODUCT_API_DOCUMENTATION.md
- **DevOps:** Start with MIGRATION_GUIDE.md

---

## 🆘 Need Help?

### Common Issues
- **401 Unauthorized:** Token expired or invalid → Re-login
- **404 Not Found:** Wrong ID or resource deleted → Verify ID
- **400 Bad Request:** Invalid input → Check request body
- **500 Server Error:** Check server logs

### Debugging
1. Check server logs: `tail -f backend.log`
2. Test health endpoint: `curl http://localhost:8080/api/health`
3. Verify database connection
4. Review error message in response
5. Check environment variables

### Resources
- Read relevant documentation
- Check Swagger UI for schemas
- Use Postman for testing
- Review cURL examples

---

## ✅ Complete File List

### Documentation Files (14)
- README.md
- QUICK_START.md
- MIGRATION_GUIDE.md
- AUTH_API_DOCUMENTATION.md
- IMPLEMENTATION_SUMMARY.md
- PRODUCT_API_DOCUMENTATION.md
- PRODUCT_IMPLEMENTATION_SUMMARY.md
- SWAGGER_POSTMAN_GUIDE.md
- API_TESTING_QUICK_REFERENCE.md
- SWAGGER_POSTMAN_SUMMARY.md
- SEEDER_SUMMARY.md
- API_DOCUMENTATION_GUIDE.md
- VISUAL_GUIDE.md
- **DOCUMENTATION_INDEX.md** (this file)

### API Collections (4)
- docs/swagger.json
- docs/swagger.yaml
- E-Commerce_Complete_API.postman_collection.json
- E-Commerce_API.postman_environment.json

### Database Seeders (3)
- seeders/004_users.sql
- seeders/005_categories.sql
- seeders/006_products.sql

---

## 🎉 Summary

### What's Available

✅ **Complete API Implementation**
- Authentication system
- Category management
- Product management with pagination
- Product gallery with Google Drive integration
- Get products by category endpoint

✅ **Interactive Documentation**
- Swagger UI with try it out
- Postman collection with auto-save
- cURL examples ready to use

✅ **Comprehensive Guides**
- Setup and configuration
- API reference
- Testing workflows
- Troubleshooting tips

✅ **Developer Tools**
- Environment variables
- Test scripts
- Sample data seeders
- Error handling

---

## 📞 Quick Links

- **Swagger UI:** http://localhost:8080/swagger/index.html
- **Health Check:** http://localhost:8080/api/health
- **Main Guide:** [SWAGGER_POSTMAN_GUIDE.md](SWAGGER_POSTMAN_GUIDE.md)
- **Quick Reference:** [API_TESTING_QUICK_REFERENCE.md](API_TESTING_QUICK_REFERENCE.md)
- **API Docs:** [PRODUCT_API_DOCUMENTATION.md](PRODUCT_API_DOCUMENTATION.md)

---

**Happy Coding! 🚀**

All documentation is complete and ready to use for development, testing, and integration!
