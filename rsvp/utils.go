package rsvp

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TEMPORARY TOKEN, to be later transferred to .env
var DEMO_TOKEN = "DEMO"

type WhoAmI struct {
	Token string `json:"token" xml:"token" form:"token" cookie:"token"`
}

func (w *WhoAmI) ValidateToken() bool {
	// Implement token validations
	return DEMO_TOKEN == w.Token
}

func SetTokenCookie(c *fiber.Ctx, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = "DEMO"
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Secure = true

	c.Cookie(cookie)
}

func GetTokenCookie(c *fiber.Ctx) (WhoAmI, error) {
	whoami := new(WhoAmI)
	if err := c.CookieParser(whoami); err != nil {
		return *whoami, err
	}
	fmt.Println(whoami)

	return *whoami, nil
}
