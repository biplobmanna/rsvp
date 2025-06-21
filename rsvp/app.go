package rsvp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

// Separate function to initialize App, to modularize
// and to test
func InitApp(s Settings) *fiber.App {
	engine := django.New(s.TEMPLATE_DIR, s.TEMPLATE_EXTENSION)

	var app *fiber.App
	app = fiber.New(fiber.Config{
		Views: engine,
	})

	return app
}

func App() *fiber.App {
	// Build Configuration Settings
	s := Settings{}
	s.BuildConf()

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
