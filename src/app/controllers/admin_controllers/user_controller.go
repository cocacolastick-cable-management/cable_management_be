package admin_controllers

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/admin_usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type IUserController interface {
	CreateUser(ctx *fiber.Ctx) error
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
	accessToken := strings.Split(ctx.Get("Authorization"), " ")[1]

	request := requests.CreateUserRequest{}
	err = ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	//handle
	var userRes *responses.UserResponse
	userRes, err = uc.createUserCase.Handle(accessToken, request)

	//check error
	if err != nil {
		if validErrors, ok := err.(validator.ValidationErrors); ok {
			return ctx.JSON(utils.ValidationErrorToStruct(validErrors))
		}
		if errors.Is(err, errs.ErrAuthFailed) {
			return ctx.Status(401).JSON(utils.Response{
				Message: "authenticate failed",
				Code:    "AF",
			})
		}
		panic(err)
	}

	// return happy result
	return ctx.Status(201).JSON(utils.Response{
		Message: "authenticate successfully",
		Code:    "AS",
		Payload: userRes,
	})
}
