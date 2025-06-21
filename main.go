package main

import (
	"github.com/biplobmanna/rsvp/rsvp"
)

func main() {
	app := rsvp.App()

	app.Listen(":3000")
}
