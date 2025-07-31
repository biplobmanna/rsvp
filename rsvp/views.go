package rsvp

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ---- HELPER FUNCTIONS ----

// extract token from query param or cookie |
// validate the token and return validation status |
// on successful validation, get the user as well
// also return the struct with the token
func extractTokenFromQueryOrCookieAndValidate(c *fiber.Ctx) (bool, WhoAmI, User) {
	isTokenValid := false
	user := User{}

	//LOG.Println("Check the URL QueryParams for Valid Token...")
	whoami := GetTokenQuery(c)
	isTokenValid, user = whoami.ValidateTokenAndGetUser()
	if isTokenValid {
		//LOG.Println("URL QueryParams Token is valid...")
		return isTokenValid, whoami, user
	}

	//LOG.Println("URL QueryParams Token is Invalid, or not present...")
	//LOG.Println("Check the Token in Cookie...")
	whoami = GetTokenCookie(c)
	//LOG.Println("Checking if Token is valid, and also getting User...")
	isTokenValid, user = whoami.ValidateTokenAndGetUser()
	return isTokenValid, whoami, user
}

// ----- VIEW FUNCTIONS -----

// GET: Form to validate the token |
// POST: Validate the token, and redirect to /card
func WhoAmIView(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: WhoAmIView")
	if c.Method() == "GET" {
		//LOG.Println("Render and return the WhoAmI Page for User...")
		return c.Render("whoami", fiber.Map{
			"Title": "RSVP",
		}, "base")

	} else if c.Method() == "POST" {
		whoami := new(WhoAmI)
		//LOG.Println("Parsing the request body for Token...")
		if err := c.BodyParser(whoami); err != nil {
			//LOG.Println("Failed to parse the request body for Token...")
			return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
		}
		//LOG.Println("Validating Token...")
		isTokenValid, _ := whoami.ValidateTokenAndGetUser()
		if isTokenValid {
			//LOG.Println("Token is valid, setting it into Cookie...")
			SetTokenCookie(c, whoami.Token)
			//LOG.Println("Redirect to CARD Page...")
			return c.Redirect("/card")

		} else {
			//LOG.Println("Token INVALID, Redirect back to /whoami Page")
			return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
		}
	}
	log.Fatal("Unreachable Flow, DEBUG Immediately...")
	return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
}


// GET: Show the card upon successful validation
func CardView(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: CardView")
	//LOG.Println("Extract Token from QueryParams or Cookie...")
	isTokenValid, whoami, user := extractTokenFromQueryOrCookieAndValidate(c)
	if isTokenValid {
		//LOG.Println("Token is valid, setting it in Cookie...")
		SetTokenCookie(c, whoami.Token)

		//LOG.Println("Generating a shareable URL for User...")
		cardUrl, _ := c.GetRouteURL("card", fiber.Map{})
		cardUrl = c.BaseURL() + cardUrl + "/?t=" + user.Token[:32]
		imageUrl := c.BaseURL() + "/static/img/og-logo.jpg"

		//LOG.Println("Render and return CARD page with all data...")
		return c.Render("card", fiber.Map{
			"Title": "RSVP CARD",
			"User": user,
			"CardUrl": cardUrl,
			"ImageUrl": imageUrl,
		}, "card-base")

	} else {
		//LOG.Println("Token INVALID, Redirect back to /whoami Page")
		return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
	}
}

// GET: RSVP action for the user
// | return the HTML template with Card
func RsvpView(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: RsvpView")
	//LOG.Println("Extract Token from QueryParams or Cookie...")
	isTokenValid, whoami, user := extractTokenFromQueryOrCookieAndValidate(c)
	if isTokenValid {
		//LOG.Println("Token is valid, setting it in Cookie...")
		SetTokenCookie(c, whoami.Token)

		//LOG.Println("Parsing the RSVP Status from Request Body...")
		rsvp := new(Rsvp)
		if err := c.BodyParser(rsvp); err != nil {
			//LOG.Println("Parsing failed, return StatusBadRequest, and redirect back to Card")
			return c.Status(fiber.StatusBadRequest).Redirect("/card")
		}
		//LOG.Println("Parsing successful, setting the User.Rsvp...")
		user.Rsvp = rsvp.Rsvp
		result := DB.Save(&user)
		if result.Error != nil {
			//LOG.Println("Failed to set the RSVP status in DB...")
			return c.Status(fiber.StatusBadRequest).Redirect("/card")
		}
		//LOG.Println("RSVP Status of User changed...")
		return c.Redirect("/card")
	} else {
		//LOG.Println("Token Invalid, return StatusUnauthorized")
		return c.Status(fiber.StatusUnauthorized).Redirect("/whoami")
	}
}


// ALL: Redirect User to GET: /whoami page
// | a catch-all view to catch all misc URLs and redirect
// instead of throwing errors
func RedirectToWhoAmI(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: RedirectToWhoAmI")
	//LOG.Println("Redirecting and undefined Admin URL to WhoAmI View...")
	return c.Status(fiber.StatusNotFound).Redirect("/whoami")
}
