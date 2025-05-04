package router

import (
	"be-fiber/handler"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(router fiber.Router, handler *handler.AuthHandler) {
	auth := router.Group("/auth")
	
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}