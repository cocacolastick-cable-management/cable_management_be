package validations

import (
	"github.com/cable_management/cable_management_be/src/domain/validations"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/infra/db"
	"github.com/cable_management/cable_management_be/src/infra/db/repositories"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateUserRequest(sl validator.StructLevel) {

	request := sl.Current().Interface().(requests.CreateUserRequest)

	if !validations.ValidateRole(request.Role) {
		sl.ReportError(request.Role, "Role", "Role", "invalid", "invalid role")
	}

	if !validations.ValidatePassword(request.Password) {
		sl.ReportError(request.Password, "Password", "Password", "invalid", "invalid password format")
	}

	userRepo := repositories.NewUserRepository(db.DB)
	matchingUser, _ := userRepo.FindByEmail(request.Email)
	if matchingUser != nil {
		sl.ReportError(request.Email, "Email", "Email", "duplicated", "duplicated email")
	}
}
