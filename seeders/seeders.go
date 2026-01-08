package seeders

import (
	"log"
)

type Seeder interface {
	Seed() error
}

func RunSeeders() error {
	seeders := []Seeder{
		&RoleSeeder{},
		&AdminSeeder{},
	}

	for _, seeder := range seeders {
		if err := seeder.Seed(); err != nil {
			return err
		}
	}

	log.Println("All seeders completed successfully")
	return nil
}
