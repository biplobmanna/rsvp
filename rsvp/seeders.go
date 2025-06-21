package rsvp

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
)

// Seed to MAX SEEDERS
var SEEDERS = 100

func SeedAdmin(s Settings, db *gorm.DB) {
	fake := faker.New()
	errors, rowsAffected := 0, 0
	for i := range 10 {
		user := fake.Internet().User()
		token, err := EncryptAES(s, user+strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		admin := Admin{
			User:  user,
			Token: token,
		}

		result := db.Create(admin)
		if result.Error != nil {
			errors += 1
		} else {
			rowsAffected += 1
		}
	}
	fmt.Println("Seeding Admin:")
	fmt.Println("Rows Affected: ", rowsAffected)
	fmt.Println("Errors While Seeding: ", errors)
}

func seedUser(s Settings, db *gorm.DB) {
	fake := faker.New()
	errors, rowsAffected := 0, 0
	for i := range SEEDERS {
		username := fake.Internet().User()
		token, err := EncryptAES(s, username+strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		user := User{
			User:     username,
			Token:    token,
			FullName: fake.Person().Name(),
			Rsvp:     false,
			Comments: "",
		}
		result := db.Create(user)

		if result.Error != nil {
			errors += 1
		} else {
			rowsAffected += 1
		}
	}
	fmt.Println("Seeding User:")
	fmt.Println("Rows Affected: ", rowsAffected)
	fmt.Println("Errors While Seeding: ", errors)
}
