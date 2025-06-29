package rsvp

import (
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// any model struct that you create, add to this interface
type Model interface {
	Admin | User
}

type Admin struct {
	gorm.Model
	User  string
	Token string
}

type User struct {
	gorm.Model
	User     string
	FullName string
	Token    string
	Email    string
	Phone    string
	Rsvp     bool
	Comments string
}
