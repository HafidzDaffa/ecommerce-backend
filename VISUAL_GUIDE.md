# 📸 Visual Guide - Swagger & Postman

## 🎯 Quick Access

| Tool | URL | Description |
|------|-----|-------------|
| **Swagger UI** | http://localhost:8080/swagger/index.html | Interactive API documentation |
| **Swagger JSON** | http://localhost:8080/swagger/doc.json | OpenAPI specification |
| **Health Check** | http://localhost:8080/api/health | Server status |
| **API Base** | http://localhost:8080/api/v1 | API version 1 |

---

## 📱 Swagger UI Preview

### Main Page
```
┌─────────────────────────────────────────────────────────┐
│  E-Commerce Backend API                      [Authorize] │
│  Version 1.0                                             │
├─────────────────────────────────────────────────────────┤
│                                                          │
│  Schemes: [http] [https]                                │
│  Base URL: localhost:8080/api/v1                        │
│                                                          │
│  ▼ Authentication                                        │
│    ┌──────────────────────────────────────────┐        │
│    │ POST /auth/register                       │        │
│    │ Register a new user                       │        │
│    └──────────────────────────────────────────┘        │
│    ┌──────────────────────────────────────────┐        │
│    │ POST /auth/login                          │        │
│    │ Login user                                │        │
│    └──────────────────────────────────────────┘        │
│    ┌──────────────────────────────────────────┐        │
│    │ POST /auth/logout                         │        │
│    │ Logout user                               │        │
│    └──────────────────────────────────────────┘        │
│    ┌──────────────────────────────────────────┐        │
│    │ GET /auth/me                         🔒   │        │
│    │ Get current user                          │        │
│    └──────────────────────────────────────────┘        │
│                                                          │
└─────────────────────────────────────────────────────────┘

🔒 = Requires authentication
```

### Login Endpoint Expanded
```
┌─────────────────────────────────────────────────────────┐
│  POST /auth/login                        [Try it out]   │
├─────────────────────────────────────────────────────────┤
│  Authenticate user with email and password, returns     │
│  JWT token and sets session cookie (24 hours)           │
│                                                          │
│  Parameters                                              │
│  ┌─────────────────────────────────────────────────┐   │
│  │ request * (body)                                 │   │
│  │ ┌─────────────────────────────────────────────┐ │   │
│  │ │ {                                            │ │   │
│  │ │   "email": "admin@ecommerce.com",           │ │   │
│  │ │   "password": "password123"                 │ │   │
│  │ │ }                                            │ │   │
│  │ └─────────────────────────────────────────────┘ │   │
│  └─────────────────────────────────────────────────┘   │
│                                                          │
│                        [Execute]                         │
│                                                          │
│  Responses                                               │
│  ┌─────────────────────────────────────────────────┐   │
│  │ Code: 200                                        │   │
│  │ Response body:                                   │   │
│  │ {                                                │   │
│  │   "message": "Login successful",                │   │
│  │   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI...",  │   │
│  │   "user": {                                      │   │
│  │     "id": "a0000000-0000-0000-0000-000...",     │   │
│  │     "email": "admin@ecommerce.com",             │   │
│  │     "role_id": 3                                │   │
│  │   }                                              │   │
│  │ }                                                │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

### Authorize Dialog
```
┌──────────────────────────────────────┐
│  Available authorizations            │
├──────────────────────────────────────┤
│  BearerAuth (apiKey)                 │
│  ┌────────────────────────────────┐ │
│  │ Value: Bearer <token-here>     │ │
│  └────────────────────────────────┘ │
│                                      │
│  Type "Bearer" followed by a space  │
│  and JWT token.                     │
│                                      │
│         [Authorize]  [Close]         │
└──────────────────────────────────────┘
```

---

## 📮 Postman Preview

### Collection Structure
```
📁 E-Commerce Backend API
  │
  ├─ 📁 Authentication
  │   ├─ 📄 Register
  │   ├─ 📄 Login
  │   ├─ 📄 Login as Seller
  │   ├─ 📄 Login as Customer
  │   ├─ 📄 Get Current User (Me)
  │   └─ 📄 Logout
  │
  └─ 📁 Health Check
      └─ 📄 Health

