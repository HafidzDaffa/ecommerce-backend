package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/config"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
	"github.com/yourusername/ecommerce-go-vue/backend/seeders"
)

var (
	migrateCmd = flag.NewFlagSet("migrate", flag.ExitOnError)
	seedCmd    = flag.NewFlagSet("seed", flag.ExitOnError)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/migrate.go [command]")
		fmt.Println("Commands:")
		fmt.Println("  migrate - Run database migrations")
		fmt.Println("  seed    - Run database seeders")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "migrate":
		runMigrate()
	case "seed":
		runSeed()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

func runMigrate() {
	cfg := config.LoadConfig()

	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	log.Println("Running database migrations...")

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		),
	)

	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	currentVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Failed to get migration version: %v", err)
	}

	if err == migrate.ErrNilVersion {
		log.Println("No migrations found. Database is clean.")
	} else {
		log.Printf("Current migration version: %d, dirty: %v", currentVersion, dirty)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err == migrate.ErrNoChange {
		log.Println("✓ Database is already up to date")
	} else {
		log.Println("✓ Migrations completed successfully")
	}
}

func runSeed() {
	cfg := config.LoadConfig()

	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	log.Println("Running database seeders...")

	if err := seeders.RunSeeders(); err != nil {
		log.Fatalf("Failed to run seeders: %v", err)
	}
}
