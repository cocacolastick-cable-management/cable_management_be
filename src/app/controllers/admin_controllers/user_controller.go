package admin_controllers

import (
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/admin_usecases"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserList(ctx *fiber.Ctx) error
}

type UserController struct {
	createUserCase admin_usecases.ICreateUserCase
}

func NewUserController(createUserCase admin_usecases.ICreateUserCase) *UserController {
	return &UserController{createUserCase: createUserCase}
}

func (uc UserController) CreateUser(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals("access-token").(string)

	request := ctx.Locals("body").(requests.CreateUserRequest)

	//handle
	var userRes *responses.UserResponse
	userRes, err = uc.createUserCase.Handle(accessToken, request)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	// return happy result
	return ctx.Status(201).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
		Payload: userRes,
	})
}

func (uc UserController) GetUserList(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
