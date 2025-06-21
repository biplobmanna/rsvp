package rsvp

import (
	"gorm.io/gorm"
)

func Migrate[M Model](db *gorm.DB, m *M) {
	db.AutoMigrate(&m)
}

func MigrateAll(db *gorm.DB) {
	Migrate(db, &Admin{})
	Migrate(db, &User{})
}
