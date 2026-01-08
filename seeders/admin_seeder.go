package seeders

import (
	"log"

	"github.com/yourusername/ecommerce-go-vue/backend/common/utils"
	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
)

type AdminSeeder struct{}

func (s *AdminSeeder) Seed() error {
	var count int64

	database.DB.Table("users").Where("role_id = ?", 3).Count(&count)

	if count > 0 {
		log.Println("✓ Admin user already seeded, skipping...")
		return nil
	}

	passwordHash, err := utils.HashPassword("12341234")
	if err != nil {
		log.Printf("✗ Failed to hash password: %v", err)
		return err
	}

	adminUser := map[string]interface{}{
		"email":         "admin@ecommerce.com",
		"password_hash": passwordHash,
		"full_name":     "Administrator",
		"role_id":       3,
		"is_active":     true,
		"is_verified":   true,
	}

	if err := database.DB.Table("users").Create(adminUser).Error; err != nil {
		log.Printf("✗ Failed to seed admin user: %v", err)
		return err
	}

	log.Println("✓ Admin user seeded successfully:")
	log.Println("  - Email: admin@ecommerce.com")
	log.Println("  - Password: 12341234")
	log.Println("  - Role: Admin")

	return nil
}
