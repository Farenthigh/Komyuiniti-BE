package utils

import (
	"github.com/Farenthigh/Fitbuddy-BE/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsExist(c *fiber.Ctx) error {
	auth := c.Cookies("auth")
	if auth == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	token, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt_secret), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	claim := token.Claims.(jwt.MapClaims)
	if claim["exp"] == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	c.Locals("userID", claim["id"])
	c.Locals("email", claim["email"])
	c.Locals("username", claim["username"])
	return c.Next()
}
