package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
)

type IGetUserListCase interface {
	Handle(accessToken string, request requests.GetUserListRequest) ([]*responses.UserResponse, error)
}

type GetUserListCase struct {
	validator          *validator.Validate
	userRepo           repositories.IUserRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
}

func NewGetUserListCase(validator *validator.Validate, userRepo repositories.IUserRepository, makeSureAuthorized helpers.IMakeSureAuthorized) *GetUserListCase {
	return &GetUserListCase{validator: validator, userRepo: userRepo, makeSureAuthorized: makeSureAuthorized}
}

func (gul GetUserListCase) Handle(accessToken string, request requests.GetUserListRequest) ([]*responses.UserResponse, error) {

	_, err := gul.makeSureAuthorized.Handle(accessToken, constants.AdminRole, constants.PlannerRole)
	if err != nil {
		return nil, err
	}

	userList, _ := gul.userRepo.FindManyByRoles(request.Roles, nil)

	response := make([]*responses.UserResponse, len(userList))
	for i, user := range userList {
		userRes, _ := helpers.MapUserResponse(user)
		response[i] = userRes
	}

	return response, nil
}
