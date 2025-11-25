#!/bin/bash

echo "🔐 Step 1: Login to get token..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@ecommerce.com",
    "password": "password123"
  }')

TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')

if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
  echo "❌ Login failed!"
  echo "$LOGIN_RESPONSE" | jq .
  exit 1
fi

echo "✅ Login successful! Token: ${TOKEN:0:50}..."
echo ""
echo "📦 Step 2: Creating product..."

PRODUCT_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "product_name": "Smartphone XYZ Pro",
    "slug": "smartphone-xyz-pro",
    "sku": "ELEC-001",
    "price": 7999000,
    "discount_percent": 10,
    "short_description": "Latest flagship smartphone",
    "description": "High-end smartphone with advanced features, 6.7 inch display, 128GB storage",
    "weight_gram": 200,
    "stock_quantity": 50,
    "is_published": true,
    "category_ids": [16]
  }')

echo "$PRODUCT_RESPONSE" | jq .

# Check if successful
if echo "$PRODUCT_RESPONSE" | jq -e '.product' > /dev/null 2>&1; then
  echo ""
  echo "✅ Product created successfully!"
  PRODUCT_ID=$(echo "$PRODUCT_RESPONSE" | jq -r '.product.id')
  echo "   Product ID: $PRODUCT_ID"
  echo "   Product Name: $(echo "$PRODUCT_RESPONSE" | jq -r '.product.product_name')"
else
  echo ""
  echo "❌ Product creation failed!"
fi
