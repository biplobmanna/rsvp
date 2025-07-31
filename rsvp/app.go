package rsvp

import (
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Separate function to initialize App, to modularize
// and to test
func InitApp() *fiber.App {
	engine := html.New(SETTINGS.TEMPLATE_DIR, SETTINGS.TEMPLATE_EXTENSION)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	return app
}

func App() *fiber.App {
	// Build Configuration Settings
	SETTINGS = Settings{}
	SETTINGS.BuildConf()

	// check if DB has previously existed
	isDBExist := CheckDBExists(DATABASE_NAME)

	if isDBExist {
		// Use for normal situations
		DB = ConnectDB()
	} else {
		// create the DB file since it does not exist
		file, _ := os.Create(DATABASE_NAME)
		defer file.Close()
		// Migrate Refresh And connect DB
		DB = MigrateRefreshAndConnectDB()
		// Use when you want a fresh DB
		// SeedAdmin(SETTINGS, DB)
		// seedUser(SETTINGS, DB)
	}
	fmt.Println("  ⚫ DB:", DB.Name())

	// Initialise the App
	app := InitApp()

	// Apply any custom s here, as needed
	AddStatic(app)

	// Add the URLs
	AddUrls(app)

	// Return the App instance
	return app
}

func AddStatic(app *fiber.App) {
	app.Static(SETTINGS.STATIC_URL, SETTINGS.STATIC_DIR, fiber.Static{
		Compress:      SETTINGS.STATIC_COMPRESS,
		ByteRange:     SETTINGS.STATIC_BYTE_RANGE,
		Browse:        SETTINGS.STATIC_BROWSE,
		Download:      SETTINGS.STATIC_DOWNLOAD,
		Index:         SETTINGS.STATIC_INDEX,
		CacheDuration: SETTINGS.STATIC_CACHE_DURATION,
	})
}
