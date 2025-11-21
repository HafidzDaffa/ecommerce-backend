package server

import (
	"ecommerce-backend/internal/infrastructure/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

func NewFiberServer(cfg *config.Config, db *sqlx.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorHandler: customErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     getOrigins(cfg.CORS.AllowedOrigins),
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	setupRoutes(app, db)

	return app
}

func setupRoutes(app *fiber.App, db *sqlx.DB) {
	api := app.Group("/api")
	
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v1",
		})
	})
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}

func getOrigins(origins []string) string {
	result := ""
	for i, origin := range origins {
		if i > 0 {
			result += ","
		}
		result += origin
	}
	return result
}
