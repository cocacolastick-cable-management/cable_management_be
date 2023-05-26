package common_controllers

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	SignIn(fiber *fiber.Ctx) error
}

type AuthController struct {
	signInCase common_usecases.ISignInCase
}

func NewAuthController(signIn common_usecases.ISignInCase) *AuthController {
	return &AuthController{signInCase: signIn}
}

func (ac AuthController) SignIn(ctx *fiber.Ctx) error {

	var err error

	//parse request
	request := requests.SignInRequest{}
	err = ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	//handle
	var authData *services.AuthData
	authData, err = ac.signInCase.Handle(request)

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
	return ctx.Status(200).JSON(utils.Response{
		Message: "authenticate successfully",
		Code:    "AS",
		Payload: authData,
	})
}
