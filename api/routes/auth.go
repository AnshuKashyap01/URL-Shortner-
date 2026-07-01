package routes

import (
	"fmt"
	"os"
	"time"

	models "github.com/AnshuKashyap01/URL_Shortner/Models"
	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c fiber.Ctx) error {
	user := new(models.User)

	err := c.Bind().Body(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	r := database.CreateClient(2)
	defer r.Close()

	exists, err := r.Exists(
		database.Ctx,
		"user:"+user.Username,
	).Result()

	if err != nil {

		fmt.Println("Redis Error:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Database error",
			},
		)
	}
	if exists == 1 {
		return c.Status(fiber.StatusConflict).JSON(
			fiber.Map{
				"error": "User already exists",
			},
		)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	// If you want a more production-like design → use HSet().
	//HSet in Redis is used to store multiple fields under a single key. Think of it as storing an object or struct.

	err = r.HSet(database.Ctx,
		"user:"+user.Username,
		"email", user.Email,
		"password", string(hashedPassword),
	).Err()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not save user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})

}

func Login(c fiber.Ctx) error {

	user := new(models.User)

	if err := c.Bind().Body(user); err != nil { //Json to struct conversion
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	r := database.CreateClient(2)
	defer r.Close()

	hashedPassword, err := r.HGet(database.Ctx,
		"user:"+user.Username,
		"password").Result()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User Not Found",
		})
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	tokenString, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})

	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})

}
