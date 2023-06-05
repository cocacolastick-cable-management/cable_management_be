package initalizers

import (
	"github.com/cable_management/cable_management_be/config"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/contracts/email"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/cable_management/cable_management_be/src/features/usecases/admin_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/common_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/contractor_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/planner_usecases"
	"github.com/cable_management/cable_management_be/src/features/usecases/supplier_usecases"
	featValidations "github.com/cable_management/cable_management_be/src/features/validations"
	"github.com/cable_management/cable_management_be/src/infra/db"
	implRepositories "github.com/cable_management/cable_management_be/src/infra/db/repositories"
	implEmail "github.com/cable_management/cable_management_be/src/infra/email"
	"github.com/cable_management/cable_management_be/src/infra/http/controllers/admin_controllers"
	"github.com/cable_management/cable_management_be/src/infra/http/controllers/common_controllers"
	"github.com/cable_management/cable_management_be/src/infra/http/controllers/contractor_controllers"
	"github.com/cable_management/cable_management_be/src/infra/http/controllers/planner_controller"
	"github.com/cable_management/cable_management_be/src/infra/http/controllers/supplier_controllers"
	"github.com/cable_management/cable_management_be/src/infra/http/middlewares"
	implSse "github.com/cable_management/cable_management_be/src/infra/sse"
	"github.com/cable_management/cable_management_be/src/infra/valider"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	//infra - sse
	SseServer *implSse.ServerSentEventProvider

	//infra - db
	DB *gorm.DB

	//services - contract - email
	EmailConfig *email.EmailConfig
	EmailHelper email.IEmailHelper

	//services - contract - db - repositories
	UserRepo           repositories.IUserRepository
	ContractRepo       repositories.IContractRepository
	WithDrawReqRepo    repositories.IWithDrawRequestRepository
	WithDrawReqHisRepo repositories.IWithDrawRequestHistoryRepository
	NotifRepo          repositories.INotificationRepository

	//infra - valider
	ValidCreateUserReqDep     *featValidations.ValidateCreateUserRequestDependency
	ValidCreateWithDrawReqDep *featValidations.ValidateCreateWithDrawRequestDependency
	Validator                 *validator.Validate

	//services
	PasswordService  services.IPasswordService
	AuthTokenService services.IAuthTokenService
	AuthService      services.IAuthService
	UserFac          services.IUserFactory
	WithDrawReqFac   services.IWithDrawRequestFactory
	EmailService     services.IEmailService
	NotiFac          services.INotifFactory

	//helpers
	MakeSureAuthorized helpers.IMakeSureAuthorized

	//common_usecases
	SignInCase               common_usecases.ISignInCase
	GetUserListCase          common_usecases.IGetUserListCase
	UpdateWithDrawStatusCase common_usecases.IUpdateWithDrawStatusCase

	//admin_usecases
	CreateUserCase admin_usecases.ICreateUserCase

	//planner_usecases
	GetContractListCase planner_usecases.IGetContractListCase
	CreateWithDrawCase  planner_usecases.ICreateWithDrawCase
	GetWithDrawListCase planner_usecases.IGetWithDrawListCase

	//supplier_usecase
	SupplierGetContractListCase supplier_usecases.IGetContractListCase
	SupplierGetWithDrawListCase supplier_usecases.IGetWithDrawListCase

	//contractor_usecase
	ContractorGetWithDrawListCase contractor_usecases.IGetWithDrawListCase

	//common_controllers
	AuthController           common_controllers.IAuthController
	CommonUserControllers    common_controllers.IUserController
	CommonWithDrawController common_controllers.IWithDrawController

	//admin_controllers
	AdminUserController admin_controllers.IUserController

	//planner_controller
	ContractController planner_controller.IContractController
	WithDrawController planner_controller.IWithDrawController

	//supplier_controller
	SupplierContractController supplier_controllers.IContractController
	SupplierWithDrawController supplier_controllers.IWithDrawController

	//contractor_controller
	ContractorWithDrawController contractor_controllers.IWithDrawController

	//middleware
	AuthorizedMiddleware middlewares.IAuthorizedMiddleware
)

