package common_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/go-playground/validator/v10"
)

type ISignInCase interface {
	Handle(request requests.SignInRequest) (*services.AuthData, error)
}

type SignInCase struct {
	authService services.IAuthService
	validator   *validator.Validate
}

func NewSignInCase(authService services.IAuthService, validator *validator.Validate) *SignInCase {
	return &SignInCase{authService: authService, validator: validator}
}

func (sic SignInCase) Handle(request requests.SignInRequest) (*services.AuthData, error) {

	err := sic.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	return sic.authService.Authenticate(request.Role, request.Email, request.Password)
}
