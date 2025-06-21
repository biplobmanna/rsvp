package rsvp

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("rsvp.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