🌍 E-Commerce Local Environment  [Active]
```

### Login Request View
```
┌─────────────────────────────────────────────────────────┐
│  POST  http://localhost:8080/api/v1/auth/login    Send │
├─────────────────────────────────────────────────────────┤
│  Params  Authorization  Headers  Body  Pre-req  Tests   │
├─────────────────────────────────────────────────────────┤
│  ● raw    JSON ▼                                         │
│  ┌─────────────────────────────────────────────────┐   │
│  │ {                                                │   │
│  │   "email": "admin@ecommerce.com",               │   │
│  │   "password": "password123"                     │   │
│  │ }                                                │   │
│  │                                                  │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘

Response:
┌─────────────────────────────────────────────────────────┐
│  Body  Cookies  Headers  Test Results                   │
├─────────────────────────────────────────────────────────┤
│  Status: 200 OK    Time: 45ms    Size: 548 B            │
│  ┌─────────────────────────────────────────────────┐   │
│  │ {                                                │   │
│  │   "message": "Login successful",                │   │
│  │   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6...", │   │
│  │   "user": {                                      │   │
│  │     "id": "a0000000-0000-0000-0000-000...",     │   │
│  │     "email": "admin@ecommerce.com",             │   │
│  │     "full_name": "Admin User",                  │   │
│  │     "role_id": 3,                               │   │
│  │     "is_active": true                           │   │
│  │   }                                              │   │
│  │ }                                                │   │
│  └─────────────────────────────────────────────────┘   │
│                                                          │
│  ✓ jwt_token saved to environment                       │
│  ✓ user_id saved to environment                         │
│  ✓ user_email saved to environment                      │
│  ✓ user_role_id saved to environment                    │
└─────────────────────────────────────────────────────────┘
```

### Environment Variables View
```
┌──────────────────────────────────────────────────────┐
│  E-Commerce Local Environment              [Edit]    │
├──────────────────────────────────────────────────────┤
│  Variable        Initial Value    Current Value      │
├──────────────────────────────────────────────────────┤
│  base_url        localhost:8080   localhost:8080     │
│  jwt_token                         eyJhbGciOiJI...  │
│  user_id                           a0000000-000...   │
│  user_email                        admin@ecomm...   │
│  user_role_id                      3                 │
└──────────────────────────────────────────────────────┘
```

### Protected Request (Get Me)
```
┌─────────────────────────────────────────────────────────┐
│  GET  http://localhost:8080/api/v1/auth/me       Send  │
├─────────────────────────────────────────────────────────┤
│  Params  Authorization  Headers  Body  Pre-req  Tests   │
├─────────────────────────────────────────────────────────┤
│  Type: Bearer Token                                      │
│  Token: {{jwt_token}}  ✓ Auto-filled from environment   │
└─────────────────────────────────────────────────────────┘

Authorization Header (auto-generated):
┌─────────────────────────────────────────────────────────┐
│  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI...  │
└─────────────────────────────────────────────────────────┘
```

---

## 🔄 Workflow Diagrams

### Swagger Authentication Flow
```
┌──────────┐     ┌────────────┐     ┌──────────┐
│ Open UI  │────▶│ POST Login │────▶│ Copy     │
│          │     │ Try it out │     │ Token    │
└──────────┘     └────────────┘     └──────────┘
                                           │
                                           ▼
┌──────────┐     ┌────────────┐     ┌──────────┐
│ Test     │◀────│ Test       │◀────│ Click    │
│ /auth/me │     │ Protected  │     │ Authorize│
│          │     │ Endpoints  │     │          │
└──────────┘     └────────────┘     └──────────┘
```

### Postman Testing Flow
```
┌──────────┐     ┌────────────┐     ┌──────────┐
│ Import   │────▶│ Select     │────▶│ Click    │
│ Files    │     │ Environment│     │ Login    │
└──────────┘     └────────────┘     └──────────┘
                                           │
                                           ▼
                      ┌────────────────────────────┐
                      │ Token Auto-Saved           │
                      │ to {{jwt_token}}           │
                      └────────────────────────────┘
                                           │
                                           ▼
