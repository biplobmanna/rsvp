package rsvp

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddUrls(app *fiber.App, db *gorm.DB) {
	// Add the URLs
	app.Get("/", IndexView).Name("index")
	app.Get("/whoami", WhoAmIView).Name("whoami")
	app.Get("/card", CardView).Name("card")

	app.Post("/check-whoami", CheckWhoAmI).Name("check-whoami")
}
