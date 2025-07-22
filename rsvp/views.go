package rsvp

import (
	"github.com/gofiber/fiber/v2"
)

// GET: Form to validate the token |
// POST: Validate the token, and redirect to /card
func WhoAmIView(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		// handle the GET request

		// return the HTML template form to enter token
		return c.Render("whoami", fiber.Map{
			"Title": "RSVP",
		}, "base")

	} else if c.Method() == "POST" {
		// handle the POST request

		// declare a WhoAmI struct to hold the token data
		// fetched from the request
		whoami := new(WhoAmI)
		// parse request body for the data in WhoAmI
		if err := c.BodyParser(whoami); err != nil {
			// if error in parsing, redirect back to the same page
			// this should ensure that the input page is shown again
			return c.Status(fiber.StatusBadRequest).Redirect("/whoami")
		}
		// check if token is valid, and take actions accordingly
		if whoami.ValidateToken() {
			// Set cookie with the token
			SetTokenCookie(c, whoami.Token)
			// redirect to the Card View
			return c.Redirect("/card")

		} else {
			// redirect is back to the same page to enter
			// the token once again
			return c.Redirect("/whoami")
		}
	}
	// this is a path flow, and redirects to the token form page
	return c.Redirect("/whoami")
}

// GET: Show the card upon successful validation
func CardView(c *fiber.Ctx) error {
	// first, fetch the cookie and extract token
	whoami, err := GetTokenCookie(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
	}

	if whoami.ValidateToken() {
		// if token is valid, 
		// set token in the Cookie
		SetTokenCookie(c, whoami.Token)
		
		// return HTML template for teh RSVP card
		return c.Render("card", fiber.Map{
			"Title": "RSVP CARD",
		}, "card-base")

	} else {
		// in case of invalid token, redirect to the token form page
		return c.Redirect("/whoami")
	}
}

// GET: RSVP action for the user
// | return the HTML template with Card
func RsvpView(c *fiber.Ctx) error {
	return c.Render("rsvp", fiber.Map{
		"Title": "RSVP",
	}, "base")
}

// ALL: Redirect User to GET: /whoami page
// a catch-all view to catch all misc URLs and redirect
// instead of throwing errors
func RedirectToWhoAmI(c *fiber.Ctx) error {
	return c.Redirect("/whoami")
}
