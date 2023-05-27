package middlewares

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandleMiddleware(ctx *fiber.Ctx) error {

	err := ctx.Locals("err").(error)

	if validErrors, ok := err.(validator.ValidationErrors); ok {
		return ctx.Status(400).JSON(utils.ValidationErrorToStruct(validErrors))
	}

	if errors.Is(err, errs.ErrAuthFailed) {
		return utils.UnAuthenticatedResponse(ctx)
	}

	if errors.Is(err, errs.ErrUnAuthorized) {
		return utils.UnAuthorizedResponse(ctx)
	}

	return ctx.Status(500).JSON(err)
}
