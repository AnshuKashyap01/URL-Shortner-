package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AnshuKashyap01/URL_Shortner/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/analytics/:url",routes.GetAnalytics)
	app.Post("/signup",routes.Signup)
	
}

func main() {

	err:=godotenv.Load() //godotenv.Load() ka kaam .env file ko load karna hai taaki os.Getenv() se uske variables access kar sako.

	if err!= nil{
		fmt.Println(err)
	}
  
	app := fiber.New()
	
	app.Use(logger.New())//Fiber mein ek middleware add karta hai jo har HTTP request 
	// (GET, POST, PUT, DELETE, etc.) ki details console mein log karta hai, 
	// jisse debugging aur monitoring easy ho jaati hai.

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))



}