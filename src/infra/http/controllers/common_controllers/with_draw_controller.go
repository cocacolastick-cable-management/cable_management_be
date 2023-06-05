package common_controllers

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/cable_management/cable_management_be/src/infra/http/utils"
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

	var withDrawRes *responses.WithDrawResponse
	withDrawRes, err = w.updateWithDrawStatusCase.Handle(accessToken, withDrawId, request)

	//check error
	if err != nil {
		if errors.Is(err, errs.ErrUnAuthorized) && withDrawRes != nil {
			return ctx.Status(400).JSON(utils.Response{
				Message: "you can not update the current status",
				Code:    "BP",
				Payload: withDrawRes,
			})
		}
		ctx.Locals("err", err)
		return ctx.Next()
	}

	return ctx.Status(200).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
		Payload: withDrawRes,
	})
}
