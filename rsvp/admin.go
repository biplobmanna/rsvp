package rsvp

import (
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v2"
)

// --- HELPER FUNCTIONS ---

func extractTokenCookieAndValidateAdmin(c *fiber.Ctx) (bool, AdminWhoAmI) {
	//LOG.Println("extract token from cookie and validate admin...")
	whoami := AdminGetTokenCookie(c)
	//LOG.Println("TOKEN:", whoami.SuperToken)
	return whoami.ValidateAdminToken(), whoami
}

// ---- ADMIN VIEWS ---

// redirect any url caught here to /admin/whoami
func RedirectToAdmin(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: RedirectToAdmin")
	//LOG.Println("Redirecting URL:", c.OriginalURL(), "to /admin/whoami...")
	return c.Redirect("/admin/whoami")
}

// fetch the HTML template with the token submit form for GET
// , and check the "token" for the POST request
func AdminCheckWhoAmI(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: AdminCheckWhoAmI")
	if c.Method() == "GET" {
		//LOG.Println("Render the template for WhoAmI for Admin...")
		return c.Render("whoami", fiber.Map{
			"Title":          "RSVP: Admin",
			"CheckWhoAmIUrl": "/admin/whoami",
		}, "base")
	} else if c.Method() == "POST" {
		whoami := new(AdminWhoAmI)
		// since "token" is passed in POST request
		// must extract that from body first
		whoamiToken := new(WhoAmI)

		//LOG.Println("Parsing the request body for Token...")
		if err := c.BodyParser(whoamiToken); err != nil {
			//LOG.Println("Failed to parse the request body for Token...")
			return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
		}
		//LOG.Println("Setting the token parsed from Body to SuperToken...")
		whoami.SuperToken = whoamiToken.Token
		//LOG.Println("Check if the SuperToken is valid...")
		isTokenValid := whoami.ValidateAdminToken()
		if isTokenValid {
			//LOG.Println("The token is valid, set the token in cookie...")
			SetTokenCookie(c, "supertoken", whoami.SuperToken)
			//LOG.Println("Redirecting to /admin/users with the cookie...")
			return c.Redirect("/admin/users")
		}

		//LOG.Println("Token is INVALID!")
		//LOG.Println("Redirecting to /admin/users with Unauthorized status...")
		return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
	}

	//LOG.Panicln("Unreachable Path, check immediately...")
	return c.Status(fiber.StatusBadRequest).Redirect("/admin/whoami")
}

// return the HTML template with the complete of users
// | no pagination yet, can be added later as an enhancement
func AdminViewUsers(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: AdminViewUsers")
	//LOG.Println("Extract the Token from Cookie, and validate the Admin correctly...")
	isTokenValid, whoami := extractTokenCookieAndValidateAdmin(c)

	if isTokenValid {
		//LOG.Println("The Token is valid, setting it back to the Cookie...")
		SetTokenCookie(c, "supertoken", whoami.SuperToken)

		var results []User
		result := DB.Table("users").Find(&results)
		//LOG.Println("Fetching all Users from the DB...")

		if result.Error != nil {
			//LOG.Println("Failed to Fetch Users from DB, Error:", result.Error)
			//LOG.Println("The list of results are empty...")
		}

		//LOG.Println("Rendering the Admin-Users page with a list of all Users...")
		return c.Render("users", fiber.Map{
			"Title": "RSVP: Admin - Users",
			"Users": results,
		}, "base")
	}

	//LOG.Println("Token is INVALID!")
	//LOG.Println("Redirecting back to /admin/whoami with Status: Unauthorized...")
	return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
}

