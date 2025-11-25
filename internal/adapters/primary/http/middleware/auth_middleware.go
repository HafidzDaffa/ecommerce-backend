package middleware

import (
	"ecommerce-backend/internal/infrastructure/auth"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtService *auth.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("🔐 AuthMiddleware called for: %s %s", c.Method(), c.Path())
		
		authHeader := c.Get("Authorization")
		
		if authHeader == "" {
			token := c.Cookies("token")
			if token == "" {
				log.Println("   ❌ No token found")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Missing authorization token",
				})
			}
			authHeader = "Bearer " + token
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("   ❌ Invalid auth header format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		token := parts[1]

		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			log.Printf("   ❌ Token validation failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		log.Printf("   ✅ Token valid. UserID: %s (type: %T)", claims.UserID, claims.UserID)
		
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("role_id", claims.RoleID)
		
		log.Printf("   ✅ Locals set. Calling next handler...")

		return c.Next()
	}
}
