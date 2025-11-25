# Authentication API Documentation

## Overview
This document describes the authentication endpoints implemented in the ecommerce backend using JWT tokens and session cookies.

## Configuration
JWT configuration is stored in `.env`:
```
JWT_SECRET=2QYfqtLY3F7+XoKDN9nMNWu0pk3fX4/VENrtSh2uGUc=
JWT_EXPIRATION=24h
```

## Seeded Users
Three test users are available after running seeders:

| Email | Password | Role | Role ID |
|-------|----------|------|---------|
| admin@ecommerce.com | password123 | Admin | 3 |
| seller@ecommerce.com | password123 | Seller | 2 |
| customer@ecommerce.com | password123 | Customer | 1 |

## API Endpoints

### 1. Register
**POST** `/api/v1/auth/register`

Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "full_name": "User Name",
  "phone_number": "081234567890",
  "role_id": 1
}
```

**Role IDs:**
- `1` - Customer
- `2` - Seller
- `3` - Admin

**Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "full_name": "User Name",
    "phone_number": "081234567890",
    "role_id": 1,
    "is_email_verified": false,
    "is_active": true,
    "created_at": "2025-11-21T00:00:00Z"
  }
}
```

### 2. Login
**POST** `/api/v1/auth/login`

Login with email and password. Returns JWT token and sets session cookie.

**Request Body:**
```json
{
  "email": "admin@ecommerce.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "a0000000-0000-0000-0000-000000000001",
    "email": "admin@ecommerce.com",
    "full_name": "Admin User",
    "phone_number": "081234567890",
    "role_id": 3,
    "is_email_verified": true,
    "is_active": true,
    "last_login_at": "2025-11-21T19:08:47Z",
    "created_at": "2025-11-21T12:07:42Z"
  }
}
```

**Session Cookie:**
- Name: `token`
- Value: JWT token
- HttpOnly: `true`
- MaxAge: `86400` seconds (24 hours)

### 3. Logout
**POST** `/api/v1/auth/logout`

Clear the session cookie.

**Response (200 OK):**
```json
{
  "message": "Logout successful"
}
```

### 4. Get Current User
**GET** `/api/v1/auth/me`

Get current authenticated user information. Requires authentication.

**Headers:**
```
Authorization: Bearer <token>
```

Or use the session cookie automatically.

**Response (200 OK):**
```json
{
  "user_id": "a0000000-0000-0000-0000-000000000001",
  "email": "admin@ecommerce.com",
  "role_id": 3
}
```

## Authentication Methods

The API supports two authentication methods:

### 1. Bearer Token (Header)
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer <token>"
```

### 2. Session Cookie
The cookie is automatically set on login and sent with subsequent requests:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@ecommerce.com","password":"password123"}' \
  -c cookies.txt

curl -X GET http://localhost:8080/api/v1/auth/me \
  -b cookies.txt
```

## Example Usage

### Login as Admin
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@ecommerce.com",
    "password": "password123"
  }'
```

### Register New Customer
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newcustomer@example.com",
    "password": "password123",
    "full_name": "New Customer",
    "role_id": 1
  }'
```

### Access Protected Endpoint
```bash
TOKEN="<your-jwt-token>"
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer $TOKEN"
```

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
  "error": "Invalid email or password"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

## Security Features

1. **Password Hashing**: Passwords are hashed using bcrypt with default cost
2. **JWT Tokens**: Signed with HMAC-SHA256 algorithm
3. **Token Expiration**: Tokens expire after 24 hours
4. **HttpOnly Cookies**: Session cookies are HttpOnly to prevent XSS attacks
5. **Last Login Tracking**: Updates user's last login timestamp on successful authentication

## Implementation Details

### Files Created/Modified

**Domain Models:**
- `internal/core/domain/user.go` - User entity and DTOs
- `internal/core/domain/role.go` - Role entity and constants

**Authentication:**
- `internal/infrastructure/auth/jwt.go` - JWT generation and validation
- `internal/infrastructure/auth/password.go` - Password hashing utilities

**Repositories:**
- `internal/core/ports/user_repository.go` - User repository interface
- `internal/adapters/secondary/repository/user_repository.go` - PostgreSQL implementation

**Services:**
- `internal/core/ports/auth_service.go` - Auth service interface
- `internal/core/services/auth_service.go` - Authentication business logic

**HTTP Handlers:**
- `internal/adapters/primary/http/auth_handler.go` - HTTP handlers for auth endpoints
- `internal/adapters/primary/http/middleware/auth_middleware.go` - JWT authentication middleware

**Configuration:**
- `internal/infrastructure/config/config.go` - Updated with JWT configuration
- `.env` - JWT secret and expiration settings

**Database Seeders:**
- `seeders/004_users.sql` - Seed data for admin, seller, and customer users
