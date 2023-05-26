package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/usecase/dtos/requests"
	"github.com/cable_management/cable_management_be/src/usecase/dtos/responses"
	"github.com/go-playground/validator/v10"
)

type IGetUserListCase interface {
	Handle(accessToken string, request *requests.PaginationRequest) ([]responses.UserResponse, error)
}

type GetUserListCase struct {
	tokenService services.IAuthTokenService
	validator    *validator.Validate
	userRepo     repositories.IUserRepository
}

func (gul GetUserListCase) Handle(accessToken string, request requests.PaginationRequest) ([]*responses.UserResponse, error) {

	isTokenValid, claims := gul.tokenService.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		return nil, errs.ErrAuthFailed
	}

	if claims.Role != constants.AdminRole {
		return nil, errs.ErrUnAuthorized
	}

	err := gul.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	userList, _ := gul.userRepo.GetList(request.Page, request.Size, request.LastTimestamp)

	response := make([]*responses.UserResponse, len(userList))
	for i, user := range userList {
		response[i] = &responses.UserResponse{
			Id:        user.Id,
			Role:      user.Role,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	return response, nil
}
