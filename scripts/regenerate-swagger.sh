#!/bin/bash

# Script to regenerate Swagger documentation
# This should be run after modifying API handlers

set -e

echo "🔄 Regenerating Swagger documentation..."
echo ""

# Check if swag is installed
if ! command -v ~/go/bin/swag &> /dev/null; then
    echo "❌ swag tool not found!"
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
    echo "✅ swag installed successfully"
    echo ""
fi

# Generate Swagger docs
echo "📝 Generating Swagger docs..."
~/go/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -d ./,./internal/adapters/primary/http

echo ""
echo "✅ Swagger documentation generated successfully!"
echo ""
echo "📊 Endpoints found:"
jq '.paths | keys | length' docs/swagger.json | xargs echo "   Total endpoints:"
echo ""
echo "📖 To view documentation:"
echo "   1. Start server: go run cmd/api/main.go"
echo "   2. Open browser: http://localhost:8080/swagger/index.html"
echo ""
echo "💡 Tip: If Swagger UI doesn't show new endpoints:"
echo "   1. Hard refresh browser: Ctrl+Shift+R (or Cmd+Shift+R on Mac)"
echo "   2. Clear browser cache"
echo "   3. Restart the server"
echo ""
