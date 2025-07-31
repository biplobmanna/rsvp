package rsvp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// --- HELPER FUNCTIONS ---

func extractTokenCookieAndValidateAdmin(c *fiber.Ctx) (bool, WhoAmI) {
	fmt.Println(c.Cookies("token"))
	whoami := GetTokenCookie(c)
	return whoami.ValidateAdminToken(), whoami
}

// ---- ADMIN VIEWS ---

// redirect any url caught here to /admin/whoami
func RedirectToAdmin(c *fiber.Ctx) error {
	// any URLs other than the ones specified redirects to .../whoami
	return c.Redirect("/admin/whoami")
}

// fetch the HTML template with the token submit form for GET
// , and check the "token" for the POST request
func AdminCheckWhoAmI(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		// for GET, render the HTML template for whoami
		return c.Render("whoami", fiber.Map{
			"Title":          "RSVP: Admin",
			"CheckWhoAmIUrl": "/admin/whoami",
		}, "base")
	} else if c.Method() == "POST" {
		// for POST, check the "token"
		// and if valid, then update token

		// declare a WhoAmI struct to hold the token data
		// fetched from the request
		whoami := new(WhoAmI)

		// parse request body for the data in WhoAmI
		if err := c.BodyParser(whoami); err != nil {
			// if error in parsing, redirect back to the same page
			// this should ensure that the input page is shown again
			return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
		}
		isTokenValid := whoami.ValidateAdminToken()
		if isTokenValid {
			// Set/Refresh cookie with the token
			SetTokenCookie(c, whoami.Token)
			return c.Redirect("/admin/users")
		}

		// return unauthorized if token is invalidated
		return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
	}

	// this part of the code should be unreachable
	// even if reached due to some sort of mishap,
	// return the StatusBadRequest
	return c.Status(fiber.StatusBadRequest).Redirect("/admin/whoami")
}

// return the HTML template with the complete of users
// | no pagination yet, can be added later as an enhancement
func AdminViewUsers(c *fiber.Ctx) error {
	isTokenValid, whoami := extractTokenCookieAndValidateAdmin(c)

	if isTokenValid {
		// refresh cookie
		SetTokenCookie(c, whoami.Token)

		// create a array slice of Users struct
		var results []User
		// populate teh array slice with the list of all users
		DB.Table("users").Find(&results)

		// render the HTML template with all the users and return
		return c.Render("users", fiber.Map{
			"Title": "RSVP: Admin - Users",
			"Users": results,
		}, "base")
	}

	// redirect to /admin/whoami for invalidated token
	return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
}

// CRUD operations for Admin
// | Naming is bad, to be fixed later
// | handles both new, and existing users for GET
func AdminViewUserCrud(c *fiber.Ctx) error {
	// check for permissions before all operations
	isTokenValid, whoami := extractTokenCookieAndValidateAdmin(c)
	if isTokenValid {
		// refresh cookie token
		SetTokenCookie(c, whoami.Token)
	} else {
		// return unauthorized
		return c.Status(fiber.StatusUnauthorized).Redirect("/admin/whoami")
	}

	// for a validated request, then perform the rest of the actions
	// handle the GET request, for both new and existing users
	if c.Method() == "GET" {
		// the param which accepts either
		// "new" or "<id>"
		param := c.Params("id")

		// User struct to store the new/existing user data
		var user User
		// flag to mark new user addition or existing user update
		// this flag is user in HTML template to populate or leave
		// the input fields empty
		var update bool

		// try to conver the <id> param into int
		id, err := strconv.Atoi(param)

		if param == "new" {
			// if the param is new, the request is for
			// adding a new user, which is empty
			user = User{}
			// mark flag for new user addition instead of update
			update = false
		} else if err == nil {
			// if err is nil, means that atoi() is successful
			// which means that <id> is an integer
			// which related to the PK in Users table in DB

			// fetch data from DB
			queryResult := DB.Table("users").First(&user, id)
			if queryResult.Error != nil {
				// if the <id> doesn't belong to any Users<PK>
				// the user does not exist, so send similar status
				return c.SendStatus(fiber.StatusNotFound)
			}
			// mark flag to denote user exists, and updation operation
			update = true
		} else {
			// this block is a wrong flow,
			// so StatusNotFound
			return c.SendStatus(fiber.StatusNotFound)
		}

		// render the HTML with the user details if present
		// , or with empty fields for new users
		return c.Render("user", fiber.Map{
			"Title":  "Rsvp: Admin - Edit User",
			"Update": update,
			"User":   user,
		}, "base")

	} else if c.Method() == "POST" {
		// handle the POST request which also captures both:
		// user creation and user update

		// new User struct to store the parsed User data from request
		user := User{}
		// parse the body for the User data
		if err := c.BodyParser(&user); err != nil {
			// if parsing fails for any reason, 
			// return StatusExpectationFailed
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		// trim the space from users comments
		user.Comments = strings.TrimSpace(user.Comments)

		// update the DB
		// if the user is a new user, the "id" is NIL,
		// and will be auto-generated when creating the row
		// for update, the "id" is the PK identifying a row
		result := DB.Save(&user)
		if result.Error != nil {
			// if somehow the DB update fails, then return StatusExpectationFailed
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		fmt.Printf("✅ User {ID: %d} updated...\n", user.ID)

		// once user is added, redirect to view the list of all users
		return c.Redirect("/admin/users")

	} else if c.Method() == "DELETE" {
		// handle the DELETE request

		// get the <id> from request
		id := c.Params("id")

		// will be used to point to the correct table in DB
		var user User
		// try to delete the user identified by the <id>
		queryResult := DB.Delete(&user, id)

		if queryResult.Error != nil {
			// if failure to delete, then return StatusExpectationFailed
			return c.SendStatus(fiber.StatusExpectationFailed)
		}
		// this part is reached for successful deletion of resource
		return c.SendStatus(fiber.StatusNoContent)
	}
	// this code should not be reachable, so returns a StatusBadRequest
	return c.SendStatus(fiber.StatusBadRequest)
}

// TODO: to return Admin HTML template in future,
// not yet implemented
func AdminView(c *fiber.Ctx) error {
	return c.SendString("Admin View")
}

// get the user shareable link for the card page
func UserShareLinkView(c *fiber.Ctx) error {
	isTokenValid, _ := extractTokenCookieAndValidateAdmin(c)

	if !isTokenValid {
		return c.SendStatus(fiber.StatusUnauthorized)
	} 

	// get the user from the <id>
	user := User{}

	// extract <id> params, 
	// defaults to -9999
	// since id is PK, this should not exist
	id := c.Params("id", "-9999")

	// get user details from DB
	result := DB.First(&user, id)

	// if no result, and/or error
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// if user is found, create a shareable URL for the user
	cardUrl, _ := c.GetRouteURL("card", fiber.Map{})
	url := c.BaseURL() + cardUrl + "/?t=" + user.Token[:32]
	fmt.Println(url)

	return c.SendString(url)
}
