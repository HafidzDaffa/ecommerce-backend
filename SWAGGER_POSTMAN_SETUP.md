# Swagger & Postman - Quick Setup Guide

## 🎯 Quick Start

### Start Server
```bash
go run cmd/api/main.go
```

### Access Swagger UI
```
http://localhost:8080/swagger/index.html
```

### Import Postman Collection
1. Open Postman
2. Import `E-Commerce_API.postman_collection.json`
3. Import `E-Commerce_API.postman_environment.json`
4. Select "E-Commerce Local Environment"

---

## 📁 Files Created

### Swagger Files
```
docs/
├── docs.go          # Generated Go code
├── swagger.json     # OpenAPI JSON spec
└── swagger.yaml     # OpenAPI YAML spec
```

### Postman Files
```
E-Commerce_API.postman_collection.json     # API collection
E-Commerce_API.postman_environment.json    # Environment variables
```

### Documentation
```
API_DOCUMENTATION_GUIDE.md         # Comprehensive guide
AUTH_API_DOCUMENTATION.md          # Authentication details
SWAGGER_POSTMAN_SETUP.md          # This file
```

---

## 🔧 Swagger Setup Details

### Dependencies Installed
```bash
go get -u github.com/swaggo/fiber-swagger
go get -u github.com/swaggo/files
go install github.com/swaggo/swag/cmd/swag@latest
```

### Code Changes

**1. main.go** - Added Swagger metadata:
```go
// @title E-Commerce Backend API
// @version 1.0
// @description This is an e-commerce backend server with authentication.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

**2. auth_handler.go** - Added endpoint annotations:
```go
// @Summary Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Registration Request"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
```

**3. server/fiber.go** - Added Swagger route:
```go
import _ "ecommerce-backend/docs"
import fiberSwagger "github.com/swaggo/fiber-swagger"

app.Get("/swagger/*", fiberSwagger.WrapHandler)
```

### Generate Swagger Docs
```bash
swag init -g cmd/api/main.go -o docs
```

Run this command whenever you:
- Add new endpoints
- Modify existing endpoints
- Change request/response structures

---

## 📮 Postman Collection Details

### Collection Structure
```
E-Commerce Backend API
├── Authentication
│   ├── Register
│   ├── Login
│   ├── Login as Seller
│   ├── Login as Customer
│   ├── Get Current User (Me)
│   └── Logout
└── Health Check
    └── Health
```

### Auto-saved Variables
When you login via Postman, these are automatically saved:
- `jwt_token` - Used for authentication
- `user_id` - Current user UUID
- `user_email` - User email
- `user_role_id` - User role (1=Customer, 2=Seller, 3=Admin)

### Authorization
Collection-level auth is set to Bearer Token using `{{jwt_token}}`.
Just login once and all protected endpoints work automatically!

---

## 🧪 Testing Workflow

### Using Swagger UI

1. **Open Swagger:**
   ```
   http://localhost:8080/swagger/index.html
   ```

2. **Login to get token:**
   - Expand `POST /auth/login`
   - Click "Try it out"
   - Use: `admin@ecommerce.com` / `password123`
   - Click "Execute"
   - Copy the token from response

3. **Authorize:**
   - Click "Authorize" button (top right)
   - Enter: `Bearer <your-token>`
   - Click "Authorize"

4. **Test protected endpoint:**
   - Go to `GET /auth/me`
   - Click "Try it out"
   - Click "Execute"

### Using Postman

1. **Import files** (see Quick Start above)

2. **Test flow:**
   ```
   1. Send "Login" request → Token saved automatically
   2. Send "Get Current User (Me)" → Uses saved token
   3. No manual token copying needed!
   ```

3. **Switch users:**
   - Use "Login as Admin"
   - Use "Login as Seller"
   - Use "Login as Customer"

---

## 📊 API Endpoints Summary

| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/api/health` | GET | ❌ | Health check |
| `/api/v1/auth/register` | POST | ❌ | Register user |
| `/api/v1/auth/login` | POST | ❌ | Login user |
| `/api/v1/auth/logout` | POST | ❌ | Logout user |
| `/api/v1/auth/me` | GET | ✅ | Get current user |
| `/swagger/*` | GET | ❌ | Swagger UI |

---

## 🔐 Test Accounts

| User | Email | Password | Role |
|------|-------|----------|------|
| Admin | admin@ecommerce.com | password123 | Admin (3) |
| Seller | seller@ecommerce.com | password123 | Seller (2) |
| Customer | customer@ecommerce.com | password123 | Customer (1) |

---

## 🎨 Swagger UI Features

✅ Interactive API documentation
✅ Try out endpoints directly in browser
✅ View request/response schemas
✅ Test authentication flows
✅ No additional tools needed
✅ Auto-generated from code annotations

---

## 🚀 Postman Collection Features

✅ Pre-configured requests
✅ Environment variables
✅ Auto-save JWT tokens
✅ Test scripts
✅ Multiple user roles
✅ One-click testing
✅ Export/share with team

---

## 🔄 Updating Documentation

### After changing API code:

**1. Regenerate Swagger docs:**
```bash
swag init -g cmd/api/main.go -o docs
```

**2. Restart server:**
```bash
go run cmd/api/main.go
```

**3. Refresh Swagger UI:**
```
http://localhost:8080/swagger/index.html
```

**4. Update Postman (if needed):**
- Modify requests in Postman
- Export collection
- Commit updated JSON file

---

## 📝 Swagger Annotation Examples

### Basic Endpoint
```go
// @Summary Short description
// @Description Longer description
// @Tags CategoryName
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} ResponseType
// @Failure 400 {object} ErrorType
// @Router /path [method]
func Handler(c *fiber.Ctx) error {
    // ...
}
```

### Protected Endpoint
```go
// @Security BearerAuth
// @Summary Protected endpoint
// @Description Requires JWT token
// @Success 200 {object} ResponseType
// @Failure 401 {object} ErrorType
// @Router /protected [get]
func ProtectedHandler(c *fiber.Ctx) error {
    // ...
}
```

---

## 🐛 Troubleshooting

### Swagger UI not loading
```bash
# Check if docs exist
ls docs/

# Regenerate
swag init -g cmd/api/main.go -o docs

# Restart server
go run cmd/api/main.go
```

### Postman token not working
1. Make sure environment is selected
2. Check token in environment variables (eye icon)
3. Token expires in 24h - login again
4. Verify Authorization header format: `Bearer <token>`

### Cannot authenticate in Swagger
1. Login first via `POST /auth/login`
2. Copy token from response
3. Click "Authorize" button
4. Enter: `Bearer <token>` (with space)
5. Click "Authorize"

---

## 📖 Learn More

- **Full API Guide:** `API_DOCUMENTATION_GUIDE.md`
- **Auth Details:** `AUTH_API_DOCUMENTATION.md`
- **Swagger Docs:** [swagger.io](https://swagger.io/docs/)
- **Postman Docs:** [learning.postman.com](https://learning.postman.com/)

---

## ✅ Checklist

- [x] Swagger dependencies installed
- [x] Swagger annotations added to handlers
- [x] Swagger docs generated
- [x] Swagger UI accessible at `/swagger/index.html`
- [x] Postman collection created
- [x] Postman environment created
- [x] All endpoints tested
- [x] Documentation created

**Status:** ✅ Ready to use!
