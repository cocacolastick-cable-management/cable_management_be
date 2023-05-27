package middlewares

import (
	"github.com/cable_management/cable_management_be/src/app/initalizers"
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func AuthorizeMiddleware(roles ...string) func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")

		accessToken := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			accessToken = authHeader[7:]
		} else {
			return utils.UnAuthenticatedResponse(ctx)
		}

		isValid, claims := initalizers.AuthTokenService.IsAccessTokenValid(accessToken)
		if !isValid {
			return utils.UnAuthenticatedResponse(ctx)
		}

		if !slices.Contains(roles, claims.Role) {
			return utils.UnAuthorizedResponse(ctx)
		}

		ctx.Locals("access-token", accessToken)

		return ctx.Next()
	}
}
