package router

import (
	"be-fiber/handler"
	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app fiber.Router) {
	authApi := app.Group("/auth")
	authApi.Post("/login", handler.Login)
}
