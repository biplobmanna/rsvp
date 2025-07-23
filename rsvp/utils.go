package rsvp

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// TEMPORARY TOKEN, to be later transferred to .env
var DEMO_TOKEN = "DEMO"

// struct to hold the token
type WhoAmI struct {
	Token string `json:"token" xml:"token" form:"token" cookie:"token"`
}

// method: vaildate token
func (w *WhoAmI) ValidateToken() bool {
	// token must not be empty
	if w.Token == "" {
		return false
	}

	// token must belong to exactly one user
	_, err := GetUserFromToken(w.Token)
	if err != nil {
		return false
	}
	return true
}

// set token to cookie
func SetTokenCookie(c *fiber.Ctx, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Secure = true

	c.Cookie(cookie)
}

// get token from cookie
func GetTokenCookie(c *fiber.Ctx) WhoAmI {
	whoami := new(WhoAmI)
	c.CookieParser(whoami) // ignoring error handling
	return *whoami
}

// get token from query params
func GetTokenQuery(c *fiber.Ctx) WhoAmI {
	whoami := new(WhoAmI) 
	whoami.Token = c.Query("t", "")
	return *whoami
}

// get non-admin user from token
func GetUserFromToken(token string) (User, error) {
	user := User{}
	result := DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
