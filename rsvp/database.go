package rsvp

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DATABASE_NAME = "rsvp.db"

// global DB connection pool
// this needs to be initialized in App()
// without initialization, the entire functionality will not work
var DB *gorm.DB

func ConnectDB() *gorm.DB {
	fmt.Println("⚫ Connecting to database...")
	var err error
	// using the same global variable as above
	// to make it accessible across all func
	// in the same package
	DB, err = gorm.Open(sqlite.Open(DATABASE_NAME), &gorm.Config{})
	if err != nil {
		fmt.Println("  ⚫ Failed to connect ❌")
		log.Fatal(err)
	}

	fmt.Println("  ⚫ Connected OK ✔️")
	return DB
}

func DestroyDB() {
	err := os.Remove(DATABASE_NAME)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("🔴 Deleting ", DATABASE_NAME, "...")
	}
}
