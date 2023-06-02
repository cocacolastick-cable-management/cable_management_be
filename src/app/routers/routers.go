package routers

import (
	"github.com/cable_management/cable_management_be/src/app/initalizers"
	"github.com/cable_management/cable_management_be/src/app/middlewares"
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func MountApiRouters(app *fiber.App) {

	app.Use(cors.New())

	api := app.Group("api/")

	// common
	common := api.Group("/")
	common.Post("/sign-in",
		middlewares.BodyParserMiddleware[requests.SignInRequest],
		initalizers.AuthController.SignIn,
		middlewares.GlobalErrorHandleMiddleware)
	// change password

	// admin
	admin := api.Group("/admin", initalizers.AuthorizedMiddleware.Handle(constants.AdminRole))

	admin.Post("/users",
		middlewares.BodyParserMiddleware[requests.CreateUserRequest],
		initalizers.UserController.CreateUser) // should generate password

	admin.Get("/users",
		initalizers.UserController.GetUserList)
	//feat: disable account instead of remove it

	admin.Use(middlewares.GlobalErrorHandleMiddleware)

	// planner
	planner := api.Group("/planner", initalizers.AuthorizedMiddleware.Handle(constants.PlannerRole))

	planner.Get("/contracts",
		initalizers.ContractController.GetContractList)

	planner.Get("/with-draws",
		initalizers.WithDrawController.GetWithDrawList)

	planner.Post("/with-draws",
		middlewares.BodyParserMiddleware[requests.CreateWithDrawRequest],
		initalizers.WithDrawController.CreateWithDrawRequest)
	// cancel requests

	// supplier
	// get my contracts
	// get my with draw requests
	// update with draw request to ready

	// contractor
	// get my with draw requests
	// update with draw request to collected

	planner.Use(middlewares.GlobalErrorHandleMiddleware)
}
