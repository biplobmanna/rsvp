package rsvp

import (
	"github.com/gofiber/fiber/v2"
)

func AddUrls(app *fiber.App) {
	// ****************************************************************
	// URLs for non-admin users
	// ****************************************************************
	// Validate normal User
	app.Get("/whoami", WhoAmIView).Name("whoami")
	app.Post("/whoami", WhoAmIView).Name("whoami-check")

	// Card view for the validated users
	app.Get("/card", CardView).Name("card")

	// ****************************************************************
	// URLs related to ADMIN
	// ****************************************************************
	// Group all Admin URLs
	admin := app.Group("/admin").Name("admin-")
	admin.Get("/users", AdminViewUsers).Name("users-get-all")

	// Users CRUD Operations
	admin.Post("/users", AdminViewUserCrud).Name("users-add")
	admin.Get("/users/:id", AdminViewUserCrud).Name("users-get")
	//? since default form only supports GET/POST, using instead of PATCH
	admin.Post("/users/:id", AdminViewUserCrud).Name("users-post")
	admin.Delete("/users/:id", AdminViewUserCrud).Name("users-delete")

	// Validate Admin
	admin.Get("/whoami", AdminCheckWhoAmI).Name("admin-whoami")
	admin.Post("/whoami", AdminCheckWhoAmI).Name("admin-whoami-check")

	// all other admin urls will be redirected to /admin/whoami
	admin.All("/*", RedirectToAdmin).Name("admin-redirect")

	// ***************************************************************
	// MISC URLs
	// ***************************************************************
	// all other urls will be redirected to /whoami
	app.All("/*", RedirectToWhoAmI).Name("whoami-redirect")
}
