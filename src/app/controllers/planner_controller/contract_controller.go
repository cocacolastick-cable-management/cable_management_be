package planner_controller

import (
	"github.com/cable_management/cable_management_be/src/app/utils"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/planner_usecases"
	"github.com/gofiber/fiber/v2"
)

type IContractController interface {
	GetContractList(ctx *fiber.Ctx) error
}

type ContractController struct {
	getContractListCase planner_usecases.IGetContractListCase
}

func NewContractController(getContractListCase planner_usecases.IGetContractListCase) *ContractController {
	return &ContractController{getContractListCase: getContractListCase}
}

func (cc ContractController) GetContractList(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)
	query := ctx.Locals("query").(requests.PaginationRequest)

	//handle
	var contractListRes []*responses.ContractResponse
	contractListRes, err = cc.getContractListCase.Handle(accessToken, query)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	return ctx.Status(200).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
		Payload: contractListRes,
	})
}
