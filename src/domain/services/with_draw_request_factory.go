package services

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/google/uuid"
	"time"
)

type IWithDrawRequestFactory interface {
	CreateRequest(cableAmount uint, contractId uuid.UUID, contractorId uuid.UUID) (*entities.WithDrawRequest, error)
}

type WithDrawRequestFactory struct {
	contractRepo repositories.IContractRepository
	userRepo     repositories.IUserRepository
}

func (wdf WithDrawRequestFactory) CreateRequest(cableAmount uint, contractId uuid.UUID, contractorId uuid.UUID) (*entities.WithDrawRequest, error) {

	matchingContract, _ := wdf.contractRepo.FindById(contractId, []string{"WithDrawRequests"})
	if matchingContract == nil {
		return nil, errs.ErrNotFoundContract
	}

	cableStock, _ := matchingContract.CalCableStock()
	if cableAmount <= 0 || int(cableAmount) > cableStock {
		return nil, errs.ErrInvalidCableAmount
	}

	matchingContractor, _ := wdf.userRepo.FindById(contractorId)
	if matchingContract == nil || matchingContractor.Role != constants.ContractorRole {
		return nil, errs.ErrNotFoundContractor
	}

	return entities.NewWithDrawRequest(constants.WD_NewStatus, cableAmount, time.Now(), contractId, contractorId), nil
}
