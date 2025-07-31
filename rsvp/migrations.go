package rsvp

import (
	"gorm.io/gorm"
)

func Migrate[M Model](db *gorm.DB, m *M, desc ...string) {
	printDesc := "database table..."
	if len(desc) > 0 {
		printDesc = desc[0]
	}
	LOG.Println("  🟣 Migrating ", printDesc)
	db.AutoMigrate(&m)
}

func MigrateAll(db *gorm.DB) {
	LOG.Println("🟣 Running All Migrations...")
	Migrate(db, &User{}, "User")
}

func MigrateRefreshAndConnectDB() *gorm.DB {
	LOG.Println("Deleting DB...")
	DestroyDB()
	db := ConnectDB()
	MigrateAll(db)
	return db
}
