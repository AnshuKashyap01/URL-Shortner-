package routes

import (
	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
)

func DeleteURL(c fiber.Ctx) error {

	id := c.Params("id")
	username := c.Locals("username").(string)

	r := database.CreateClient(0)
	

	_, err := r.HGet(
		database.Ctx,
		"user_urls:"+username,
		id,
	).Result()

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"error": "You do not own this URL",
			},
		)
	}

	r.Del(database.Ctx, id)

	r.Del(database.Ctx, "clicks:"+id)

	r.HDel(
		database.Ctx,
		"user_urls:"+username,
		id,
	)

	return c.JSON(fiber.Map{
		"message": "URL deleted successfully",
	})
}
 