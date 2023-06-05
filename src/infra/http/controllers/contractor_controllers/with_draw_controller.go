package contractor_controllers

import (
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/contractor_usecases"
	"github.com/cable_management/cable_management_be/src/infra/http/utils"
	"github.com/gofiber/fiber/v2"
)

type IWithDrawController interface {
	GetWithDrawList(ctx *fiber.Ctx) error
}

type WithDrawController struct {
	getWithDrawListCase contractor_usecases.IGetWithDrawListCase
}

func NewWithDrawController(getWithDrawListCase contractor_usecases.IGetWithDrawListCase) *WithDrawController {
	return &WithDrawController{getWithDrawListCase: getWithDrawListCase}
}

func (w WithDrawController) GetWithDrawList(ctx *fiber.Ctx) error {
	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)

	//handle
	var withDrawListRes []*responses.WithDrawResponse
	withDrawListRes, err = w.getWithDrawListCase.Handle(accessToken)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	return ctx.Status(200).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
		Payload: withDrawListRes,
	})
}
