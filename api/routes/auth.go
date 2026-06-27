package routes

import (
	models "github.com/AnshuKashyap01/URL_Shortner/Models"
	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
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

	_, err2 := r.Get(database.Ctx, "user:"+user.Username).Result() //to check whether user already exist or not

	if err2 == nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "USer already exists",
		})
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
			"email",user.Email,
			"password",string(hashedPassword),
		).Err()

	
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error":"Could not save user",
		})
	}


	return c.Status(201).JSON(fiber.Map{
		"message":"User created successfully",
	})

}
