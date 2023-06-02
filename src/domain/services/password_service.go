package services

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type IPasswordService interface {
	Hash(password string) (string, error)
	Compare(passwordHash, password string) bool
	GenerateRandomPassword(length int) string
}

type PasswordService struct {
}

func NewPasswordHashService() *PasswordService {
	return &PasswordService{}
}

func (phs PasswordService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (phs PasswordService) Compare(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func (phs PasswordService) GenerateRandomPassword(length int) string {

	const (
		letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberBytes  = "0123456789"
		specialBytes = "!@#$%^&*()_+-=[]{}|;:,.<>?`~"
	)

	rand.Seed(time.Now().UnixNano())

	allBytes := letterBytes + numberBytes + specialBytes
	allBytesLen := len(allBytes)

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = allBytes[rand.Intn(allBytesLen)]
	}

	return string(password)
}
