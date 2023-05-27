package middlewares

import "github.com/gofiber/fiber/v2"

func BodyParserMiddleware[T any](ctx *fiber.Ctx) error {

	body := new(T)
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	ctx.Locals("body", *body)

	return ctx.Next()
}
