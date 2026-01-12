package seeders

import (
	"log"

	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
)

type CategorySeeder struct{}

func (s *CategorySeeder) Seed() error {
	var count int64

	database.DB.Table("categories").Count(&count)

	if count > 0 {
		log.Println("✓ Categories already seeded, skipping...")
		return nil
	}

	categories := []map[string]interface{}{
		{
			"id":          1,
			"name":        "Electronics",
			"slug":        "electronics",
			"description": "Electronic devices and gadgets",
			"image_url":   "/images/categories/electronics.jpg",
			"is_active":   true,
		},
		{
			"id":          2,
			"name":        "Men's Fashion",
			"slug":        "mens-fashion",
			"description": "Clothing and accessories for men",
			"image_url":   "/images/categories/mens-fashion.jpg",
			"is_active":   true,
		},
		{
			"id":          3,
			"name":        "Women's Fashion",
			"slug":        "womens-fashion",
			"description": "Clothing and accessories for women",
			"image_url":   "/images/categories/womens-fashion.jpg",
			"is_active":   true,
		},
		{
			"id":          4,
			"name":        "Health & Beauty",
			"slug":        "health-beauty",
			"description": "Health and beauty products",
			"image_url":   "/images/categories/health-beauty.jpg",
			"is_active":   true,
		},
		{
			"id":          5,
			"name":        "Sports & Outdoor",
			"slug":        "sports-outdoor",
			"description": "Sports equipment and outdoor activities",
			"image_url":   "/images/categories/sports-outdoor.jpg",
			"is_active":   true,
		},
		{
			"id":          6,
			"name":        "Home & Living",
			"slug":        "home-living",
			"description": "Home appliances and decorations",
			"image_url":   "/images/categories/home-living.jpg",
			"is_active":   true,
		},
		{
			"id":          7,
			"name":        "Books & Stationery",
			"slug":        "books-stationery",
			"description": "Books and office stationery",
			"image_url":   "/images/categories/books-stationery.jpg",
			"is_active":   true,
		},
		{
			"id":          8,
			"name":        "Automotive",
			"slug":        "automotive",
			"description": "Vehicle accessories and care products",
			"image_url":   "/images/categories/automotive.jpg",
			"is_active":   true,
		},
		{
			"id":          9,
			"name":        "Baby & Kids",
			"slug":        "baby-kids",
			"description": "Products for babies and children",
			"image_url":   "/images/categories/baby-kids.jpg",
			"is_active":   true,
		},
		{
			"id":          10,
			"name":        "Food & Beverages",
			"slug":        "food-beverages",
			"description": "Food and drinks",
			"image_url":   "/images/categories/food-beverages.jpg",
			"is_active":   true,
		},
		{
			"id":          11,
			"name":        "Computer & Accessories",
			"slug":        "computer-accessories",
			"description": "Computers and computer accessories",
			"image_url":   "/images/categories/computer-accessories.jpg",
			"is_active":   true,
		},
		{
			"id":          12,
			"name":        "Toys & Hobbies",
			"slug":        "toys-hobbies",
			"description": "Toys and hobby products",
			"image_url":   "/images/categories/toys-hobbies.jpg",
			"is_active":   true,
		},
		{
			"id":          13,
			"name":        "Jewelry & Watches",
			"slug":        "jewelry-watches",
			"description": "Jewelry and watch products",
			"image_url":   "/images/categories/jewelry-watches.jpg",
			"is_active":   true,
		},
		{
			"id":          14,
			"name":        "Pet Supplies",
			"slug":        "pet-supplies",
			"description": "Pet food and supplies",
			"image_url":   "/images/categories/pet-supplies.jpg",
			"is_active":   true,
		},
		{
			"id":          15,
			"name":        "Home Appliances",
			"slug":        "home-appliances",
			"description": "Electronic home appliances",
			"image_url":   "/images/categories/home-appliances.jpg",
			"is_active":   true,
		},
		{
			"id":          16,
			"name":        "Furniture",
			"slug":        "furniture",
			"description": "Home and office furniture",
			"image_url":   "/images/categories/furniture.jpg",
			"is_active":   true,
		},
		{
			"id":          17,
			"name":        "Garden & Plants",
			"slug":        "garden-plants",
			"description": "Gardening tools and plants",
			"image_url":   "/images/categories/garden-plants.jpg",
			"is_active":   true,
		},
		{
			"id":          18,
			"name":        "Music & Instruments",
			"slug":        "music-instruments",
			"description": "Musical instruments and accessories",
			"image_url":   "/images/categories/music-instruments.jpg",
			"is_active":   true,
		},
		{
			"id":          19,
			"name":        "Camera & Photo",
			"slug":        "camera-photo",
			"description": "Cameras and photography equipment",
			"image_url":   "/images/categories/camera-photo.jpg",
			"is_active":   true,
		},
		{
			"id":          20,
			"name":        "Gaming",
			"slug":        "gaming",
			"description": "Video games and gaming accessories",
			"image_url":   "/images/categories/gaming.jpg",
			"is_active":   true,
		},
	}

	for _, category := range categories {
		if err := database.DB.Table("categories").Create(category).Error; err != nil {
			log.Printf("✗ Failed to seed category %s: %v", category["name"], err)
			return err
		}
	}

	log.Println("✓ Categories seeded successfully:")
	log.Println("  - Electronics")
	log.Println("  - Men's Fashion")
	log.Println("  - Women's Fashion")
	log.Println("  - Health & Beauty")
	log.Println("  - Sports & Outdoor")
	log.Println("  - Home & Living")
	log.Println("  - Books & Stationery")
	log.Println("  - Automotive")
	log.Println("  - Baby & Kids")
	log.Println("  - Food & Beverages")
	log.Println("  - Computer & Accessories")
	log.Println("  - Toys & Hobbies")
	log.Println("  - Jewelry & Watches")
	log.Println("  - Pet Supplies")
	log.Println("  - Home Appliances")
	log.Println("  - Furniture")
	log.Println("  - Garden & Plants")
	log.Println("  - Music & Instruments")
	log.Println("  - Camera & Photo")
	log.Println("  - Gaming")

	return nil
}
