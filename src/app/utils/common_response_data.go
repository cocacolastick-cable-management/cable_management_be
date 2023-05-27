package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UnAuthenticatedResponse(ctx *fiber.Ctx) error {
	return ctx.Status(401).JSON(Response{
		Message: "authenticate failed",
		Code:    "AF",
	})
}

func UnAuthorizedResponse(ctx *fiber.Ctx) error {
	return ctx.Status(403).JSON(Response{
		Message: "unauthorized",
		Code:    "UA",
	})
}

func ValidationErrorResponse(ctx *fiber.Ctx, validError validator.ValidationErrors) error {
	return ctx.Status(400).JSON(Response{
		Message: "invalidation",
		Code:    "IVL",
		Errors:  ValidationErrorToStruct(validError),
	})
}
