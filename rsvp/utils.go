package rsvp

import (
	"time"
	"github.com/gofiber/fiber/v2"
)

// struct to hold the token
type WhoAmI struct {
	Token string `json:"token" xml:"token" form:"token" cookie:"token"`
}

// struct to hold the Admin token
type AdminWhoAmI struct {
	SuperToken string `json:"supertoken" xml:"supertoken" form:"supertoken" cookie:"supertoken"`
}

// struct to hold the rsvp details
type Rsvp struct {
	Rsvp bool `json:"rsvp" xml:"rsvp" form:"rsvp" cookie:"rsvp"`
}

// method: vaildate token
func (w *WhoAmI) ValidateTokenAndGetUser() (bool, User) {
	user := User{}

	tokenLen := len(w.Token)
	if tokenLen == 16 || tokenLen == 32 {
		// do nothing
	} else {
		//LOG.Println("Token is of incorrent length:", tokenLen, "...")
		return false, user
	}

	//LOG.Println("Fetching User details using the Token...")
	token := w.Token + "%"
	result := DB.Where("token like ?", token).First(&user)

	if result.Error != nil {
		//LOG.Println("No Users found with the Token, i.e., Token is Invalid...")
		return false, user
	}

	//LOG.Println("User Found, Token Valid...")
	return true, user
}

// method: validate admin token
func (w *AdminWhoAmI) ValidateAdminToken() bool {
	//LOG.Println("Validating Admin Token...")
	return SETTINGS.ADMIN_TOKEN == w.SuperToken
}

// set token to cookie
func SetTokenCookie(c *fiber.Ctx, key, value string) {
	//LOG.Println("Setting the Token into Cookie...")
	cookie := new(fiber.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Secure = true

	c.Cookie(cookie)
}

// get token from cookie
func GetTokenCookie(c *fiber.Ctx) WhoAmI {
	//LOG.Println("Extracting Token from Cookie...")
	whoami := new(WhoAmI)
	c.CookieParser(whoami) // ignoring error handling
	return *whoami
}

// get admin token from cookie
func AdminGetTokenCookie(c *fiber.Ctx) AdminWhoAmI {
	//LOG.Println("Extracting Token from Cookie...")
	whoami := new(AdminWhoAmI)
	c.CookieParser(whoami) // ignoring error handling
	return *whoami
}

// get token from query params
func GetTokenQuery(c *fiber.Ctx) WhoAmI {
	//LOG.Println("Extracting Token from Query Params...")
	whoami := new(WhoAmI) 
	whoami.Token = c.Query("t", "")
	return *whoami
}

