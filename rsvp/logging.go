package rsvp

import (
	"log"
	"os"
)

type logger struct {
	*log.Logger
}

var LOG *logger

func SetupLogging() {
	file, _ := os.OpenFile("rsvp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

	LOG = &logger{
		Logger:   log.New(file, "", log.Lshortfile|log.LstdFlags),
	}
}
