package rsvp

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate[M Model](db *gorm.DB, m *M) {
	fmt.Println("Migrating ", m)
	db.AutoMigrate(&m)
}

func MigrateAll(db *gorm.DB) {
	Migrate(db, &Admin{})
	Migrate(db, &User{})
}

func MigrateRefreshAndConnectDB() *gorm.DB {
	DestroyDB()
	db := ConnectDB()
	MigrateAll(db)
	return db
}
