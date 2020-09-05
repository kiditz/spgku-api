package middleware

import (
	"crypto/subtle"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

// IsLoggedIn used to handle token verification
func IsLoggedIn() echo.MiddlewareFunc {
	return m.JWTWithConfig(m.JWTConfig{
		SigningKey: []byte(os.Getenv("ACCESS_SECRET")),
	})
}

// IsMidtrans used to handle token verification
func IsMidtrans() echo.MiddlewareFunc {
	return m.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(os.Getenv("MIDTRANS_SERVER_KEY"))) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("")) == 1 {
			return true, nil
		}
		return false, nil
	})
}

var (
	company = "company"
	talent  = "talent"
)

// IsCompany used to routing if loggined user is company
func IsCompany(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isCompany := claims["type"] == company
		if !isCompany {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// IsTalent used to routing if loggined user is talent
func IsTalent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isTalent := claims["type"] == talent
		fmt.Printf("Is Talent %v", isTalent)
		if !isTalent {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
