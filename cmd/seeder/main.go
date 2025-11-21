package main

import (
	"ecommerce-backend/internal/infrastructure/config"
	"ecommerce-backend/internal/infrastructure/database"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	seedersPath := "seeders"

	files, err := os.ReadDir(seedersPath)
	if err != nil {
		log.Fatalf("Failed to read seeders directory: %v", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}

	sort.Strings(sqlFiles)

	if len(sqlFiles) == 0 {
		log.Println("No seeder files found")
		return
	}

	log.Printf("Found %d seeder file(s)\n", len(sqlFiles))

	for _, fileName := range sqlFiles {
		filePath := filepath.Join(seedersPath, fileName)
		
		log.Printf("Running seeder: %s", fileName)
		
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("❌ Failed to read %s: %v", fileName, err)
			continue
		}

		_, err = db.Exec(string(content))
		if err != nil {
			log.Printf("❌ Failed to execute %s: %v", fileName, err)
			continue
		}

		log.Printf("✅ Successfully executed: %s", fileName)
	}

	log.Println("\n✅ All seeders completed successfully!")
	
	printSummary(db)
}

func printSummary(db interface {
	Get(dest interface{}, query string, args ...interface{}) error
}) {
	log.Println("\n📊 Database Summary:")
	log.Println("===================")

	tables := []struct {
		name  string
		table string
	}{
		{"Roles", "roles"},
		{"Order Statuses", "order_statuses"},
		{"Categories", "categories"},
	}

	for _, t := range tables {
		var count int
		err := db.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s", t.table))
		if err != nil {
			log.Printf("⚠️  %s: Error counting - %v", t.name, err)
			continue
		}
		log.Printf("📦 %s: %d records", t.name, count)
	}
}
