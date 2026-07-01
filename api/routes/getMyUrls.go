package routes

import (
	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
)

func GetMyURLs(c fiber.Ctx) error {
	username := c.Locals("username")

	r:= database.CreateClient(0)
	defer r.Close()


	urls ,err := r.HGetAll(
		database.Ctx,
		"user_urls:"+username.(string),
	).Result()

	if err!= nil{
		return c.Status(500).JSON(fiber.Map{
			"error":"Cannot fetch URLs",
		})
	}


	return c.JSON(fiber.Map{
		"user": username,
		"urls": urls,
	})

}

// JWT Token
//     |
//     v
// Middleware extracts username
//     |
//     v
// c.Locals("username") = "anshu"
//     |
//     v
// "user_urls:anshu"
//     |
//     v
// Redis HGetAll
//     |
//     v
// Returns all URLs created by Anshu