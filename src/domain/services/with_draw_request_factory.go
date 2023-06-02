package services

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"time"
)

type IWithDrawRequestFactory interface {
	CreateRequest(cableAmount uint, ContractUniqueName string, contractorEmail string) (*entities.WithDrawRequest, error)
}

type WithDrawRequestFactory struct {
	contractRepo repositories.IContractRepository
	userRepo     repositories.IUserRepository
}

func NewWithDrawRequestFactory(contractRepo repositories.IContractRepository, userRepo repositories.IUserRepository) *WithDrawRequestFactory {
	return &WithDrawRequestFactory{contractRepo: contractRepo, userRepo: userRepo}
}

func (wdf WithDrawRequestFactory) CreateRequest(cableAmount uint, ContractUniqueName string, contractorEmail string) (*entities.WithDrawRequest, error) {

	matchingContract, _ := wdf.contractRepo.FindByUniqueName(ContractUniqueName, []string{"WithDrawRequests"})
	if matchingContract == nil {
		return nil, errs.ErrNotFoundContract
	}

	cableStock, _ := matchingContract.CalCableStock()
	if cableAmount <= 0 || int(cableAmount) > cableStock {
		return nil, errs.ErrInvalidCableAmount
	}

	matchingContractor, _ := wdf.userRepo.FindByEmail(contractorEmail)
	if matchingContractor == nil || matchingContractor.Role != constants.ContractorRole {
		return nil, errs.ErrNotFoundContractor
	}

	newWithDraw := entities.NewWithDrawRequest(constants.WD_NewStatus, cableAmount, time.Now(), matchingContract.Id, matchingContractor.Id)
	return newWithDraw, nil
}
