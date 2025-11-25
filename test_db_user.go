package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("✅ Database connected successfully")

	// Query user seller@ecommerce.com
	var id, email, fullName, passwordHash string
	var roleID int
	var isActive bool

	err = db.QueryRow(`
		SELECT id, email, full_name, password_hash, role_id, is_active 
		FROM users 
		WHERE email = $1
	`, "seller@ecommerce.com").Scan(&id, &email, &fullName, &passwordHash, &roleID, &isActive)

	if err == sql.ErrNoRows {
		fmt.Println("❌ User 'seller@ecommerce.com' TIDAK DITEMUKAN di database!")
		fmt.Println("\n📝 Jalankan seeder untuk menambahkan user:")
		fmt.Println("   psql -U postgres -d ecommerce_db -f seeders/004_users.sql")
		return
	} else if err != nil {
		log.Fatal("Query error:", err)
	}

	fmt.Println("\n✅ User DITEMUKAN di database:")
	fmt.Printf("   ID: %s\n", id)
	fmt.Printf("   Email: %s\n", email)
	fmt.Printf("   Full Name: %s\n", fullName)
	fmt.Printf("   Role ID: %d\n", roleID)
	fmt.Printf("   Is Active: %t\n", isActive)
	fmt.Printf("   Password Hash: %s...\n", passwordHash[:30])

	// Test password verification
	password := "password123"
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err == nil {
		fmt.Println("\n✅ Password 'password123' COCOK dengan hash di database!")
		fmt.Println("\n🎯 Kesimpulan: Data user dan password VALID, login seharusnya berhasil!")
	} else {
		fmt.Println("\n❌ Password 'password123' TIDAK COCOK!")
		fmt.Println("   Error:", err)
	}

	// Check all users
	fmt.Println("\n📋 Semua user di database:")
	rows, _ := db.Query("SELECT email, full_name, role_id, is_active FROM users ORDER BY email")
	defer rows.Close()
	
	for rows.Next() {
		var e, fn string
		var rid int
		var ia bool
		rows.Scan(&e, &fn, &rid, &ia)
		fmt.Printf("   - %s (%s) [Role: %d, Active: %t]\n", e, fn, rid, ia)
	}
}
