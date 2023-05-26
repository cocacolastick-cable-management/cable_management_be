package validations

import (
	"github.com/cable_management/cable_management_be/infras/db"
	"github.com/cable_management/cable_management_be/infras/db/repositories"
	"github.com/cable_management/cable_management_be/src/dtos/requests"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateUserRequest(sl validator.StructLevel) {

	request := sl.Current().Interface().(requests.CreateUserRequest)

	if !ValidateRole(request.Role) {
		sl.ReportError(request.Role, "role", "Role", "invalid", "invalid role")
	}

	if !ValidatePassword(request.Password) {
		sl.ReportError(request.Password, "password", "Password", "invalid", "invalid password format")
	}

	userRepo := repositories.NewUserRepository(db.DB)
	matchingUser, _ := userRepo.FindByEmail(request.Email)
	if matchingUser != nil {
		sl.ReportError(request.Email, "email", "Email", "duplicated", "duplicated email")
	}
}
