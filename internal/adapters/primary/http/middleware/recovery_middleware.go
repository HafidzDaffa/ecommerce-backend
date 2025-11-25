package middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("🔥 PANIC RECOVERED!")
				log.Printf("   Method: %s", c.Method())
				log.Printf("   Path: %s", c.Path())
				log.Printf("   Headers: %v", c.GetReqHeaders())
				log.Printf("   Panic: %v", r)
				log.Printf("   Stack trace:\n%s", string(debug.Stack()))
				
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   true,
					"message": fmt.Sprintf("%v", r),
				})
			}
		}()
		
		return c.Next()
	}
}
