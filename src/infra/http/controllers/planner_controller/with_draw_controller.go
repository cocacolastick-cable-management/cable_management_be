package planner_controller

import (
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/planner_usecases"
	"github.com/cable_management/cable_management_be/src/infra/http/utils"
	"github.com/gofiber/fiber/v2"
)

type IWithDrawController interface {
	CreateWithDrawRequest(ctx *fiber.Ctx) error
	GetWithDrawList(ctx *fiber.Ctx) error
}

type WithDrawController struct {
	createWithDrawCase  planner_usecases.ICreateWithDrawCase
	getWithDrawListCase planner_usecases.IGetWithDrawListCase
}

func NewWithDrawController(createWithDrawCase planner_usecases.ICreateWithDrawCase, getWithDrawListCase planner_usecases.IGetWithDrawListCase) *WithDrawController {
	return &WithDrawController{createWithDrawCase: createWithDrawCase, getWithDrawListCase: getWithDrawListCase}
}

func (wdc WithDrawController) CreateWithDrawRequest(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)
	request := ctx.Locals("body").(requests.CreateWithDrawRequest)

	//handle
	var withDrawRes *responses.WithDrawResponse
	withDrawRes, err = wdc.createWithDrawCase.Handle(accessToken, request)

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
		Payload: withDrawRes,
	})
}

func (wdc WithDrawController) GetWithDrawList(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)

	//handle
	var withDrawListRes []*responses.WithDrawResponse
	withDrawListRes, err = wdc.getWithDrawListCase.Handle(accessToken)

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
