package services

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/validations"
	"time"
)

type IUserFactory interface {
	CreateUser(role, displayName, email, password string) (*entities.User, error)
}

type UserFactory struct {
	userRepo    repositories.IUserRepository
	hashService IPasswordHashService
}

func NewUserFactory(userRepo repositories.IUserRepository, hashService IPasswordHashService) *UserFactory {
	return &UserFactory{userRepo: userRepo, hashService: hashService}
}

func (uf UserFactory) CreateUser(role, displayName, email, password string) (*entities.User, error) {

	if !validations.ValidateRole(role) {
		return nil, errs.ErrInvalidRole
	}

	if !validations.ValidatePassword(password) {
		return nil, errs.ErrInvalidPasswordFormat
	}

	matchingUser, _ := uf.userRepo.FindByEmail(email)
	if matchingUser != nil {
		return nil, errs.ErrDuplicatedEmail
	}

	passwordHash, _ := uf.hashService.Hash(password)

	return entities.NewUser(role, displayName, email, passwordHash, time.Now()), nil
}
