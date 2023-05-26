package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidJwtToken = errors.New("invalid jwt-token")
)

const (
	AccessTokenExpire  = time.Hour * 2
	RefreshTokenExpire = time.Hour * 24 * 30
)

const (
	AccessTokenTypeName  = "access-token"
	RefreshTokenTypeName = "refresh-token"
)

var (
	jwtSecret = config.ENV.JwtSecret
)

type AuthData struct {
	AccessToken  string
	RefreshToken string
}

type AuthTokenClaims struct {
	jwt.RegisteredClaims
	Role      string
	Type      string
	AccountId uuid.UUID
}

type IAuthTokenService interface {
	GenerateAuthData(role string, accountId uuid.UUID) (*AuthData, error)
	IsAccessTokenValid(accessToken string) (bool, *AuthTokenClaims)
	IsRefreshTokenValid(refreshToken string) (bool, *AuthTokenClaims)
	ParseToClaims(tokenStr string) (*AuthTokenClaims, error)
}

type AuthTokenService struct {
}

func NewAuthTokenService() *AuthTokenService {
	return &AuthTokenService{}
}

func (ats AuthTokenService) GenerateAuthData(role string, accountId uuid.UUID) (*AuthData, error) {

	//TODO validate role???

	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		AuthTokenClaims{
			Role:      role,
			AccountId: accountId,
			Type:      AccessTokenTypeName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
			},
		})

	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		AuthTokenClaims{
			Role:      role,
			AccountId: accountId,
			Type:      RefreshTokenTypeName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire)),
			},
		})

	accessTokenStr, _ := accessToken.SignedString([]byte(jwtSecret))
	refreshTokenStr, _ := refreshToken.SignedString([]byte(jwtSecret))

	return &AuthData{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}

func (ats AuthTokenService) IsAccessTokenValid(accessToken string) (bool, *AuthTokenClaims) {

	claims, err := ats.ParseToClaims(accessToken)
	if err != nil {
		return false, nil
	}

	tokenType := claims.Type
	if tokenType != AccessTokenTypeName {
		return false, nil
	}

	return true, claims
}

func (ats AuthTokenService) IsRefreshTokenValid(refreshToken string) (bool, *AuthTokenClaims) {

	claims, err := ats.ParseToClaims(refreshToken)
	if err != nil {
		return false, nil
	}

	tokenType := claims.Type
	if tokenType != RefreshTokenTypeName {
		return false, nil
	}

	return true, claims
}

func (ats AuthTokenService) ParseToClaims(tokenStr string) (*AuthTokenClaims, error) {

	claims := &AuthTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidJwtToken
	}

	return claims, nil
}
