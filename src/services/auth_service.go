package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/infra/database/repositories"
)

var (
	ErrAuthFailed   = errors.New("authenticate failed")
	ErrUnAuthorized = errors.New("unauthorized")
)

type IAuthService interface {
	Authenticate(role, email, password string) (*AuthData, error)
	authAdmin(email, password string) (*AuthData, error)
	authPlanner(email, password string) (*AuthData, error)
	authSupplier(email, password string) (*AuthData, error)
	authContractor(email, password string) (*AuthData, error)
}

type AuthService struct {
	passwordHashService IPasswordHashService
	tokenService        IAuthTokenService
	adminRepo           repositories.IAdminRepository
	plannerRepo         repositories.IPlannerRepository
	supplierRepo        repositories.ISupplierRepository
	contractorRepo      repositories.IContractorRepository
}

func NewAuthService(
	passwordHashService IPasswordHashService,
	tokenService IAuthTokenService,
	adminRepo repositories.IAdminRepository,
	plannerRepo repositories.IPlannerRepository,
	supplierRepo repositories.ISupplierRepository,
	contractorRepo repositories.IContractorRepository) *AuthService {

	return &AuthService{
		passwordHashService: passwordHashService,
		tokenService:        tokenService,
		adminRepo:           adminRepo,
		plannerRepo:         plannerRepo,
		supplierRepo:        supplierRepo,
		contractorRepo:      contractorRepo}
}

func (as AuthService) Authenticate(role, email, password string) (*AuthData, error) {

	switch role {
	case AdminRole:
		return as.authAdmin(email, password)
	case SupplierRole:
		return as.authSupplier(email, password)
	case PlannerRole:
		return as.authPlanner(email, password)
	case ContractorRole:
		return as.authContractor(email, password)
	default:
		return nil, ErrIncorrectRole
	}
}

func (as AuthService) authAdmin(email, password string) (*AuthData, error) {

	matchingAdmin, _ := as.adminRepo.FindByEmail(email)
	if matchingAdmin == nil {
		return nil, ErrAuthFailed
	}
	isPasswordMatch := as.passwordHashService.Compare(matchingAdmin.PasswordHash, password)
	if !isPasswordMatch {
		return nil, ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(AdminRole, matchingAdmin.Id)
}

func (as AuthService) authPlanner(email, password string) (*AuthData, error) {
	matchingPlanner, _ := as.plannerRepo.FindByEmail(email)
	if matchingPlanner == nil {
		return nil, ErrAuthFailed
	}
	isPasswordMatch := as.passwordHashService.Compare(matchingPlanner.PasswordHash, password)
	if !isPasswordMatch {
		return nil, ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(PlannerRole, matchingPlanner.Id)
}

func (as AuthService) authSupplier(email, password string) (*AuthData, error) {
	matchingSupplier, _ := as.supplierRepo.FindByEmail(email)
	if matchingSupplier == nil {
		return nil, ErrAuthFailed
	}
	isPasswordMatch := as.passwordHashService.Compare(matchingSupplier.PasswordHash, password)
	if !isPasswordMatch {
		return nil, ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(SupplierRole, matchingSupplier.Id)
}

func (as AuthService) authContractor(email, password string) (*AuthData, error) {
	matchingContractor, _ := as.contractorRepo.FindByEmail(email)
	if matchingContractor == nil {
		return nil, ErrAuthFailed
	}
	isPasswordMatch := as.passwordHashService.Compare(matchingContractor.PasswordHash, password)
	if !isPasswordMatch {
		return nil, ErrAuthFailed
	}

	return as.tokenService.GenerateAuthData(ContractorRole, matchingContractor.Id)
}
