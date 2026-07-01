package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {

	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		tokenString := strings.TrimPrefix(  //Removed bearer from the string we only token
			authHeader, "Bearer ",
		)

		token, err := jwt.Parse(tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid Token",
			})
		}

		 claims := token.Claims.(jwt.MapClaims)//JWT token ke andar jo data store tha use nikal rahe ho

		c.Locals(
			"username",
			claims["username"],
		)

		return c.Next() //Mera kaam khatam ab next handler chalao

	}
}
