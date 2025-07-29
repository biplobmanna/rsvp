package rsvp

import (
	"fmt"

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

	// Migrate Refresh And connect DB
	// Use when you want a fresh DB
	// DB := MigrateRefreshAndConnectDB()
	// SeedAdmin(s, DB)
	// seedUser(s, DB)

	// Use for normal situations
	DB := ConnectDB()
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
