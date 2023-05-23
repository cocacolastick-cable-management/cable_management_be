package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/src/entities"
	"github.com/cable_management/cable_management_be/src/infra/database/repositories"
	"time"
)

const (
	AdminRole      = "admin"
	PlannerRole    = "planner"
	SupplierRole   = "supplier"
	ContractorRole = "contractor"
)

var (
	ErrIncorrectRole             = errors.New("incorrect role")
	ErrDuplicatedAdminEmail      = errors.New("duplicated admin email")
	ErrDuplicatedSupplierEmail   = errors.New("duplicated supplier email")
	ErrDuplicatedPlannerEmail    = errors.New("duplicated planner email")
	ErrDuplicatedContractorEmail = errors.New("duplicated contractor email")
)

type IAccountFactory interface {
	CreateNewAccount(role, email, password string) (entities.IAbstractAccount, error)
	createAdmin(email, password string) (*entities.Admin, error)
	createSupplier(email, password string) (*entities.Supplier, error)
	createPlanner(email, password string) (*entities.Planner, error)
	createContractor(email, password string) (*entities.Contractor, error)
}

type AccountFactory struct {
	passwordHashService IPasswordHashService
	adminRepo           repositories.IAdminRepository
	supplierRepo        repositories.ISupplierRepository
	plannerRepo         repositories.IPlannerRepository
	contractorRepo      repositories.IContractorRepository
}

func (af AccountFactory) CreateNewAccount(role, email, password string) (entities.IAbstractAccount, error) {

	switch role {
	// TODO should I remove AdminRole on production ???
	case AdminRole:
		return af.createAdmin(email, password)
	case SupplierRole:
		return af.createSupplier(email, password)
	case PlannerRole:
		return af.createPlanner(email, password)
	case ContractorRole:
		return af.createContractor(email, password)
	default:
		return nil, ErrIncorrectRole
	}
}

func (af AccountFactory) createAdmin(email, password string) (*entities.Admin, error) {

	matchingAdmin, _ := af.adminRepo.FindByEmail(email)
	if matchingAdmin != nil {
		return nil, ErrDuplicatedAdminEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	return entities.NewAdmin(email, passwordHash), nil
}

func (af AccountFactory) createSupplier(email, password string) (*entities.Supplier, error) {

	matchingSupplier, _ := af.supplierRepo.FindByEmail(email)
	if matchingSupplier != nil {
		return nil, ErrDuplicatedSupplierEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	supplier := entities.NewSupplier(email, passwordHash)
	supplier.CreatedAt = time.Now()

	return supplier, nil
}

func (af AccountFactory) createPlanner(email, password string) (*entities.Planner, error) {

	matchingPlanner, _ := af.plannerRepo.FindByEmail(email)
	if matchingPlanner != nil {
		return nil, ErrDuplicatedPlannerEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	planner := entities.NewPlanner(email, passwordHash)
	planner.CreatedAt = time.Now()

	return planner, nil
}

func (af AccountFactory) createContractor(email, password string) (*entities.Contractor, error) {

	matchingContractor, _ := af.plannerRepo.FindByEmail(email)
	if matchingContractor != nil {
		return nil, ErrDuplicatedContractorEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	contractor := entities.NewContractor(email, passwordHash)
	contractor.CreatedAt = time.Now()

	return contractor, nil
}
