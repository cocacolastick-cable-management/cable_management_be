package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/entities"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
	"time"
)

var (
	ErrDuplicatedEmail = errors.New("email is already in use")
)

type IUserFactory interface {
	CreateUser(email, password string) (*entities.User, error)
}

type UserFactory struct {
	userRepo    repositories.IUserRepository
	hashService IPasswordHashService
}

func NewUserFactory(userRepo repositories.IUserRepository, hashService IPasswordHashService) *UserFactory {
	return &UserFactory{userRepo: userRepo, hashService: hashService}
}

func (uf UserFactory) CreateUser(email, password string) (*entities.User, error) {

	matchingUser, _ := uf.userRepo.FindByEmail(email)
	if matchingUser != nil {
		return nil, ErrDuplicatedEmail
	}

	passwordHash, _ := uf.hashService.Hash(password)

	return entities.NewUser(email, passwordHash, time.Now()), nil
}
