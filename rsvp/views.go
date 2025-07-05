package rsvp

import (
	"github.com/gofiber/fiber/v2"
)

// GET: Form to validate the token |
// POST: Validate the token, and redirect to /card
func WhoAmIView(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("whoami", fiber.Map{
			"Title": "RSVP",
		}, "base")
	} else if c.Method() == "POST" {
		whoami := new(WhoAmI)
		if err := c.BodyParser(whoami); err != nil {
			return c.Status(fiber.StatusBadRequest).Redirect("/whoami")
		}
		// check if token is valid, and take actions accordingly
		if whoami.ValidateToken() {
			SetTokenCookie(c, whoami.Token)
			return c.Redirect("/card")
		} else {
			return c.Redirect("/whoami")
		}
	}
	return c.Redirect("/whoami")
}

// GET: Show the card upon successful validation
func CardView(c *fiber.Ctx) error {
	whoami, err := GetTokenCookie(c)
	if err != nil {
		return err
	}

	if whoami.ValidateToken() {
		SetTokenCookie(c, whoami.Token)
		return c.Render("card", fiber.Map{
			"Title": "RSVP CARD",
		}, "card-base")
	} else {
		return c.Redirect("/whoami")
	}
}

// GET: RSVP action for the user
func RsvpView(c *fiber.Ctx) error {
	return c.Render("rsvp", fiber.Map{
		"Title": "RSVP",
	}, "base")
}

// ALL: Redirect User to GET: /whoami page
func RedirectToWhoAmI(c *fiber.Ctx) error {
	return c.Redirect("/whoami")
}
