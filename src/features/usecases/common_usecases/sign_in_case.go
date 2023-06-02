package common_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/go-playground/validator/v10"
)

type ISignInCase interface {
	Handle(request requests.SignInRequest) (*responses.AuthResponse, error)
}

type SignInCase struct {
	userRepo     repositories.IUserRepository
	hashService  services.IPasswordService
	tokenService services.IAuthTokenService
	validator    *validator.Validate
}

func NewSignInCase(userRepo repositories.IUserRepository, hashService services.IPasswordService, tokenService services.IAuthTokenService, validator *validator.Validate) *SignInCase {
	return &SignInCase{userRepo: userRepo, hashService: hashService, tokenService: tokenService, validator: validator}
}

func (sic SignInCase) Handle(request requests.SignInRequest) (*responses.AuthResponse, error) {

	err := sic.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	matchingUser, _ := sic.userRepo.FindByEmailAndRole(request.Email, request.Role)
	if matchingUser == nil || !sic.hashService.Compare(matchingUser.PasswordHash, request.Password) {
		return nil, errs.ErrAuthFailed
	}
	if !matchingUser.IsActive {
		return nil, errs.ErrDisableAccount
	}

	authData, _ := sic.tokenService.GenerateAuthData(matchingUser.Role, matchingUser.Id)

	return &responses.AuthResponse{
		AuthData: *authData,
		Role:     matchingUser.Role,
		Name:     matchingUser.DisplayName,
		Email:    matchingUser.Email,
	}, nil
}
