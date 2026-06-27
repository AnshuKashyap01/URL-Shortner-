package routes

import (
	"strconv"

	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func GetAnalytics(c fiber.Ctx) error{
	shortUrl:=c.Params("url")

	r:= database.CreateClient(1)
	defer r.Close()

	clicks,err := r.Get(database.Ctx,"clicks"+shortUrl).Result()

	if err==redis.Nil{
		clicks = "0"
	}else if err!= nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"Database error",
		})
	}

	totalClicks,_:= strconv.Atoi(clicks)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"short_url": shortUrl,
		"clicks": totalClicks,
	})



}