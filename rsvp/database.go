package rsvp

import (
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
	LOG.Println("⚫ Connecting to database...")
	var err error
	DB, err = gorm.Open(sqlite.Open(DATABASE_NAME), &gorm.Config{})
	if err != nil {
		LOG.Panicln("  ⚫ Failed to connect to DB... ❌")
	}
	LOG.Println("  ⚫ DB Connected OK ✔️")
	return DB
}

func DestroyDB() {
	LOG.Println("Trying to remove Database:", DATABASE_NAME, "...")
	err := os.Remove(DATABASE_NAME)
	if err != nil {
		LOG.Panicln("Failed to remove DB...", err)
	} else {
		LOG.Println("🔴 Deleting ", DATABASE_NAME, "...")
	}
}

// check if the DB file exists
func CheckDBExists(path string) bool {
	LOG.Println("Checking if File", path, "Exist...")
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
