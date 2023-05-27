package middlewares

import "github.com/gofiber/fiber/v2"

func QueryParserMiddleware[T any](ctx *fiber.Ctx) error {

	query := new(T)
	err := ctx.QueryParser(query)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	ctx.Locals("query", *query)

	return ctx.Next()
}
