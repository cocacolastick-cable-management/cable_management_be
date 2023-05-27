package utils

import "github.com/gofiber/fiber/v2"

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
