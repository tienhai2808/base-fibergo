package handler

import (
	"be-fiber/common"
	"be-fiber/router/request"

	"github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := common.ValidateBody(c, req); err != nil {
			if fiberErr, ok := err.(*fiber.Error); ok {
					return c.Status(fiberErr.Code).JSON(fiber.Map{
							"error": fiberErr.Message,
					})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Đăng nhập thành công",
	})
}
