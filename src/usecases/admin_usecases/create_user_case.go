package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/constants"
	"github.com/cable_management/cable_management_be/src/errs"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/cable_management/cable_management_be/src/usecases/_commons/requests"
	"github.com/cable_management/cable_management_be/src/usecases/_commons/responses"
	"github.com/go-playground/validator/v10"
)

type ICreateUserCase interface {
	Handle(accessToken string, request *requests.CreateUserRequest) (*responses.UserResponse, error)
}

type CreateUserCase struct {
	tokenService services.IAuthTokenService
	userFac      services.IUserFactory
	userRepo     repositories.IUserRepository
	validator    *validator.Validate
}

func NewCreateUserCase(tokenService services.IAuthTokenService, userFac services.IUserFactory, userRepo repositories.IUserRepository, validator *validator.Validate) *CreateUserCase {
	return &CreateUserCase{tokenService: tokenService, userFac: userFac, userRepo: userRepo, validator: validator}
}

func (cac CreateUserCase) Handle(accessToken string, request *requests.CreateUserRequest) (*responses.UserResponse, error) {

	isTokenValid, claims := cac.tokenService.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		return nil, errs.ErrAuthFailed
	}

	if claims.Role != constants.AdminRole {
		return nil, errs.ErrUnAuthorized
	}

	err := cac.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	newUser, _ := cac.userFac.CreateUser(request.Role, request.Email, request.Password)
	_ = cac.userRepo.Insert(newUser)

	return nil, nil
}
