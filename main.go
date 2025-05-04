package main

import (
	"be-fiber/router"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	api := app.Group("/fiber-go")
	router.AuthRouter(api)

	log.Fatal(app.Listen(":5003"))
}
