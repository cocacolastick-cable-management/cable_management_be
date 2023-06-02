package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
)

type ICreateUserCase interface {
	Handle(accessToken string, request requests.CreateUserRequest) (*responses.UserResponse, error)
}

type CreateUserCase struct {
	tokenService       services.IAuthTokenService
	userFac            services.IUserFactory
	userRepo           repositories.IUserRepository
	validator          *validator.Validate
	makeSureAuthorized helpers.IMakeSureAuthorized
	passService        services.IPasswordService
	emailService       services.IEmailService
}

func NewCreateUserCase(tokenService services.IAuthTokenService, userFac services.IUserFactory, userRepo repositories.IUserRepository, validator *validator.Validate, makeSureAuthorized helpers.IMakeSureAuthorized, passService services.IPasswordService, emailService services.IEmailService) *CreateUserCase {
	return &CreateUserCase{tokenService: tokenService, userFac: userFac, userRepo: userRepo, validator: validator, makeSureAuthorized: makeSureAuthorized, passService: passService, emailService: emailService}
}

func (cac CreateUserCase) Handle(accessToken string, request requests.CreateUserRequest) (*responses.UserResponse, error) {

	_, err := cac.makeSureAuthorized.Handle(accessToken, constants.AdminRole)
	if err != nil {
		return nil, err
	}

	err = cac.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	password := cac.passService.GenerateRandomPassword(8)
	newUser, _ := cac.userFac.CreateUser(request.Role, request.DisplayName, request.Email, password)
	_ = cac.userRepo.Insert(newUser)

	go func() {
		_ = cac.emailService.SendEmailNewUser(newUser.Role, newUser.Email, password)
	}()

	return helpers.MapUserResponse(newUser)
}
