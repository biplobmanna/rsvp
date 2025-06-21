package rsvp

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DATABASE_NAME = "rsvp.db"

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DATABASE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DestroyDB() {
	err := os.Remove(DATABASE_NAME)
	if err != nil {
		log.Fatal(err)
	}
}
