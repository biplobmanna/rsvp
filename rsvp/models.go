package rsvp

import (
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// any model struct that you create, add to this interface
type Model interface {
	User
}

type User struct {
	gorm.Model
	FullName string
	Token    string
	Email    string
	Phone    string
	Rsvp     bool
	Comments string
}
