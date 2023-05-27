package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
)

type IGetUserListCase interface {
	Handle(accessToken string, request requests.PaginationRequest) ([]*responses.UserResponse, error)
}

type GetUserListCase struct {
	tokenService       services.IAuthTokenService
	validator          *validator.Validate
	userRepo           repositories.IUserRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
}

func NewGetUserListCase(tokenService services.IAuthTokenService, validator *validator.Validate, userRepo repositories.IUserRepository, makeSureAuthorized helpers.IMakeSureAuthorized) *GetUserListCase {
	return &GetUserListCase{tokenService: tokenService, validator: validator, userRepo: userRepo, makeSureAuthorized: makeSureAuthorized}
}

func (gul GetUserListCase) Handle(accessToken string, request requests.PaginationRequest) ([]*responses.UserResponse, error) {

	_, err := gul.makeSureAuthorized.Handle(accessToken, constants.AdminRole)
	if err != nil {
		return nil, err
	}

	err = gul.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	userList, _ := gul.userRepo.GetList(request.Page, request.Size, request.OrderBy, request.LastTimestamp)

	response := make([]*responses.UserResponse, len(userList))
	for i, user := range userList {
		response[i] = &responses.UserResponse{
			Id:          user.Id,
			Role:        user.Role,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt.UTC(),
		}
	}

	return response, nil
}
