package seeders

import (
	"log"

	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
)

type RoleSeeder struct{}

func (s *RoleSeeder) Seed() error {
	var count int64

	database.DB.Table("roles").Count(&count)

	if count > 0 {
		log.Println("✓ Roles already seeded, skipping...")
		return nil
	}

	roles := []map[string]interface{}{
		{
			"id":          1,
			"name":        "customer",
			"description": "Customer dengan akses belanja standar",
		},
		{
			"id":          2,
			"name":        "toko",
			"description": "Toko dengan akses manajemen produk",
		},
		{
			"id":          3,
			"name":        "admin",
			"description": "Administrator dengan akses penuh ke sistem",
		},
	}

	for _, role := range roles {
		if err := database.DB.Table("roles").Create(role).Error; err != nil {
			log.Printf("✗ Failed to seed role %s: %v", role["name"], err)
			return err
		}
	}

	log.Println("✓ Roles seeded successfully:")
	log.Println("  - ID 1: customer")
	log.Println("  - ID 2: toko")
	log.Println("  - ID 3: admin")

	return nil
}
