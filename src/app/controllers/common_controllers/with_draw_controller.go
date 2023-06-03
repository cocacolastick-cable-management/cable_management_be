package common_controllers

import (
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IWithDrawController interface {
	UpdateWithDrawStatusCase(ctx *fiber.Ctx) error
}

type WithDrawController struct {
	updateWithDrawStatusCase common_usecases.IUpdateWithDrawStatusCase
}

func NewWithDrawController(updateWithDrawStatusCase common_usecases.IUpdateWithDrawStatusCase) *WithDrawController {
	return &WithDrawController{updateWithDrawStatusCase: updateWithDrawStatusCase}
}

func (w WithDrawController) UpdateWithDrawStatusCase(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)

	request := ctx.Locals("body").(requests.UpdateWithDrawStatusRequest)

	var withDrawId uuid.UUID
	withDrawId, err = uuid.Parse(ctx.Params("id"))
	if err != nil {
		ctx.Locals("err", errs.ErrNotFound)
		return ctx.Next()
	}

	err = w.updateWithDrawStatusCase.Handle(accessToken, withDrawId, request)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	return ctx.Status(200).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
	})
}
