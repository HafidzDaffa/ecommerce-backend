package seeders

import (
	"log"

	"github.com/yourusername/ecommerce-go-vue/backend/infrastructure/database"
)

type TagSeeder struct{}

func (s *TagSeeder) Seed() error {
	var count int64

	database.DB.Table("tags").Count(&count)

	if count > 0 {
		log.Println("✓ Tags already seeded, skipping...")
		return nil
	}

	tags := []map[string]interface{}{
		{
			"id":   1,
			"name": "Discount",
			"slug": "discount",
		},
		{
			"id":   2,
			"name": "Sale",
			"slug": "sale",
		},
		{
			"id":   3,
			"name": "Best Seller",
			"slug": "best-seller",
		},
		{
			"id":   4,
			"name": "New",
			"slug": "new",
		},
		{
			"id":   5,
			"name": "Flash Sale",
			"slug": "flash-sale",
		},
		{
			"id":   6,
			"name": "Premium",
			"slug": "premium",
		},
		{
			"id":   7,
			"name": "Free Shipping",
			"slug": "free-shipping",
		},
		{
			"id":   8,
			"name": "Local",
			"slug": "local",
		},
		{
			"id":   9,
			"name": "Import",
			"slug": "import",
		},
		{
			"id":   10,
			"name": "Limited Edition",
			"slug": "limited-edition",
		},
		{
			"id":   11,
			"name": "Exclusive",
			"slug": "exclusive",
		},
		{
			"id":   12,
			"name": "Trending",
			"slug": "trending",
		},
		{
			"id":   13,
			"name": "Eco-Friendly",
			"slug": "eco-friendly",
		},
		{
			"id":   14,
			"name": "Organic",
			"slug": "organic",
		},
		{
			"id":   15,
			"name": "Handmade",
			"slug": "handmade",
		},
		{
			"id":   16,
			"name": "Vintage",
			"slug": "vintage",
		},
		{
			"id":   17,
			"name": "Clearance",
			"slug": "clearance",
		},
		{
			"id":   18,
			"name": "Bundle",
			"slug": "bundle",
		},
		{
			"id":   19,
			"name": "Top Rated",
			"slug": "top-rated",
		},
		{
			"id":   20,
			"name": "Featured",
			"slug": "featured",
		},
		{
			"id":   21,
			"name": "Hot Deal",
			"slug": "hot-deal",
		},
		{
			"id":   22,
			"name": "Cashback",
			"slug": "cashback",
		},
		{
			"id":   23,
			"name": "Buy 1 Get 1",
			"slug": "buy-1-get-1",
		},
		{
			"id":   24,
			"name": "Wholesale",
			"slug": "wholesale",
		},
		{
			"id":   25,
			"name": "Pre-order",
			"slug": "pre-order",
		},
	}

	for _, tag := range tags {
		if err := database.DB.Table("tags").Create(tag).Error; err != nil {
			log.Printf("✗ Failed to seed tag %s: %v", tag["name"], err)
			return err
		}
	}

	log.Println("✓ Tags seeded successfully:")
	log.Println("  - Discount")
	log.Println("  - Sale")
	log.Println("  - Best Seller")
	log.Println("  - New")
	log.Println("  - Flash Sale")
	log.Println("  - Premium")
	log.Println("  - Free Shipping")
	log.Println("  - Local")
	log.Println("  - Import")
	log.Println("  - Limited Edition")
	log.Println("  - Exclusive")
	log.Println("  - Trending")
	log.Println("  - Eco-Friendly")
	log.Println("  - Organic")
	log.Println("  - Handmade")
	log.Println("  - Vintage")
	log.Println("  - Clearance")
	log.Println("  - Bundle")
	log.Println("  - Top Rated")
	log.Println("  - Featured")
	log.Println("  - Hot Deal")
	log.Println("  - Cashback")
	log.Println("  - Buy 1 Get 1")
	log.Println("  - Wholesale")
	log.Println("  - Pre-order")

	return nil
}
