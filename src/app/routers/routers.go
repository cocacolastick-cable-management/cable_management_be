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

	api := app.Group("/api")

	// common
	common := api.Group("/common")

	common.Post("/sign-in",
		middlewares.BodyParserMiddleware[requests.SignInRequest],
		initalizers.AuthController.SignIn,
		middlewares.GlobalErrorHandleMiddleware)

	common.Get("/users",
		initalizers.AuthorizedMiddleware.Handle(constants.AdminRole, constants.PlannerRole),
		initalizers.CommonUserControllers.GetUserList)

	common.Patch("/with-draws/:id",
		initalizers.AuthorizedMiddleware.Handle(constants.AdminRole, constants.PlannerRole),
		middlewares.BodyParserMiddleware[requests.UpdateWithDrawStatusRequest],
		initalizers.CommonWithDrawController.UpdateWithDrawStatusCase)
	// change password
	// reset password

	// admin
	admin := api.Group("/admin", initalizers.AuthorizedMiddleware.Handle(constants.AdminRole))

	admin.Post("/users",
		middlewares.BodyParserMiddleware[requests.CreateUserRequest],
		initalizers.AdminUserController.CreateUser) // should generate password

	//TODO: disable account

	// planner
	planner := api.Group("/planner", initalizers.AuthorizedMiddleware.Handle(constants.PlannerRole))

	planner.Get("/contracts",
		initalizers.ContractController.GetContractList)

	planner.Get("/with-draws",
		initalizers.WithDrawController.GetWithDrawList)

	//TODO: need to send realtime noti
	planner.Post("/with-draws",
		middlewares.BodyParserMiddleware[requests.CreateWithDrawRequest],
		initalizers.WithDrawController.CreateWithDrawRequest)

	// TODO notification

	// supplier
	supplier := api.Group("/supplier",
		initalizers.AuthorizedMiddleware.Handle(constants.SupplierRole))

	supplier.Get("/contracts",
		initalizers.SupplierContractController.GetContractList)

	supplier.Get("/with-draws",
		initalizers.SupplierWithDrawController.GetWithDrawList)
	// TODO notification

	// TODO contractor
	// TODO get my with draw requests
	// TODO notification

	api.Use(middlewares.GlobalErrorHandleMiddleware)
}
