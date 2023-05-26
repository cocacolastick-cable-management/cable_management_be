package adminUsecases

import (
	"github.com/cable_management/cable_management_be/src_test/entities"
	"github.com/cable_management/cable_management_be/src_test/infra/database"
	"github.com/cable_management/cable_management_be/src_test/infra/database/repositories"
	"github.com/cable_management/cable_management_be/src_test/services"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
	"regexp"
	"time"
)

type CreateAccountRequestDto struct {
	Role     string `json:"role" validate:"required,validateRoleExceptAdmin"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,validatePassword"`
}

type CreateAccountResponseDto struct {
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type ICreateAccount interface {
	Handle(accessToken string, createRequest *CreateAccountRequestDto) (*CreateAccountResponseDto, error)
}

type CreateAccount struct {
	validator      *validator.Validate
	accountFac     services.IAccountFactory
	supplierRepo   repositories.ISupplierRepository
	plannerRepo    repositories.IPlannerRepository
	contractorRepo repositories.IContractorRepository
	tokenService   services.IAuthTokenService
}

func NewCreateAccount(
	validator *validator.Validate,
	accountFac services.IAccountFactory,
	supplierRepo repositories.ISupplierRepository,
	plannerRepo repositories.IPlannerRepository,
	contractorRepo repositories.IContractorRepository,
	tokenService services.IAuthTokenService) *CreateAccount {

	return &CreateAccount{
		validator:      validator,
		accountFac:     accountFac,
		supplierRepo:   supplierRepo,
		plannerRepo:    plannerRepo,
		contractorRepo: contractorRepo,
		tokenService:   tokenService}
}

func (ca CreateAccount) Handle(accessToken string, createRequest *CreateAccountRequestDto) (*CreateAccountResponseDto, error) {

	isValid, claims := ca.tokenService.IsAccessTokenValid(accessToken)
	if !isValid || claims.Role != services.AdminRole {
		return nil, services.ErrUnAuthorized
	}

	err := ca.validator.Struct(createRequest)
	if err != nil {
		return nil, err
	}

	account, _ := ca.accountFac.CreateNewAccount(createRequest.Role, createRequest.Email, createRequest.Password, claims.AccountId)

	switch entity := account.(type) {
	case *entities.Planner:
		_ = ca.plannerRepo.Insert(entity)
	case *entities.Supplier:
		_ = ca.supplierRepo.Insert(entity)
	case *entities.Contractor:
		_ = ca.contractorRepo.Insert(entity)
	}

	absAccount := account.(entities.AbstractAccount)
	response := &CreateAccountResponseDto{createRequest.Role, absAccount.Email, absAccount.CreatedAt}

	return response, nil
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasEnoughLength := len(password) >= 8

	return hasNumber && hasUppercase && hasEnoughLength
}

func ValidateRoleExceptAdmin(fl validator.FieldLevel) bool {

	roles := []string{
		services.PlannerRole,
		services.SupplierRole,
		services.ContractorRole}

	value := fl.Field().Interface().(string)

	return slices.Contains[string](roles, value)
}

func ValidateCreateAccountRequest(sl validator.StructLevel) {

	plannerRepo := repositories.NewPlannerRepository(database.DB)
	supplierRepo := repositories.NewSupplierRepository(database.DB)
	contractorRepo := repositories.NewContractorRepository(database.DB)

	createRequest := sl.Current().Interface().(CreateAccountRequestDto)

	switch createRequest.Role {
	case services.PlannerRole:
		matchingPlanner, _ := plannerRepo.FindByEmail(createRequest.Email)
		if matchingPlanner != nil {
			sl.ReportError(createRequest.Email, "email", "Email", "email planner is duplicated", "")
		}
	case services.SupplierRole:
		matchingSupplier, _ := supplierRepo.FindByEmail(createRequest.Email)
		if matchingSupplier != nil {
			sl.ReportError(createRequest.Email, "email", "Email", "email supplier is duplicated", "")
		}
	case services.ContractorRole:
		matchingContractor, _ := contractorRepo.FindByEmail(createRequest.Email)
		if matchingContractor != nil {
			sl.ReportError(createRequest.Email, "email", "Email", "email contractor is duplicated", "")
		}
	}
}
