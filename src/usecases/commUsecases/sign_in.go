package commUsecases

import (
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
)

type SignInRequestDto struct {
	Role     string `validate:"validateRole"`
	Email    string
	Password string
}

type ISignIn interface {
	Handle(signInRequest SignInRequestDto) (*services.AuthData, error)
}

type SignIn struct {
	authService services.IAuthService
	validator   *validator.Validate
}

func (si SignIn) Handle(signInRequest SignInRequestDto) (*services.AuthData, error) {

	var err error
	
	err = si.validator.Struct(signInRequest)
	if err != nil {
		return nil, err
	}

	var authData *services.AuthData
	authData, err = si.authService.Authenticate(
		signInRequest.Role,
		signInRequest.Email,
		signInRequest.Password)

	return authData, err
}
