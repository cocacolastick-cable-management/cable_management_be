package commUsecases

import "github.com/cable_management/cable_management_be/src/services"

type IRefreshToken interface {
	Handle(refreshToken string) (*services.AuthData, error)
}

type RefreshToken struct {
	tokenService services.IAuthTokenService
}

func (rt RefreshToken) Handle(refreshToken string) (*services.AuthData, error) {

	isValid, claims := rt.tokenService.IsRefreshTokenValid(refreshToken)
	if !isValid {
		return nil, services.ErrInvalidJwtToken
	}

	return rt.tokenService.GenerateAuthData(claims.Role, claims.AccountId)
}
