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
	admin := api.Group("/admin", initalizers.AuthorizedMiddleware.Handle(constants.AdminRole))

	admin.Post("/users",
		middlewares.BodyParserMiddleware[requests.CreateUserRequest],
		initalizers.UserController.CreateUser)

	admin.Get("/users",
		middlewares.QueryParserMiddleware[requests.PaginationRequest],
		initalizers.UserController.GetUserList)

	admin.Use(middlewares.GlobalErrorHandleMiddleware)

	// planner
	planner := api.Group("/planner", initalizers.AuthorizedMiddleware.Handle(constants.PlannerRole))

	planner.Get("/contracts",
		middlewares.QueryParserMiddleware[requests.PaginationRequest],
		initalizers.ContractController.GetContractList)

	planner.Post("/with-draws",
		middlewares.BodyParserMiddleware[requests.CreateWithDrawRequest],
		initalizers.WithDrawController.CreateWithDrawRequest)

	planner.Use(middlewares.GlobalErrorHandleMiddleware)
}
