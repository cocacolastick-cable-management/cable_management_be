package helpers

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"golang.org/x/exp/slices"
)

type IMakeSureAuthorized interface {
	Handle(accessToken string, targetRoles ...string) (*services.AuthTokenClaims, error)
}

type MakeSureAuthorized struct {
	tokenService services.IAuthTokenService
	userRepo     repositories.IUserRepository
}

func NewMakeSureAuthorized(tokenService services.IAuthTokenService, userRepo repositories.IUserRepository) *MakeSureAuthorized {
	return &MakeSureAuthorized{tokenService: tokenService, userRepo: userRepo}
}

func (msa MakeSureAuthorized) Handle(accessToken string, targetRoles ...string) (*services.AuthTokenClaims, error) {

	isTokenValid, claims := msa.tokenService.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		return nil, errs.ErrAuthFailed
	}

	if !slices.Contains(targetRoles, claims.Role) {
		return nil, errs.ErrUnAuthorized
	}

	matchingUser, err := msa.userRepo.FindById(claims.AccountId)
	if err != nil || matchingUser == nil {
		return nil, errs.ErrAuthFailed
	}

	if !slices.Contains(targetRoles, matchingUser.Role) || !(matchingUser.Role == claims.Role) {
		return nil, errs.ErrUnAuthorized
	}

	return claims, nil
}
