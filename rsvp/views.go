package rsvp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// The views are used to serve HTML pages
func IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Base HTML",
		"Contents": "Hello, World!",
	}, "base")
}

func WhoAmIView(c *fiber.Ctx) error {
	return c.Render("whoami", fiber.Map{
		"Title":          "RSVP",
		"CheckWhoAmIUrl": "/check-whoami",
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
	whoami, err := GetTokenCookie(c)
	if err != nil {
		return err
	}

	if whoami.ValidateToken() {
		SetTokenCookie(c, whoami.Token)
		return c.Render("card", fiber.Map{
			"Title": "RSVP CARD",
		}, "base")
	} else {
		return c.Redirect("/whoami")
	}
}

func RedirectToIndex(c *fiber.Ctx) error {
	return c.Redirect("/whoami")
}

func RedirectToAdmin(c *fiber.Ctx) error {
	return c.Redirect("/admin/whoami")
}

func AdminCheckWhoAmI(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("whoami", fiber.Map{
			"Title":          "RSVP: Admin",
			"CheckWhoAmIUrl": "/admin/whoami",
		}, "base")
	} else if c.Method() == "POST" {
		whoami, err := GetTokenCookie(c)
		if err != nil {
			return c.SendStatus(fiber.StatusExpectationFailed)
		}
		if whoami.ValidateToken() {
			return c.SendString("Token Validated")
		}

	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func AdminViewUsers(c *fiber.Ctx) error {
	var results []User
	DB.Table("users").Find(&results)

	return c.Render("users", fiber.Map{
		"Title": "RSVP: Admin - Users",
		"Users": results,
	}, "base")
}

func AdminViewUserCrud(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		param := c.Params("id")
		var user User
		var update bool
		id, err := strconv.Atoi(param)

		if param == "new" {
			user = User{}
			update = false
		} else if err == nil {
			// fetch data from DB
			queryResult := DB.Table("users").First(&user, id)
			if queryResult.Error != nil {
				return c.SendStatus(fiber.StatusNotFound)
			}
			update = true
		} else {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Render("user", fiber.Map{
			"Title":  "Rsvp: Admin - Edit User",
			"Update": update,
			"User":   user,
		}, "base")
	} else if c.Method() == "POST" {
		user := User{}
		if err := c.BodyParser(&user); err != nil {
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		user.Comments = strings.TrimSpace(user.Comments)

		// update the DB
		result := DB.Save(&user)
		if result.Error != nil {
			c.SendStatus(fiber.StatusExpectationFailed)
		}
		fmt.Printf("✅ User {ID: %d} updated...\n", user.ID)
		return c.Redirect("/admin/users")
	} else if c.Method() == "DELETE" {
		id := c.Params("id")
		var user User
		queryResult := DB.Delete(&user, id)
		if queryResult.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func AdminView(c *fiber.Ctx) error {
	return c.SendString("Admin View")
}
