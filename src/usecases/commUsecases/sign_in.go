package commUsecases

import (
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

type SignInRequestDto struct {
	Role     string `validate:"validateRole"`
	Email    string
	Password string
}

type ISignIn interface {
	Handle(signInRequest *SignInRequestDto) (*services.AuthData, error)
}

type SignIn struct {
	authService services.IAuthService
	validator   *validator.Validate
}

func NewSignIn(authService services.IAuthService, validator *validator.Validate) *SignIn {
	return &SignIn{authService: authService, validator: validator}
}

func (si SignIn) Handle(signInRequest *SignInRequestDto) (*services.AuthData, error) {

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

func ValidateRole(fl validator.FieldLevel) bool {

	roles := []string{
		services.AdminRole,
		services.PlannerRole,
		services.SupplierRole,
		services.ContractorRole}

	value := fl.Field().Interface().(string)

	return slices.Contains[string](roles, value)
}
