#!/bin/bash

echo "🧪 Testing Login Endpoint"
echo "========================="
echo ""

# Test 1: Login dengan seller@ecommerce.com (BENAR)
echo "✅ Test 1: Login dengan seller@ecommerce.com (email BENAR)"
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@ecommerce.com",
    "password": "password123"
  }' \
  -w "\nHTTP Status: %{http_code}\n" \
  -s

echo ""
echo "---"
echo ""

# Test 2: Login dengan seller@example.com (SALAH - untuk demo error)
echo "❌ Test 2: Login dengan seller@example.com (email SALAH)"
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "password123"
  }' \
  -w "\nHTTP Status: %{http_code}\n" \
  -s

echo ""
echo "---"
echo ""

# Test 3: Login dengan password salah
echo "❌ Test 3: Login dengan password SALAH"
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@ecommerce.com",
    "password": "wrongpassword"
  }' \
  -w "\nHTTP Status: %{http_code}\n" \
  -s

echo ""
