package initalizers

import (
	"github.com/cable_management/cable_management_be/src/app/controllers/common_controllers"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/cable_management/cable_management_be/src/infra/db"
	implRepositories "github.com/cable_management/cable_management_be/src/infra/db/repositories"
	"github.com/cable_management/cable_management_be/src/infra/valider"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	//infra - db
	DB *gorm.DB

	//infra - valider
	Validator *validator.Validate

	//services - contract - db - repositories
	UserRepo repositories.IUserRepository

	//services
	HashService      services.IPasswordHashService
	AuthTokenService services.IAuthTokenService
	AuthService      services.IAuthService

	//common_usecases
	SignInCase common_usecases.ISignInCase

	//common_controllers
	AuthController common_controllers.IAuthController
)

func init() {
	//infra - db
	db.Init()
	DB = db.DB

	//infra - valider
	valider.Init()
	valider.SetStructValidations(nil)
	Validator = valider.Validator

	//services - contract - db - repositories
	UserRepo = implRepositories.NewUserRepository(DB)

	//services
	HashService = services.NewPasswordHashService()
	AuthTokenService = services.NewAuthTokenService()
	AuthService = services.NewAuthService(UserRepo, HashService, AuthTokenService)

	//common_usecases
	SignInCase = common_usecases.NewSignInCase(AuthService, Validator)

	//common_controllers
	AuthController = common_controllers.NewAuthController(SignInCase)
}
