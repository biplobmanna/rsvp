package rsvp

import (
	"fmt"
	"log"

	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
)

// Seed to MAX SEEDERS
var SEEDERS = 100

func SeedAdmin(s Settings, db *gorm.DB) {
	fake := faker.New()
	errors, rowsAffected := 0, 0
	for range 2 {
		user := fake.Internet().User()
		token := s.ADMIN_TOKEN
		admin := Admin{
			User:  user,
			Token: token,
		}

		result := db.Create(&admin)
		if result.Error != nil {
			errors += 1
		} else {
			rowsAffected += 1
		}
	}
	fmt.Println("🔵 Seeding Admin:")
	fmt.Println("  🔵 Rows Affected: ", rowsAffected)
	fmt.Println("  🔵 Errors While Seeding: ", errors)
}

func seedUser(s Settings, db *gorm.DB) {
	fake := faker.New()
	errors, rowsAffected := 0, 0
	for i := range SEEDERS {
		username := fake.Internet().User()
		token, err := EncryptAES(s, username+s.ADMIN_TOKEN)
		if err != nil {
			log.Fatal(err)
		}

		user := User{
			User:     username,
			FullName: fake.Person().Name(),
			Token:    token,
			Email:    fake.Internet().Email(),
			Phone:    "+91-9123456789",
			Rsvp:     false,
			Comments: fake.Lorem().Sentence(i),
		}
		result := db.Create(&user)

		if result.Error != nil {
			errors += 1
		} else {
			rowsAffected += 1
		}
	}
	fmt.Println("🔵 Seeding User:")
	fmt.Println("  🔵 Rows Affected: ", rowsAffected)
	fmt.Println("  🔵 Errors While Seeding: ", errors)
}