func init() {
	//infra - sse
	SseServer = implSse.NewSSEServer()

	//infra - db
	db.Init()
	DB = db.DB

	//infra - email
	EmailConfig = email.NewEmailConfig(config.ENV.SmtpEmail, config.ENV.SmtpHost, config.ENV.SmtpPort, config.ENV.SmtpPassword)
	EmailHelper = implEmail.NewEmailHelper(EmailConfig)

	//services - contract - email
	EmailConfig = email.NewEmailConfig(config.ENV.SmtpEmail, config.ENV.SmtpHost, config.ENV.SmtpPort, config.ENV.SmtpPassword)

	//services - contract - db - repositories
	UserRepo = implRepositories.NewUserRepository(DB)
	ContractRepo = implRepositories.NewContractRepository(DB)
	WithDrawReqRepo = implRepositories.NewWithDrawRequestRepository(DB)
	WithDrawReqHisRepo = implRepositories.NewWithDrawRequestHistoryRepository(DB)
	NotifRepo = implRepositories.NewNotificationRepository(DB)

	//infra - valider
	ValidCreateUserReqDep = featValidations.NewValidateCreateUserRequestDependency(UserRepo)
	ValidCreateWithDrawReqDep = featValidations.NewValidateCreateWithDrawRequestDependency(ContractRepo, UserRepo)
	InitValidator()
	Validator = valider.Validator

	//services
	PasswordService = services.NewPasswordHashService()
	AuthTokenService = services.NewAuthTokenService()
	AuthService = services.NewAuthService(UserRepo, PasswordService, AuthTokenService)
	UserFac = services.NewUserFactory(UserRepo, PasswordService)
	WithDrawReqFac = services.NewWithDrawRequestFactory(ContractRepo, UserRepo)
	EmailService = services.NewEmailService(EmailHelper)
	NotiFac = services.NewNotifFactory(UserRepo, WithDrawReqRepo)

	//helpers
	MakeSureAuthorized = helpers.NewMakeSureAuthorized(AuthTokenService, UserRepo)

	//common_usecases
	SignInCase = common_usecases.NewSignInCase(UserRepo, PasswordService, AuthTokenService, Validator)
	GetUserListCase = common_usecases.NewGetUserListCase(Validator, UserRepo, MakeSureAuthorized)
	UpdateWithDrawStatusCase = common_usecases.NewUpdateWithDrawStatusCase(MakeSureAuthorized, Validator, WithDrawReqRepo, WithDrawReqHisRepo, NotifRepo, NotiFac)

	//admin_usecases
	CreateUserCase = admin_usecases.NewCreateUserCase(AuthTokenService, UserFac, UserRepo, Validator, MakeSureAuthorized, PasswordService, EmailService)

	//planner_usecase
	GetContractListCase = planner_usecases.NewGetContractListCase(ContractRepo, MakeSureAuthorized)
	CreateWithDrawCase = planner_usecases.NewCreateWithDrawCase(WithDrawReqFac, WithDrawReqRepo, WithDrawReqHisRepo, ContractRepo, MakeSureAuthorized, Validator, EmailHelper, NotifRepo, NotiFac)
	GetWithDrawListCase = planner_usecases.NewGetWithDrawListCase(MakeSureAuthorized, WithDrawReqRepo)

	//supplier_usecase
	SupplierGetContractListCase = supplier_usecases.NewGetContractListCase(ContractRepo, MakeSureAuthorized)
	SupplierGetWithDrawListCase = supplier_usecases.NewGetWithDrawListCase(MakeSureAuthorized, WithDrawReqRepo)

	//contractor_usecase
	ContractorGetWithDrawListCase = contractor_usecases.NewGetWithDrawListCase(MakeSureAuthorized, WithDrawReqRepo)

	//common_controllers
	AuthController = common_controllers.NewAuthController(SignInCase)
	CommonUserControllers = common_controllers.NewUserController(GetUserListCase)
	CommonWithDrawController = common_controllers.NewWithDrawController(UpdateWithDrawStatusCase)

	//admin_controllers
	AdminUserController = admin_controllers.NewUserController(CreateUserCase)

	//planner_controller
	ContractController = planner_controller.NewContractController(GetContractListCase)
	WithDrawController = planner_controller.NewWithDrawController(CreateWithDrawCase, GetWithDrawListCase)

	//supplier_controller
	SupplierContractController = supplier_controllers.NewContractController(SupplierGetContractListCase)
	SupplierWithDrawController = supplier_controllers.NewWithDrawController(SupplierGetWithDrawListCase)

	//contractor_controller
	ContractorWithDrawController = contractor_controllers.NewWithDrawController(ContractorGetWithDrawListCase)

	//middleware
	AuthorizedMiddleware = middlewares.NewAuthorizeMiddleware(AuthTokenService)
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
