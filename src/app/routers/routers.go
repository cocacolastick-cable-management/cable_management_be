package routers

import (
	"github.com/cable_management/cable_management_be/src/app/initalizers"
	"github.com/cable_management/cable_management_be/src/app/middlewares"
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/gofiber/fiber/v2"
)

func MountApiRouters(app *fiber.App) {

	api := app.Group("api/")

	// common
	common := api.Group("/")
	common.Post("/sign-in",
		middlewares.BodyParserMiddleware[requests.SignInRequest],
		initalizers.AuthController.SignIn,
		middlewares.GlobalErrorHandleMiddleware)

	// admin
	admin := api.Group("/admin")
	admin.Post("/users",
		middlewares.AuthorizeMiddleware(constants.AdminRole),
		middlewares.BodyParserMiddleware[requests.CreateUserRequest],
		initalizers.UserController.CreateUser,
		middlewares.GlobalErrorHandleMiddleware)
}
