package initalizers

import (
	"github.com/cable_management/cable_management_be/src/app/controllers/admin_controllers"
	"github.com/cable_management/cable_management_be/src/app/controllers/common_controllers"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/cable_management/cable_management_be/src/features/usecases/admin_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	featValidations "github.com/cable_management/cable_management_be/src/features/validations"
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
	UserFac          services.IUserFactory

	//helpers
	MakeSureAuthorized helpers.IMakeSureAuthorized

	//common_usecases
	SignInCase common_usecases.ISignInCase

	//admin_usecases
	CreateUserCase  admin_usecases.ICreateUserCase
	GetUserListCase admin_usecases.IGetUserListCase

	//common_controllers
	AuthController common_controllers.IAuthController

	//admin_controllers
	UserController admin_controllers.IUserController
)

func init() {
	//infra - db
	db.Init()
	DB = db.DB

	//infra - valider
	InitValidator()
	Validator = valider.Validator

	//services - contract - db - repositories
	UserRepo = implRepositories.NewUserRepository(DB)

	//services
	HashService = services.NewPasswordHashService()
	AuthTokenService = services.NewAuthTokenService()
	AuthService = services.NewAuthService(UserRepo, HashService, AuthTokenService)
	UserFac = services.NewUserFactory(UserRepo, HashService)

	//helpers
	MakeSureAuthorized = helpers.NewMakeSureAuthorized(AuthTokenService, UserRepo)

	//common_usecases
	SignInCase = common_usecases.NewSignInCase(AuthService, Validator)

	//admin_usecases
	CreateUserCase = admin_usecases.NewCreateUserCase(AuthTokenService, UserFac, UserRepo, Validator, MakeSureAuthorized)
	GetUserListCase = admin_usecases.NewGetUserListCase(AuthTokenService, Validator, UserRepo, MakeSureAuthorized)

	//common_controllers
	AuthController = common_controllers.NewAuthController(SignInCase)
	UserController = admin_controllers.NewUserController(CreateUserCase, GetUserListCase)
}

func InitValidator() {
	valider.Init()
	valider.SetStructValidations([]struct {
		Fn   validator.StructLevelFunc
		Type any
	}{
		{
			Fn:   featValidations.ValidateCreateUserRequest,
			Type: requests.CreateUserRequest{},
		},
	})
}
