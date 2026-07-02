package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/AnshuKashyap01/URL_Shortner/database"
	"github.com/AnshuKashyap01/URL_Shortner/helpers"
	"github.com/asaskevich/govalidator/v12"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	QRcode          string        `json:"qr"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c fiber.Ctx) error {

	body := new(request) //Defining new empty request  datatype

	err := c.Bind().Body(body) //Json se data leke body mein bhar raha hai

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	// 	User request
	//       |
	//       v
	// Get IP
	//       |
	//       v
	// Redis mein IP search karo
	//       |
	//       +---- Nahi mili?
	//       |         |
	//       |         v
	//       |   New quota set karo
	//       |
	//       +---- Mili?
	//                 |
	//                 v
	//         Remaining quota check karo
	//                 |
	//          +------+------+
	//          |             |
	//      quota > 0     quota <= 0
	//          |             |
	//          v             v
	//     Request allow   Error return

	//implement rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.IP()).Result() //Current IP ke liye quota check karna

	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err() //Har naye user ko 10 requests milengi jo 30 minutes ke liye valid hain.
	} else {
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result() //TTL ->Time To Live  time batao kitna bacha hai even if Quouta khatam hogaya hai.
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate Limit exceeded",
				"rate_limit_reset": limit / time.Minute, //limit time ke bad  try karna limit khatam ho gayi hai.
			})
		}
	}

	//check if input is an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	//check for domain error

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Nice Try Diddy (`-`)"})
	}

	//enforce https,SSL

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL custom short is already in use",
		})

	}
	if body.Expiry == 0 {
		body.Expiry = 24

	}
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	username := c.Locals("username")
	err = r.HSet(
		database.Ctx,
		"user_urls:"+username.(string),
		id,
		body.URL,
	).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to save user URL",
		})
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		QRcode:          "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	r2.Decr(database.Ctx, c.IP())

	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = os.Getenv("DOMAIN") + "/go/" + id
	resp.QRcode = os.Getenv("DOMAIN") + "/qr/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
