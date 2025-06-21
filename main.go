package main

import (
	"github.com/biplobmanna/rsvp/rsvp"
)

func main() {
	// comment/un-comment this line as needed
	// this is to run migrationso
	// db := rsvp.Connect()
	// rsvp.MigrateAll(db)

	app := rsvp.App()

	app.Listen(":3000")
}
