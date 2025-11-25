# API Documentation Guide

This guide explains how to use the Swagger UI and Postman collection for testing the E-Commerce Backend API.

## Table of Contents
- [Swagger UI](#swagger-ui)
- [Postman Collection](#postman-collection)
- [Available Endpoints](#available-endpoints)
- [Test Users](#test-users)

---

## Swagger UI

Swagger UI provides interactive API documentation where you can test endpoints directly from your browser.

### Accessing Swagger UI

1. **Start the server:**
   ```bash
   go run cmd/api/main.go
   ```

2. **Open Swagger UI in your browser:**
   ```
   http://localhost:8080/swagger/index.html
   ```

3. **View Swagger JSON:**
   ```
   http://localhost:8080/swagger/doc.json
   ```

### Using Swagger UI

#### Testing Public Endpoints (Register, Login)

1. Expand the endpoint you want to test (e.g., `POST /auth/login`)
2. Click "Try it out"
3. Edit the request body with your data
4. Click "Execute"
5. View the response below

**Example - Login:**
```json
{
  "email": "admin@ecommerce.com",
  "password": "password123"
}
```

#### Testing Protected Endpoints (Requires Authentication)

1. **First, login to get a JWT token:**
   - Go to `POST /auth/login`
   - Click "Try it out"
   - Enter credentials
   - Click "Execute"
   - Copy the `token` from the response

2. **Authorize Swagger:**
   - Click the "Authorize" button at the top right
   - Enter: `Bearer <your-token>` (replace `<your-token>` with the token you copied)
   - Click "Authorize"
   - Click "Close"

3. **Test protected endpoints:**
   - Now you can test `GET /auth/me` and other protected endpoints
   - The token will be automatically included in the Authorization header

### Regenerating Swagger Documentation

If you modify the API endpoints or add new handlers, regenerate the Swagger docs:

```bash
# Install swag CLI if not installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g cmd/api/main.go -o docs

# Restart the server
go run cmd/api/main.go
```

---

## Postman Collection

Postman provides a powerful interface for testing APIs with features like environment variables, test scripts, and collections.

### Files Included

- `E-Commerce_API.postman_collection.json` - Collection with all API endpoints
- `E-Commerce_API.postman_environment.json` - Local environment configuration

### Importing into Postman

1. **Open Postman**

2. **Import Collection:**
   - Click "Import" button (top left)
   - Drag and drop `E-Commerce_API.postman_collection.json`
   - Or click "Upload Files" and select the file
   - Click "Import"

3. **Import Environment:**
   - Click "Import" again
   - Drag and drop `E-Commerce_API.postman_environment.json`
   - Click "Import"

4. **Select Environment:**
   - Click the environment dropdown (top right)
   - Select "E-Commerce Local Environment"

### Using Postman Collection

#### Quick Start

1. **Health Check:**
   - Open "Health Check" folder
   - Click "Health" request
   - Click "Send"
   - You should see: `{"status":"ok","message":"Server is running"}`

2. **Login:**
   - Open "Authentication" folder
   - Click "Login" request
   - Click "Send"
   - The JWT token is automatically saved to environment variable `{{jwt_token}}`

3. **Get Current User:**
   - Click "Get Current User (Me)" request
   - The Authorization header uses `{{jwt_token}}` automatically
   - Click "Send"

#### Features

**Automatic Token Management:**
- When you login, the JWT token is automatically saved to the environment
- All protected endpoints automatically use the token from the environment
- No need to manually copy/paste tokens

**Pre-configured Test Users:**
- Login as Admin: Use "Login" request (default)
- Login as Seller: Use "Login as Seller" request
- Login as Customer: Use "Login as Customer" request

**Test Scripts:**
- Login requests automatically save the token and user info
- Register requests automatically save the new user ID

### Environment Variables

The Postman environment includes these variables:

| Variable | Description | Set By |
|----------|-------------|--------|
| `base_url` | API base URL | Manual (default: localhost:8080) |
| `jwt_token` | JWT authentication token | Automatic (from login) |
| `user_id` | Current user ID | Automatic (from login) |
| `user_email` | Current user email | Automatic (from login) |
| `user_role_id` | Current user role ID | Automatic (from login) |

To view/edit environment variables:
1. Click the eye icon next to environment dropdown
2. Click "Edit" to modify values

---

## Available Endpoints

### Authentication

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/v1/auth/register` | Register new user | No |
| POST | `/api/v1/auth/login` | Login user | No |
| POST | `/api/v1/auth/logout` | Logout user | No |
| GET | `/api/v1/auth/me` | Get current user | Yes |

### Health Check

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/health` | Server health check | No |

---

## Test Users

Three test users are seeded in the database:

### Admin User
```
Email: admin@ecommerce.com
Password: password123
Role ID: 3 (Admin)
```

### Seller User
```
Email: seller@ecommerce.com
Password: password123
Role ID: 2 (Seller)
```

### Customer User
```
Email: customer@ecommerce.com
Password: password123
Role ID: 1 (Customer)
```

---

## Example API Calls

### Register New User

**Request:**
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "newuser@example.com",
  "password": "password123",
  "full_name": "New User",
  "phone_number": "081234567899",
  "role_id": 1
}
```

**Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "uuid",
    "email": "newuser@example.com",
    "full_name": "New User",
    "phone_number": "081234567899",
    "role_id": 1,
    "is_email_verified": false,
    "is_active": true,
    "created_at": "2025-11-21T00:00:00Z"
  }
}
```

### Login

**Request:**
```http
POST /api/v1/auth/login
Content-Type: application/json

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
    "role_id": 3,
    "is_email_verified": true,
    "is_active": true
  }
}
```

### Get Current User (Protected)

**Request:**
```http
GET /api/v1/auth/me
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Response (200 OK):**
```json
{
  "user_id": "a0000000-0000-0000-0000-000000000001",
  "email": "admin@ecommerce.com",
  "role_id": 3
}
```

---

## Troubleshooting

### Swagger UI Issues

**Swagger UI not loading:**
- Make sure the server is running: `go run cmd/api/main.go`
- Check if docs folder exists: `ls docs/`
- Regenerate docs: `swag init -g cmd/api/main.go -o docs`

**401 Unauthorized on protected endpoints:**
- Make sure you clicked "Authorize" and entered the Bearer token
- Token format must be: `Bearer <token>`
- Token expires after 24 hours, login again to get a new token

### Postman Issues

**Requests failing:**
- Check if server is running on port 8080
- Verify environment is selected (top right dropdown)
- Check `base_url` variable is set correctly

**Token not working:**
- Make sure you ran a login request first
- Check if `{{jwt_token}}` has a value in environment variables
- Token expires after 24 hours, login again

**Environment variables not saving:**
- Make sure "E-Commerce Local Environment" is selected
- Check if test scripts are enabled in Postman settings

---

## Additional Resources

- [Swagger Documentation](https://swagger.io/docs/)
- [Postman Learning Center](https://learning.postman.com/)
- [JWT.io](https://jwt.io/) - Decode and verify JWT tokens
- [Fiber Framework](https://docs.gofiber.io/)

---

## Support

For issues or questions, please check:
- `AUTH_API_DOCUMENTATION.md` - Detailed API documentation
- `README.md` - Project setup and information
- Server logs: `backend.log`
