package services

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/errs"
)

type IAuthService interface {
	Authenticate(role, email, password string) (*AuthData, error)
}

type AuthService struct {
	userRepo     repositories.IUserRepository
	hashService  IPasswordService
	tokenService IAuthTokenService
}

func NewAuthService(userRepo repositories.IUserRepository, hashService IPasswordService, tokenService IAuthTokenService) *AuthService {
	return &AuthService{userRepo: userRepo, hashService: hashService, tokenService: tokenService}
}

func (as AuthService) Authenticate(role, email, password string) (*AuthData, error) {

	matchingUser, _ := as.userRepo.FindByEmailAndRole(email, role)
	if matchingUser == nil || !as.hashService.Compare(matchingUser.PasswordHash, password) {
		return nil, errs.ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(matchingUser.Role, matchingUser.Id)
}
