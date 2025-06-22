package rsvp

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Separate function to initialize App, to modularize
// and to test
func InitApp(s Settings) *fiber.App {
	engine := html.New(s.TEMPLATE_DIR, s.TEMPLATE_EXTENSION)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	return app
}

func App() *fiber.App {
	// Build Configuration Settings
	s := Settings{}
	s.BuildConf()

	// Migrate Refresh And connect DB
	// Use when you want a fresh DB
	// DB := MigrateRefreshAndConnectDB()
	// SeedAdmin(s, DB)
	// seedUser(s, DB)

	// Use for normal situations
	DB := ConnectDB()
	fmt.Println("  ⚫ DB:", DB.Name())

	// Initialise the App
	app := InitApp(s)

	// Apply any custom s here, as needed
	AddStatic(app, s)

	// Add the URLs
	AddUrls(app)

	// Return the App instance
	return app
}

func AddStatic(app *fiber.App, s Settings) {
	app.Static(s.STATIC_URL, s.STATIC_DIR, fiber.Static{
		Compress:      s.STATIC_COMPRESS,
		ByteRange:     s.STATIC_BYTE_RANGE,
		Browse:        s.STATIC_BROWSE,
		Download:      s.STATIC_DOWNLOAD,
		Index:         s.STATIC_INDEX,
		CacheDuration: s.STATIC_CACHE_DURATION,
	})
}
