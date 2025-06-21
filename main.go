package main

import (
	"github.com/biplobmanna/rsvp/rsvp"
)

func main() {
	// comment/un-comment this line as needed

	app := rsvp.App()

	app.Listen(":3000")
}
