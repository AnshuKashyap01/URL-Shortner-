package routes

import (
	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)

	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in database"})

	} else if err != nil { 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"cannot connect to DB"})
	}

	rInr := database.CreateClient(0)
	defer rInr.Close()

	_= rInr.Incr(database.Ctx,"clicks:"+url)


	return c.Redirect().Status(301).To(value)

}
