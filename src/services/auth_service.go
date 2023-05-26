package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
)

var (
	ErrAuthFailed = errors.New("authenticate failed")
)

// $2a$10$pQpD2YRD49hiR7as3UkZQOpuITpdepw9mGwLvw/8MHZF3eYTsNI2a
type IAuthService interface {
	Authenticate(role, email, password string) (*AuthData, error)
}

type AuthService struct {
	userRepo     repositories.IUserRepository
	hashService  IPasswordHashService
	tokenService IAuthTokenService
}

func (as AuthService) Authenticate(role, email, password string) (*AuthData, error) {

	matchingUser, _ := as.userRepo.FindByEmailAndRole(email, role)
	if matchingUser == nil || as.hashService.Compare(matchingUser.PasswordHash, password) {
		return nil, ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(matchingUser.Role, matchingUser.Id)
}
