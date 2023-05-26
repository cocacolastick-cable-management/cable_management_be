package common_usecases

import (
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Role     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type ISignInCase interface {
	Handle(request *SignInRequest) (*services.AuthData, error)
}

type SignInCase struct {
	authService services.IAuthService
	validator   *validator.Validate
}

func (sic SignInCase) Handle(request *SignInRequest) (*services.AuthData, error) {

	err := sic.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	return sic.authService.Authenticate(request.Role, request.Email, request.Password)
}