┌──────────┐     ┌────────────┐     ┌──────────┐
│ All      │◀────│ Token      │◀────│ Click    │
│ Protected│     │ Auto-Used  │     │ Get Me   │
│ Endpoints│     │            │     │          │
└──────────┘     └────────────┘     └──────────┘
```

---

## 🎨 Endpoint Color Coding

### In Swagger UI
```
🟦 GET     - Blue   - Read operations
🟩 POST    - Green  - Create operations
🟨 PUT     - Yellow - Update operations
🟥 DELETE  - Red    - Delete operations
```

### Status Codes
```
✅ 200 - OK                  Success
✅ 201 - Created             Resource created
⚠️  400 - Bad Request         Invalid input
⚠️  401 - Unauthorized        Auth required/failed
⚠️  404 - Not Found           Resource not found
❌ 500 - Internal Server Error Server error
```

---

## 📋 Quick Reference Cards

### Swagger Quick Commands
| Action | Steps |
|--------|-------|
| **Test Login** | 1. Expand POST /auth/login<br>2. Click "Try it out"<br>3. Click "Execute" |
| **Authorize** | 1. Click "Authorize" (🔓)<br>2. Enter `Bearer <token>`<br>3. Click "Authorize" |
| **Test Protected** | 1. Authorize first<br>2. Expand GET /auth/me<br>3. Click "Try it out"<br>4. Click "Execute" |

### Postman Quick Commands
| Action | Steps |
|--------|-------|
| **Import** | 1. Click "Import"<br>2. Drag JSON files<br>3. Click "Import" |
| **Set Environment** | 1. Click dropdown (top right)<br>2. Select "E-Commerce Local" |
| **Test Flow** | 1. Send "Login"<br>2. Send "Get Current User" |
| **View Token** | 1. Click eye icon (👁️)<br>2. Find `jwt_token` |

---

## 📊 Testing Checklist

### ✅ Swagger UI
- [ ] Open http://localhost:8080/swagger/index.html
- [ ] Verify all endpoints visible
- [ ] Test POST /auth/login
- [ ] Copy token from response
- [ ] Click "Authorize" and enter token
- [ ] Test GET /auth/me
- [ ] Verify response shows user info

### ✅ Postman
- [ ] Import collection JSON
- [ ] Import environment JSON
- [ ] Select environment
- [ ] Send "Login" request
- [ ] Verify token saved to environment
- [ ] Send "Get Current User" request
- [ ] Verify auto-authorization works
- [ ] Try "Login as Seller"
- [ ] Try "Login as Customer"

---

## 🎯 Success Indicators

### Swagger Working ✓
```
✓ Swagger UI loads
✓ API endpoints listed
✓ "Try it out" works
✓ Authorization saves token
✓ Protected endpoints accessible
```

### Postman Working ✓
```
✓ Collection imported
✓ Environment active
✓ Login saves token automatically
✓ Protected requests work without manual auth
✓ Environment variables populated
```

---

## 📱 Mobile Testing (Optional)

You can also test the API from mobile devices:

```
1. Find your computer's local IP: ifconfig or ipconfig
2. Update Postman environment: base_url = http://192.168.x.x:8080
3. Or access Swagger: http://192.168.x.x:8080/swagger/index.html
4. Make sure firewall allows connections on port 8080
```

---

## 🎓 Learning Resources

### Swagger/OpenAPI
- Swagger Editor: https://editor.swagger.io/
- OpenAPI Spec: https://swagger.io/specification/
- Swaggo Docs: https://github.com/swaggo/swag

### Postman
- Learning Center: https://learning.postman.com/
- Workspace: https://www.postman.com/
- API Testing: https://www.postman.com/api-platform/api-testing/

### JWT
- JWT.io: https://jwt.io/
- Debugger: https://jwt.io/#debugger
- Introduction: https://jwt.io/introduction

---

## 🎉 You're All Set!

Your API documentation is ready:
- ✅ Swagger UI at /swagger/index.html
- ✅ Postman collection ready to import
- ✅ Test accounts available
- ✅ Comprehensive documentation

**Happy Testing! 🚀**
