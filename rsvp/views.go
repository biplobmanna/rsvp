package rsvp

import (
	"github.com/gofiber/fiber/v2"
)

// The views are used to serve HTML pages
func IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title":    "Base HTML",
		"Contents": "Hello, World!",
	}, "base")
}

func WhoAmIView(c *fiber.Ctx) error {
	return c.Render("whoami", fiber.Map{
		"title": "RSVP",
	}, "base")
}

func CheckWhoAmI(c *fiber.Ctx) error {
	whoami := new(WhoAmI)
	if err := c.BodyParser(whoami); err != nil {
		return err
	}
	// check if token is valid, and take actions accordingly
	if whoami.ValidateToken() {
		SetTokenCookie(c, whoami.Token)
		return c.Redirect("/card")
	} else {
		return c.Redirect("/whoami")
	}
}

func CardView(c *fiber.Ctx) error {
	var whoami WhoAmI
	var err error
	err, whoami = GetTokenCookie(c)
	if err != nil {
		return err
	}

	if whoami.ValidateToken() {
		SetTokenCookie(c, whoami.Token)
		return c.Render("card", fiber.Map{
			"title": "RSVP CARD",
		}, "base")
	} else {
		return c.Redirect("/whoami")
	}
}
