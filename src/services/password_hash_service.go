package services

import "golang.org/x/crypto/bcrypt"

type IPasswordHashService interface {
	Hash(password string) (string, error)
}

type PasswordHashService struct {
}

func (phs PasswordHashService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
