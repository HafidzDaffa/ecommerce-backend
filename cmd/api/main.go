package main

import (
	"ecommerce-backend/internal/infrastructure/config"
	"ecommerce-backend/internal/infrastructure/database"
	"ecommerce-backend/internal/infrastructure/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

// @title E-Commerce Backend API
// @version 1.0
// @description Complete e-commerce backend server with authentication, product management, category management, and Google Drive integration.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@ecommerce.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	app := server.NewFiberServer(cfg, db)

	go func() {
		if err := app.Listen(":" + cfg.App.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server is running on port %s", cfg.App.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
