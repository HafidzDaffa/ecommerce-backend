package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Hash dari seeder
	hashFromSeeder := "$2a$10$Z5UnNu2kAT3qgLCWbhyhgu7NCamR6X/9SKl.fd8/I9U/mOTjWQv76"
	password := "password123"
	
	// Test verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(hashFromSeeder), []byte(password))
	if err == nil {
		fmt.Println("✅ Password COCOK! Hash di seeder valid untuk 'password123'")
	} else {
		fmt.Println("❌ Password TIDAK COCOK!")
		fmt.Printf("Error: %v\n", err)
		
		// Generate hash baru yang benar
		newHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		fmt.Printf("\n🔧 Hash yang benar untuk 'password123':\n%s\n", string(newHash))
	}
}
