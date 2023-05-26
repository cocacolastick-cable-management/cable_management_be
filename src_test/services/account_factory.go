package services

import (
	"errors"
	"github.com/cable_management/cable_management_be/src_test/entities"
	"github.com/cable_management/cable_management_be/src_test/infra/database/repositories"
	"github.com/google/uuid"
	"time"
)

const (
	AdminRole      = "admin"
	PlannerRole    = "planner"
	SupplierRole   = "supplier"
	ContractorRole = "contractor"
)

var (
	ErrIncorrectRole = errors.New("incorrect role")
	//ErrDuplicatedAdminEmail      = errors.New("duplicated admin email")
	ErrDuplicatedSupplierEmail   = errors.New("duplicated supplier email")
	ErrDuplicatedPlannerEmail    = errors.New("duplicated planner email")
	ErrDuplicatedContractorEmail = errors.New("duplicated contractor email")
)

type IAccountFactory interface {
	CreateNewAccount(role, email, password string, creatorId uuid.UUID) (entities.IAbstractAccount, error)
	//createAdmin(email, password string) (*entities.Admin, error)
	createSupplier(email, password string, creatorId uuid.UUID) (*entities.Supplier, error)
	createPlanner(email, password string, creatorId uuid.UUID) (*entities.Planner, error)
	createContractor(email, password string, creatorId uuid.UUID) (*entities.Contractor, error)
}

type AccountFactory struct {
	passwordHashService IPasswordHashService
	adminRepo           repositories.IAdminRepository
	supplierRepo        repositories.ISupplierRepository
	plannerRepo         repositories.IPlannerRepository
	contractorRepo      repositories.IContractorRepository
}

func NewAccountFactory(
	passwordHashService IPasswordHashService,
	supplierRepo repositories.ISupplierRepository,
	plannerRepo repositories.IPlannerRepository,
	contractorRepo repositories.IContractorRepository) *AccountFactory {

	return &AccountFactory{
		passwordHashService: passwordHashService,
		supplierRepo:        supplierRepo,
		plannerRepo:         plannerRepo,
		contractorRepo:      contractorRepo}
}

func (af AccountFactory) CreateNewAccount(role, email, password string, creatorId uuid.UUID) (entities.IAbstractAccount, error) {

	switch role {
	case SupplierRole:
		return af.createSupplier(email, password, creatorId)
	case PlannerRole:
		return af.createPlanner(email, password, creatorId)
	case ContractorRole:
		return af.createContractor(email, password, creatorId)
	default:
		return nil, ErrIncorrectRole
	}
}

func (af AccountFactory) createSupplier(email, password string, creatorId uuid.UUID) (*entities.Supplier, error) {

	matchingSupplier, _ := af.supplierRepo.FindByEmail(email)
	if matchingSupplier != nil {
		return nil, ErrDuplicatedSupplierEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	supplier := entities.NewSupplier(email, passwordHash)
	supplier.CreatedAt = time.Now()
	supplier.CreatorId = creatorId

	return supplier, nil
}

func (af AccountFactory) createPlanner(email, password string, creatorId uuid.UUID) (*entities.Planner, error) {

	matchingPlanner, _ := af.plannerRepo.FindByEmail(email)
	if matchingPlanner != nil {
		return nil, ErrDuplicatedPlannerEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	planner := entities.NewPlanner(email, passwordHash)
	planner.CreatedAt = time.Now()
	planner.CreatorId = creatorId

	return planner, nil
}

func (af AccountFactory) createContractor(email, password string, creatorId uuid.UUID) (*entities.Contractor, error) {

	matchingContractor, _ := af.contractorRepo.FindByEmail(email)
	if matchingContractor != nil {
		return nil, ErrDuplicatedContractorEmail
	}

	passwordHash, _ := af.passwordHashService.Hash(password)

	contractor := entities.NewContractor(email, passwordHash)
	contractor.CreatedAt = time.Now()
	contractor.CreatorId = creatorId

	return contractor, nil
}
