package common_controllers

import (
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
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

	//get body
	request := ctx.Locals("body").(requests.SignInRequest)

	//handle
	var response *responses.AuthResponse
	response, err = ac.signInCase.Handle(request)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	// return happy result
	return ctx.Status(200).JSON(utils.Response{
		Message: "authenticate successfully",
		Code:    "AS",
		Payload: response,
	})
}
