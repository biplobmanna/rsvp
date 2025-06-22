package rsvp

import (
	"github.com/gofiber/fiber/v2"
)

func AddUrls(app *fiber.App) {
	// Add the URLs
	app.Get("/", IndexView).Name("index")
	app.Get("/whoami", WhoAmIView).Name("whoami")
	app.Get("/card", CardView).Name("card")

	app.Post("/check-whoami", CheckWhoAmI).Name("check-whoami")

	// Admin View
	admin := app.Group("/admin").Name("admin-")
	admin.Get("/users", AdminViewUsers).Name("all-users-get")
	admin.Post("/users", AdminViewUserCrud).Name("users-add")
	admin.Get("/users/:id", AdminViewUserCrud).Name("users-get")
	admin.Post("/users/:id", AdminViewUserCrud).Name("users-patch")
	admin.Delete("/users/:id", AdminViewUserCrud).Name("users-delete")
	admin.Get("/whoami", AdminCheckWhoAmI).Name("whoami-get")
	admin.Post("/whoami", AdminCheckWhoAmI).Name("whoami-post")

	// all other admin urls will be redirected to /admin/whoami
	admin.All("/*", RedirectToAdmin).Name("redirect")
	// all other urls will be redirected to /whoami
	app.All("/*", RedirectToIndex).Name("index-redirect")

	// Print All URLS
	// stack, _ := json.MarshalIndent(app.Stack(), "", " ")
	// fmt.Println(string(stack))
}
