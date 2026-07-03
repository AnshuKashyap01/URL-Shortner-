package routes

import (
	"os"

	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"github.com/skip2/go-qrcode"
)

func GenerateQr(c fiber.Ctx) error {

	shortUrl := c.Params("url")

	r := database.CreateClient(0)
	

	_, err := r.Get(database.Ctx, shortUrl).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short url not found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Databse error",
		})
	}

	//To generate QR code--->

	png, err := qrcode.Encode(
		os.Getenv("DOMAIN")+"/go/"+shortUrl,
		qrcode.Medium,
		256,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate QR code",
		})
	}

	c.Set("Content-Type", "image/png")

	return c.Send(png)

}
