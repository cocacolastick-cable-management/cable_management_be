package middlewares

import (
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

type IAuthorizedMiddleware interface {
	Handle(roles ...string) func(ctx *fiber.Ctx) error
}

type AuthorizeMiddleware struct {
	authTokenService services.IAuthTokenService
}

func NewAuthorizeMiddleware(authTokenService services.IAuthTokenService) *AuthorizeMiddleware {
	return &AuthorizeMiddleware{authTokenService: authTokenService}
}

func (am AuthorizeMiddleware) Handle(roles ...string) func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")

		accessToken := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			accessToken = authHeader[7:]
		} else {
			return utils.UnAuthenticatedResponse(ctx)
		}

		isValid, claims := am.authTokenService.IsAccessTokenValid(accessToken)
		if !isValid {
			return utils.UnAuthenticatedResponse(ctx)
		}

		if !slices.Contains(roles, claims.Role) {
			return utils.UnAuthorizedResponse(ctx)
		}

		ctx.Locals(services.AccessTokenTypeName, accessToken)

		return ctx.Next()
	}
}