// CRUD operations for Admin
// | Naming is bad, to be fixed later
// | handles both new, and existing users for GET
func AdminViewUserCrud(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: AdminViewUserCrud")
	//LOG.Println("Extract the Token from Cookie, and validate the Admin correctly...")
	isTokenValid, whoami := extractTokenCookieAndValidateAdmin(c)
	if isTokenValid {
		//LOG.Println("The Token is valid, setting it back to the Cookie...")
		SetTokenCookie(c, "supertoken", whoami.SuperToken)
	} else {
		//LOG.Println("Token is INVALID!")
		//LOG.Println("Redirecting back to /admin/whoami with Status: Unauthorized...")
		return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
	}

	if c.Method() == "GET" {
		//LOG.Println("Extracting Param:<id> from the Request:URL...")
		param := c.Params("id")

		var user User
		var update bool

		//LOG.Println("Coverting Param:<id> from string to int...")
		id, err := strconv.Atoi(param)

		if param == "new" {
			//LOG.Println("Param=new - this relates to adding of a new user, unmark 'update' flag...")
			user = User{}
			update = false
		} else if err == nil {
			//LOG.Println("Param:<id> is a valid integer, and not 'new'...")
			//LOG.Println("Finding the User with <id> from the DB...")
			queryResult := DB.Table("users").First(&user, id)
			if queryResult.Error != nil {
				//LOG.Println("User with id=", id, "not found in DB...")
				//LOG.Println("Sending back a Status: NotFound")
				return c.SendStatus(fiber.StatusNotFound)
			}
			//LOG.Println("User details found in DB, mark 'update' flag...")
			update = true
		} else {
			//LOG.Println("Wrong URL, return Status: NotFound")
			return c.SendStatus(fiber.StatusNotFound)
		}

		//LOG.Println("Render the User Page for adding/updating user...")
		return c.Render("user", fiber.Map{
			"Title":  "Rsvp: Admin - Edit User",
			"Update": update,
			"User":   user,
		}, "base")

	} else if c.Method() == "POST" {
		//LOG.Println("Parsing the body for User data...")
		user := User{}
		if err := c.BodyParser(&user); err != nil {
			//LOG.Println("Failed to parse body for User data...")
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		//LOG.Println("User data parsed successfully...")
		//LOG.Println("Trimming all space in Comments...")
		user.Comments = strings.TrimSpace(user.Comments)
		//LOG.Println("Generating Token for User...")
		token, err := EncryptAES(user.FullName+SETTINGS.ADMIN_TOKEN)
		if err != nil {
			//LOG.Println("Failed to generate Token for User...")
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		user.Token = token
		//LOG.Println("Token geneated successfully...")

		//LOG.Println("Saving User data to DB...")
		result := DB.Save(&user)
		if result.Error != nil {
			//LOG.Println("Failed to save User data to DB...")
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		//LOG.Printf("✅ User {ID: %d} updated...\n", user.ID)

		//LOG.Println("Redirect back to /admin/users after successfully saving User...")
		return c.Redirect("/admin/users")

	} else if c.Method() == "DELETE" {
		//LOG.Println("Extracting Param:<id> from the Request:URL...")
		id := c.Params("id")

		var user User
		queryResult := DB.Delete(&user, id)
		//LOG.Println("Deleting User from the DB...")

		if queryResult.Error != nil {
			//LOG.Println("Failed to delete User from DB, return Status: ExpectationFailed")
			return c.SendStatus(fiber.StatusExpectationFailed)
		}
		//LOG.Println("Deleted Successfully...")
		return c.SendStatus(fiber.StatusNoContent)
	}
	//LOG.Panicln("Unreachable Block, Debug Immediately...")
	return c.SendStatus(fiber.StatusBadRequest)
}

// get the user shareable link for the card page
func UserShareLinkView(c *fiber.Ctx) error {
	//LOG.Println("[", c.Method(), "]","URL:", c.OriginalURL(), "| View: UserShareLinkView")
	//LOG.Println("Extract the Token from Cookie, and validate the Admin correctly...")
	isTokenValid, _ := extractTokenCookieAndValidateAdmin(c)

	if !isTokenValid {
		//LOG.Println("Token is INVALID, return Status: Unauthorized")
		return c.SendStatus(fiber.StatusUnauthorized)
	} 

	user := User{}
	id := c.Params("id", "-9999")
	//LOG.Println("Extracting Param:<id> from the Request:URL...")

	//LOG.Println("Fetching User details from DB...")
	result := DB.First(&user, id)

	if result.Error != nil {
		//LOG.Println("Failed to fetch user details from DB...")
		return c.SendStatus(fiber.StatusNotFound)
	}

	//LOG.Println("Generating a Shareable URL...")
	cardUrl, _ := c.GetRouteURL("card", fiber.Map{})
	url := c.BaseURL() + cardUrl + "/?t=" + user.Token[:32]

	//LOG.Println("Shareable URL:", url)
	return c.SendString(url)
}
