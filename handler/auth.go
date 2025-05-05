package handler

import (
	"be-fiber/common"
	"be-fiber/router/request"
	"be-fiber/service"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (h *AuthHandler) Test(c fiber.Ctx) error {
	tokenStr := c.Cookies("refresh_token")
	fmt.Println(tokenStr)
	req := new(request.TestRequest)
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

	fmt.Printf(`Nội dung request gửi: %s`, req.Request)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message": "Hello world",
	})
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
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

	user, err := h.svc.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Đăng nhập thành công",
		"user":    user,
	})
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	req := new(request.RegisterRequest)
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

	user, refreshToken, err := h.svc.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(1),
		HTTPOnly: true,  
		Secure:   false,  
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Đăng ký thành công",
		"user":    user,
	})
}