package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func ValidateBody(c fiber.Ctx, req any) error {
	validate := validator.New()

	if err := c.Bind().JSON(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Yêu cầu không đúng định dạng JSON")
	}

	if err := validate.Struct(req); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Lỗi xác thực dữ liệu")
		}

		var errorMessages []string
		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()
			errorMessages = append(errorMessages, getVietnameseErrorMessage(field, tag, err.Param()))
		}

		return fiber.NewError(fiber.StatusBadRequest, strings.Join(errorMessages, "; "))
	}

	return nil
}

func getVietnameseErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Trường %s là bắt buộc", field)
	case "email":
		return fmt.Sprintf("Trường %s phải là email hợp lệ", field)
	case "min":
		return fmt.Sprintf("Trường %s phải có ít nhất %s ký tự", field, param)
	case "gte":
		return fmt.Sprintf("Trường %s phải lớn hơn hoặc bằng %s", field, param)
	default:
		return fmt.Sprintf("Trường %s không hợp lệ", field)
	}
}