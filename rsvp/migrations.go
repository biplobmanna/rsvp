package rsvp

import (
	"gorm.io/gorm"
)

func Migrate[M Model](db *gorm.DB, m *M, desc ...string) {
	//LOG.Println("  🟣 Migrating Table(s)...")
	db.AutoMigrate(&m)
}

func MigrateAll(db *gorm.DB) {
	//LOG.Println("🟣 Running All Migrations...")
	Migrate(db, &User{}, "User")
}

func MigrateRefreshAndConnectDB() *gorm.DB {
	//LOG.Println("Deleting DB...")
	DestroyDB()
	db := ConnectDB()
	MigrateAll(db)
	return db
}
