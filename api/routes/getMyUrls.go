package routes

import (
	"strconv"

	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type URLInfo struct {
	Short    string `json:"short"`
	Original string `json:"original"`
	Clicks   int    `json:"clicks"`
}

func GetMyURLs(c fiber.Ctx) error {

	username := c.Locals("username")

	r := database.CreateClient(0)
	defer r.Close()

	urls, err := r.HGetAll(
		database.Ctx,
		"user_urls:"+username.(string),
	).Result()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Cannot fetch URLs",
		})
	}

	var result []URLInfo

	for short, original := range urls {

		clicks, err := r.Get(database.Ctx, "clicks:"+short).Result()

		if err == redis.Nil {
			clicks = "0"
		}

		totalClicks, _ := strconv.Atoi(clicks)

		result = append(result, URLInfo{
			Short:    short,
			Original: original,
			Clicks:   totalClicks,
		})
	}

	return c.JSON(fiber.Map{
		"user": username,
		"urls": result,
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
