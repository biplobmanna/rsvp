package rsvp

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Separate function to initialize App, to modularize
// and to test
func InitApp() *fiber.App {
	//LOG.Println("Setting the HTML Template Engine...")
	engine := html.New(SETTINGS.TEMPLATE_DIR, SETTINGS.TEMPLATE_EXTENSION)

	//LOG.Println("Initialising a Fiber Application...")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return app
}

func App() *fiber.App {
	// Build Configuration Settings
	SETTINGS = Settings{}
	SETTINGS.BuildConf()

	// Setting up //LOG.ING
	SetupLogging()

	//LOG.Println("Setup of Settings and Logging Complete...")
	//LOG.Printf("Checking if DATABSE:%s exists...\n", DATABASE_NAME)
	isDBExist := CheckDBExists(DATABASE_NAME)

	if isDBExist {
		//LOG.Println("DB Exists, connecting...")
		DB = ConnectDB()
	} else {
		//LOG.Println("DB does not exist, creating...")
		file, err := os.Create(DATABASE_NAME)
		if err != nil {
			//LOG.Panicln("Failed to create DB, Exiting...")
		}
		defer file.Close()

		//LOG.Println("DB Created, connecting...")
		//LOG.Println("Migrating the Tables...")
		DB = MigrateRefreshAndConnectDB()

		// Use when you want a fresh DB, and seeded for testing
		// SeedAdmin(DB)
		// seedUser(DB)
	}

	//LOG.Println("Initialising Application...")
	app := InitApp()

	//LOG.Println("Adding Static Settings...")
	AddStatic(app)

	//LOG.Println("Addding the URLs...")
	AddUrls(app)

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
