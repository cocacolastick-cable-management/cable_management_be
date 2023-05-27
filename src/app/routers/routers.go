package routers

import (
	"github.com/cable_management/cable_management_be/src/app/initalizers"
	"github.com/gofiber/fiber/v2"
)

func MountApiRouters(app *fiber.App) {
	api := app.Group("api/")

	common := api.Group("/")
	common.Post("/sign-in", initalizers.AuthController.SignIn)

	admin := api.Group("/admin")
	admin.Post("/users", initalizers.UserController.CreateUser)
}
