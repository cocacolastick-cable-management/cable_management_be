package services

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
	"github.com/cable_management/cable_management_be/src/validations"
	"time"
)

type IUserFactory interface {
	CreateUser(role, email, password string) (*entities.User, error)
}

type UserFactory struct {
	userRepo    repositories.IUserRepository
	hashService IPasswordHashService
}

func NewUserFactory(userRepo repositories.IUserRepository, hashService IPasswordHashService) *UserFactory {
	return &UserFactory{userRepo: userRepo, hashService: hashService}
}

func (uf UserFactory) CreateUser(role, email, password string) (*entities.User, error) {

	if !validations.ValidateRole(role) {
		return nil, ErrInvalidRole
	}

	matchingUser, _ := uf.userRepo.FindByEmail(email)
	if matchingUser != nil {
		return nil, ErrDuplicatedEmail
	}

	passwordHash, _ := uf.hashService.Hash(password)

	return entities.NewUser(role, email, passwordHash, time.Now()), nil
}
