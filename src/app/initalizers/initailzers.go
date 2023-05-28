package initalizers

import (
	"github.com/cable_management/cable_management_be/src/app/controllers/admin_controllers"
	"github.com/cable_management/cable_management_be/src/app/controllers/common_controllers"
	"github.com/cable_management/cable_management_be/src/app/controllers/planner_controller"
	"github.com/cable_management/cable_management_be/src/app/middlewares"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/cable_management/cable_management_be/src/features/usecases/admin_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/planner_usecases"
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

	//services - contract - db - repositories
	UserRepo           repositories.IUserRepository
	ContractRepo       repositories.IContractRepository
	WithDrawReqRepo    repositories.IWithDrawRequestRepository
	WithDrawReqHisRepo repositories.IWithDrawRequestHistoryRepository

	//services
	HashService      services.IPasswordHashService
	AuthTokenService services.IAuthTokenService
	AuthService      services.IAuthService
	UserFac          services.IUserFactory
	WithDrawReqFac   services.IWithDrawRequestFactory

	//helpers
	MakeSureAuthorized helpers.IMakeSureAuthorized

	//common_usecases
	SignInCase common_usecases.ISignInCase

	//admin_usecases
	CreateUserCase  admin_usecases.ICreateUserCase
	GetUserListCase admin_usecases.IGetUserListCase

	//planner_usecases
	GetContractListCase planner_usecases.IGetContractListCase
	CreateWithDrawCasae planner_usecases.ICreateWithDrawCase

	//common_controllers
	AuthController common_controllers.IAuthController

	//admin_controllers
	UserController admin_controllers.IUserController

	//planner_controller
	ContractController planner_controller.IContractController

	//middleware
	AuthorizedMiddleware middlewares.IAuthorizedMiddleware

	//infra - valider
	ValidCreateUserReqDep     *featValidations.ValidateCreateUserRequestDependency
	ValidCreateWithDrawReqDep *featValidations.ValidateCreateWithDrawRequestDependency
	Validator                 *validator.Validate
)

func init() {
	//infra - db
	db.Init()
	DB = db.DB

	//services - contract - db - repositories
	UserRepo = implRepositories.NewUserRepository(DB)
	ContractRepo = implRepositories.NewContractRepository(DB)
	WithDrawReqRepo = implRepositories.NewWithDrawRequestRepository(DB)
	WithDrawReqHisRepo = implRepositories.NewWithDrawRequestHistoryRepository(DB)

	//services
	HashService = services.NewPasswordHashService()
	AuthTokenService = services.NewAuthTokenService()
	AuthService = services.NewAuthService(UserRepo, HashService, AuthTokenService)
	UserFac = services.NewUserFactory(UserRepo, HashService)
	WithDrawReqFac = services.NewWithDrawRequestFactory(ContractRepo, UserRepo)

	//helpers
	MakeSureAuthorized = helpers.NewMakeSureAuthorized(AuthTokenService, UserRepo)

	//common_usecases
	SignInCase = common_usecases.NewSignInCase(AuthService, Validator)

	//admin_usecases
	CreateUserCase = admin_usecases.NewCreateUserCase(AuthTokenService, UserFac, UserRepo, Validator, MakeSureAuthorized)
	GetUserListCase = admin_usecases.NewGetUserListCase(Validator, UserRepo, MakeSureAuthorized)

	//planner_usecase
	GetContractListCase = planner_usecases.NewGetContractListCase(ContractRepo, MakeSureAuthorized, Validator)
	CreateWithDrawCasae = planner_usecases.NewCreateWithDrawCase(WithDrawReqFac, WithDrawReqRepo, ContractRepo, MakeSureAuthorized, Validator)

	//common_controllers
	AuthController = common_controllers.NewAuthController(SignInCase)

	//admin_controllers
	UserController = admin_controllers.NewUserController(CreateUserCase, GetUserListCase)

	//planner_controllers
	ContractController = planner_controller.NewContractController(GetContractListCase)

	//middleware
	AuthorizedMiddleware = middlewares.NewAuthorizeMiddleware(AuthTokenService)

	//infra - valider
	ValidCreateUserReqDep = featValidations.NewValidateCreateUserRequestDependency(UserRepo)
	ValidCreateWithDrawReqDep = featValidations.NewValidateCreateWithDrawRequestDependency(ContractRepo, UserRepo)
	InitValidator()
	Validator = valider.Validator
}

func InitValidator() {
	valider.Init()
	valider.SetStructValidations([]struct {
		Fn   validator.StructLevelFunc
		Type any
	}{
		{
			Fn:   featValidations.ValidateCreateUserRequest(ValidCreateUserReqDep),
			Type: requests.CreateUserRequest{},
		},
		{
			Fn:   featValidations.ValidateCreateWithDrawRequest(ValidCreateWithDrawReqDep),
			Type: requests.CreateWithDrawRequest{},
		},
	})
}
