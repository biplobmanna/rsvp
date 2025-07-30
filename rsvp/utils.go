package rsvp

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// struct to hold the token
type WhoAmI struct {
	Token string `json:"token" xml:"token" form:"token" cookie:"token"`
}

// struct to hold the rsvp details
type Rsvp struct {
	Rsvp bool `json:"rsvp" xml:"rsvp" form:"rsvp" cookie:"rsvp"`
}

// method: vaildate token
func (w *WhoAmI) ValidateTokenAndGetUser() (bool, User) {
	// user data to return
	user := User{}

	// token must not be empty
	tokenLen := len(w.Token)
	if tokenLen == 16 || tokenLen == 32 {
		// do nothing
	} else {
		return false, user
	}

	// partial find using "token"
	token := w.Token + "%"
	result := DB.Where("token like ?", token).First(&user)

	if result.Error != nil {
		return false, user
	}

	// return true and user
	return true, user
}

// method: validate admin token
func (w *WhoAmI) ValidateAdminToken() bool {
	return SETTINGS.ADMIN_TOKEN == w.Token
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

