package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/yourusername/ecommerce-go-vue/backend/application/usecases"
	_ "github.com/yourusername/ecommerce-go-vue/backend/docs"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/config"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
	http "github.com/yourusername/ecommerce-go-vue/backend/interfaces/http"
)

// @title E-commerce API
// @version 1.0
// @description E-commerce backend API with Go Fiber
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()

	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	userRepo := database.NewUserRepository()
	database.NewRoleRepository()
	userUseCase := usecases.NewUserUseCase(userRepo, cfg)

	userHandler := http.NewUserHandler(userUseCase)

	router := http.NewRouter(app, userHandler)
	router.SetupRoutes()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
