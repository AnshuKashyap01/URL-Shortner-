package main

import (
	"log"
	"os"

	middleware "github.com/AnshuKashyap01/URL_Shortner/middleWare"
	"github.com/AnshuKashyap01/URL_Shortner/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/v1", middleware.Protected(), routes.ShortenURL)
	app.Get("/analytics/:url", routes.GetAnalytics)
	app.Get("/myurls", middleware.Protected(), routes.GetMyURLs)
	app.Post("/signup", routes.Signup)
	app.Post("/login", routes.Login)
	app.Delete("/url/:id", middleware.Protected(), routes.DeleteURL)

	app.Get("/qr/:url", routes.GenerateQr)

	// Always keep this at the end
	app.Get("/go/:url", routes.ResolveURL)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	app := fiber.New()

	app.Use(logger.New()) //Fiber mein ek middleware add karta hai jo har HTTP request
	// (GET, POST, PUT, DELETE, etc.) ki details console mein log karta hai,
	// jisse debugging aur monitoring easy ho jaati hai.

	app.Use("/static", static.New("./static"))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("./static/index.html")
	})

	app.Get("/login", func(c fiber.Ctx) error {

		return c.SendFile("./static/login.html")

	})

	app.Get("/signup", func(c fiber.Ctx) error {

		return c.SendFile("./static/signup.html")

	})

	setupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
