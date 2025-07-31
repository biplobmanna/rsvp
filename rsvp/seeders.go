package rsvp

import (
	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
)

// Seed to MAX SEEDERS
var SEEDERS = 100

func seedUser(db *gorm.DB) {
	//LOG.Println("Seeding User...")
	fake := faker.New()
	errors, rowsAffected := 0, 0
	//LOG.Println("Generating", SEEDERS, "seeds for Users...")
	for i := range SEEDERS {
		username := fake.Internet().User()
		token, err := EncryptAES(username+SETTINGS.ADMIN_TOKEN)
		if err != nil {
			//LOG.Fatal(err)
		}

		user := User{
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
	//LOG.Println("  🔵 Rows Affected: ", rowsAffected)
	//LOG.Println("  🔵 Errors While Seeding: ", errors)
}
