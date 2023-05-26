package services

import "golang.org/x/crypto/bcrypt"

type IPasswordHashService interface {
	Hash(password string) (string, error)
	Compare(passwordHash, password string) bool
}

type PasswordHashService struct {
}

func NewPasswordHashService() *PasswordHashService {
	return &PasswordHashService{}
}

func (phs PasswordHashService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (phs PasswordHashService) Compare(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
