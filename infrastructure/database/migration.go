package database

import (
	"log"

	"github.com/yourusername/ecommerce-go-vue/backend/domain/entities"
)

func AutoMigrate() error {
	err := DB.AutoMigrate(
		&entities.Role{},
		&entities.User{},
	)

	if err != nil {
		log.Printf("Auto migration failed: %v", err)
		return err
	}

	log.Println("Database migration completed successfully")

	return nil
}

func SeedData() error {
	var roleCount int64
	DB.Model(&entities.Role{}).Count(&roleCount)

	if roleCount == 0 {
		roles := []entities.Role{
			{ID: 1, Name: "customer", Description: "Regular customer with standard purchasing access"},
			{ID: 2, Name: "toko", Description: "Store owner with product management access"},
			{ID: 3, Name: "admin", Description: "Administrator with full system access"},
		}

		for _, role := range roles {
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Failed to create role %s: %v", role.Name, err)
				continue
			}
		}

		log.Println("Default roles seeded successfully: customer (1), toko (2), admin (3)")
	}

	return nil
}
