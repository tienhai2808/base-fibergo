package router

import (
	"be-fiber/handler"
	"be-fiber/middleware"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(router fiber.Router, handler *handler.AuthHandler) {
	auth := router.Group("/auth")

	auth.Post("/test", middleware.AuthRequired(), func(c fiber.Ctx) error {
		return c.SendString("Test handler")
	})
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
