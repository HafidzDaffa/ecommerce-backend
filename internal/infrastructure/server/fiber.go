package server

import (
	"ecommerce-backend/internal/adapters/primary/http"
	"ecommerce-backend/internal/adapters/primary/http/middleware"
	"ecommerce-backend/internal/adapters/secondary/repository"
	"ecommerce-backend/internal/core/ports"
	"ecommerce-backend/internal/core/services"
	"ecommerce-backend/internal/infrastructure/auth"
	"ecommerce-backend/internal/infrastructure/config"
	"ecommerce-backend/internal/infrastructure/storage"
	"log"
	"time"

	_ "ecommerce-backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewFiberServer(cfg *config.Config, db *sqlx.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorHandler: customErrorHandler,
	})

	app.Use(middleware.RecoveryMiddleware())
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

	setupRoutes(app, cfg, db)

	return app
}

func setupRoutes(app *fiber.App, cfg *config.Config, db *sqlx.DB) {
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	productGalleryRepo := repository.NewProductGalleryRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	ratingRepo := repository.NewRatingRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	
	jwtService := auth.NewJWTService(cfg)
	
	var storageService ports.StorageService
	if cfg.Storage.Type == "google_drive" {
		var err error
		storageService, err = storage.NewGoogleDriveStorage(cfg)
		if err != nil {
			panic("Failed to initialize Google Drive storage: " + err.Error())
		}
	} else {
		storageService = storage.NewLocalStorage("./uploads", cfg.App.Name)
	}
	
	authService := services.NewAuthService(userRepo, jwtService)
	categoryService := services.NewCategoryService(categoryRepo, storageService)
	productService := services.NewProductService(productRepo, productGalleryRepo, storageService)
	cartService := services.NewCartService(cartRepo, productRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, productRepo)
	ratingService := services.NewRatingService(ratingRepo, orderRepo, productRepo, userRepo)
	paymentService := services.NewPaymentService(paymentRepo, orderRepo, userRepo, cfg.Xendit.APIKey)
	
	authHandler := http.NewAuthHandler(authService)
	categoryHandler := http.NewCategoryHandler(categoryService)
	productHandler := http.NewProductHandler(productService)
	cartHandler := http.NewCartHandler(cartService)
	orderHandler := http.NewOrderHandler(orderService)
	ratingHandler := http.NewRatingHandler(ratingService)
	paymentHandler := http.NewPaymentHandler(paymentService)

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

	authRoutes := v1.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/logout", authHandler.Logout)
	authRoutes.Get("/me", middleware.AuthMiddleware(jwtService), authHandler.Me)

	categoryRoutes := v1.Group("/categories")
	categoryRoutes.Get("/", categoryHandler.GetAllCategories)
	categoryRoutes.Get("/:id", categoryHandler.GetCategoryByID)
	categoryRoutes.Post("/", middleware.AuthMiddleware(jwtService), categoryHandler.CreateCategory)
	categoryRoutes.Put("/:id", middleware.AuthMiddleware(jwtService), categoryHandler.UpdateCategory)
	categoryRoutes.Delete("/:id", middleware.AuthMiddleware(jwtService), categoryHandler.DeleteCategory)

	productRoutes := v1.Group("/products")
	productRoutes.Get("/", productHandler.GetAllProducts)
	productRoutes.Get("/:id", productHandler.GetProductByID)
	productRoutes.Get("/slug/:slug", productHandler.GetProductBySlug)
	productRoutes.Get("/category/:category_id", productHandler.GetProductsByCategoryID)
	productRoutes.Post("/", middleware.AuthMiddleware(jwtService), productHandler.CreateProduct)
	productRoutes.Put("/:id", middleware.AuthMiddleware(jwtService), productHandler.UpdateProduct)
	productRoutes.Delete("/:id", middleware.AuthMiddleware(jwtService), productHandler.DeleteProduct)
	
	productRoutes.Get("/:product_id/galleries", productHandler.GetProductGalleries)
	productRoutes.Post("/galleries", middleware.AuthMiddleware(jwtService), productHandler.AddProductGallery)
	productRoutes.Put("/galleries/:id", middleware.AuthMiddleware(jwtService), productHandler.UpdateProductGallery)
	productRoutes.Delete("/galleries/:id", middleware.AuthMiddleware(jwtService), productHandler.DeleteProductGallery)

	cartRoutes := v1.Group("/cart", middleware.AuthMiddleware(jwtService))
	cartRoutes.Post("/", cartHandler.AddToCart)
	cartRoutes.Get("/", cartHandler.GetCart)
	cartRoutes.Put("/:id", cartHandler.UpdateCartItem)
	cartRoutes.Delete("/:id", cartHandler.RemoveFromCart)
	cartRoutes.Delete("/clear", cartHandler.ClearCart)

	orderRoutes := v1.Group("/orders")
	orderRoutes.Get("/statuses", orderHandler.GetOrderStatuses)
	orderRoutes.Use(middleware.AuthMiddleware(jwtService))
	orderRoutes.Post("/", orderHandler.CreateOrder)
	orderRoutes.Get("/", orderHandler.GetUserOrders)
	orderRoutes.Get("/:id", orderHandler.GetOrder)
	orderRoutes.Post("/:id/cancel", orderHandler.CancelOrder)

	adminOrderRoutes := v1.Group("/admin/orders", middleware.AuthMiddleware(jwtService))
	adminOrderRoutes.Get("/", orderHandler.GetAllOrders)
	adminOrderRoutes.Put("/:id/status", orderHandler.UpdateOrderStatus)

	ratingRoutes := v1.Group("/ratings")
	ratingRoutes.Get("/", ratingHandler.GetProductRatings)
	ratingRoutes.Get("/stats", ratingHandler.GetRatingStats)
	ratingRoutes.Use(middleware.AuthMiddleware(jwtService))
	ratingRoutes.Post("/", ratingHandler.CreateRating)
	ratingRoutes.Get("/my", ratingHandler.GetUserRatings)
	ratingRoutes.Put("/:id", ratingHandler.UpdateRating)
	ratingRoutes.Delete("/:id", ratingHandler.DeleteRating)

	paymentRoutes := v1.Group("/payments")
	paymentRoutes.Post("/xendit/callback", paymentHandler.XenditCallback)
	paymentRoutes.Use(middleware.AuthMiddleware(jwtService))
	paymentRoutes.Post("/", paymentHandler.CreatePayment)
	paymentRoutes.Get("/", paymentHandler.GetUserPayments)
	paymentRoutes.Get("/order", paymentHandler.GetPaymentByOrderID)
	paymentRoutes.Get("/:id", paymentHandler.GetPaymentByID)
	paymentRoutes.Get("/:id/status", paymentHandler.CheckPaymentStatus)
	paymentRoutes.Post("/:id/cancel", paymentHandler.CancelPayment)

	adminPaymentRoutes := v1.Group("/admin/payments", middleware.AuthMiddleware(jwtService))
	adminPaymentRoutes.Get("/", paymentHandler.GetAllPayments)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	
	// Debug logging
	log.Printf("❌ ERROR HANDLER CALLED:")
	log.Printf("   Path: %s %s", c.Method(), c.Path())
	log.Printf("   Error: %v", err)
	log.Printf("   Error Type: %T", err)

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
