# Postman Collection

This directory contains Postman collections and environments for testing the E-commerce API.

## Files

### Ecommerce-API.postman_collection.json
Complete collection of all API endpoints with pre-configured requests.

**Endpoints Included:**
- Health Check
- User Registration
- User Login
- Get User by ID
- List Users (with pagination)
- Update User
- Delete User

### Ecommerce-API.postman_environment.json
Postman environment file with variables:

**Variables:**
- `base_url`: API base URL (default: http://localhost:8080)
- `user_id`: User ID for testing
- `token`: Authentication token (for future auth implementation)

## Import Instructions

1. Open Postman
2. Click "Import" in the top left
3. Select the collection file: `Ecommerce-API.postman_collection.json`
4. Select the environment file: `Ecommerce-API.postman_environment.json`
5. Click "Import"

## Using the Collection

1. **Set Environment:**
   - Click the environment selector in the top right
   - Choose "E-commerce API Environment"

2. **Run Requests:**
   - Select a request from the collection
   - Click "Send" to execute

3. **Variable Substitution:**
   - All requests use `{{base_url}}` variable
   - Update `base_url` in environment if API runs on different port/host

## Test Scripts

Collection includes test scripts that automatically verify:
- Response status codes
- Response body structure
- Expected values

View test results in the "Test Results" tab after sending requests.

## Example Workflow

```bash
# 1. Health Check
GET {{base_url}}/api/v1/health

# 2. Register User
POST {{base_url}}/api/v1/users/register
{
  "email": "customer@example.com",
  "password": "password123",
  "full_name": "Customer User"
}

# 3. Login
POST {{base_url}}/api/v1/users/login
{
  "email": "customer@example.com",
  "password": "password123"
}

# 4. Get User
GET {{base_url}}/api/v1/users/1

# 5. Update User
PUT {{base_url}}/api/v1/users/1
{
  "full_name": "Updated Name"
}

# 6. List Users
GET {{base_url}}/api/v1/users?page=1&limit=10

# 7. Delete User
DELETE {{base_url}}/api/v1/users/1
```
