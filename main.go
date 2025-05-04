package main

import (
	"be-fiber/config"
	"be-fiber/database"
	"be-fiber/handler"
	"be-fiber/repository"
	"be-fiber/router"
	"be-fiber/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Fiber GO",
	})

	cfg := config.LoadConfig()
	mongo := database.ConnectMongoDB(&cfg.Mongo)
	defer mongo.Close()

	repo := repository.NewUserRepository(mongo)
	svc := service.NewAuthService(repo)
	hdl := handler.NewAuthHandler(svc)

	api := app.Group("/fiber-go")
	router.AuthRouter(api, hdl)

	port := fmt.Sprintf(":%s", cfg.App.Port)
	log.Fatal(app.Listen(port))
}