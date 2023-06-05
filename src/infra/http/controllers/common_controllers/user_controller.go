package common_controllers

import (
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/cable_management/cable_management_be/src/infra/http/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type IUserController interface {
	GetUserList(ctx *fiber.Ctx) error
}

type UserController struct {
	getUserList common_usecases.IGetUserListCase
}

func NewUserController(getUserList common_usecases.IGetUserListCase) *UserController {
	return &UserController{getUserList: getUserList}
}

func (uc UserController) GetUserList(ctx *fiber.Ctx) error {

	var err error

	//parse request
	accessToken := ctx.Locals(services.AccessTokenTypeName).(string)
	roles := strings.Split(ctx.Query("roles"), ",")
	request := requests.GetUserListRequest{Roles: roles}

	//handle
	var userListRes []*responses.UserResponse
	userListRes, err = uc.getUserList.Handle(accessToken, request)

	//check error
	if err != nil {
		//check for other cases before pass to global error handler
		ctx.Locals("err", err)
		return ctx.Next()
	}

	// return happy result
	return ctx.Status(200).JSON(utils.Response{
		Message: "Success",
		Code:    "OK",
		Payload: userListRes,
	})
}
