package rsvp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// redirect any url caught here to /admin/whoami
func RedirectToAdmin(c *fiber.Ctx) error {
	// any URLs other than the ones specified redirects to .../whoami
	return c.Redirect("/admin/whoami")
}

// fetch the HTML template with the token submit form for GET
// , and check the "token" for the POST request
func AdminCheckWhoAmI(c *fiber.Ctx) error {
	// for GET, render the HTML template for whoami
	if c.Method() == "GET" {
		return c.Render("whoami", fiber.Map{
			"Title":          "RSVP: Admin",
			"CheckWhoAmIUrl": "/admin/whoami",
		}, "base")
	} else if c.Method() == "POST" {
		// for POST, check the "token"
		// and if valid, then update token

		// extract token from cookie
		whoami := GetTokenCookie(c)

		// if contains valid token, return "Token Validated"
		// else, return error
		if whoami.ValidateToken() {
			return c.SendString("Token Validated")
		}

	}
	// this part of the code should be unreachable
	// even if reached due to some sort of mishap,
	// return the StatusBadRequest
	return c.SendStatus(fiber.StatusBadRequest)
}

// return the HTML template with the complete of users
// | no pagination yet, can be added later as an enhancement
func AdminViewUsers(c *fiber.Ctx) error {
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

// CRUD operations for Admin
// | Naming is bad, to be fixed later
// | handles both new, and existing users for GET
func AdminViewUserCrud(c *fiber.Ctx) error {
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
