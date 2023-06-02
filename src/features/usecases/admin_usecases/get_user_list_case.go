package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
)

type IGetUserListCase interface {
	Handle(accessToken string) ([]*responses.UserResponse, error)
}

type GetUserListCase struct {
	validator          *validator.Validate
	userRepo           repositories.IUserRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
}

func NewGetUserListCase(validator *validator.Validate, userRepo repositories.IUserRepository, makeSureAuthorized helpers.IMakeSureAuthorized) *GetUserListCase {
	return &GetUserListCase{validator: validator, userRepo: userRepo, makeSureAuthorized: makeSureAuthorized}
}
func (gul GetUserListCase) Handle(accessToken string) ([]*responses.UserResponse, error) {

	_, err := gul.makeSureAuthorized.Handle(accessToken, constants.AdminRole)
	if err != nil {
		return nil, err
	}

	userList, _ := gul.userRepo.GetList(nil)

	response := make([]*responses.UserResponse, len(userList))
	for i, user := range userList {
		response[i] = &responses.UserResponse{
			Id:          user.Id,
			Role:        user.Role,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			IsActive:    user.IsActive,
			CreatedAt:   user.CreatedAt.UTC(),
		}
	}

	return response, nil
}
