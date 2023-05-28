package validations

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/go-playground/validator/v10"
)

type ValidateCreateWithDrawRequestDependency struct {
	contractRepo repositories.IContractRepository
	userRepo     repositories.IUserRepository
}

func NewValidateCreateWithDrawRequestDependency(contractRepo repositories.IContractRepository, userRepo repositories.IUserRepository) *ValidateCreateWithDrawRequestDependency {
	return &ValidateCreateWithDrawRequestDependency{contractRepo: contractRepo, userRepo: userRepo}
}

func ValidateCreateWithDrawRequest(dependency *ValidateCreateWithDrawRequestDependency) func(sl validator.StructLevel) {

	return func(sl validator.StructLevel) {

		request := sl.Current().Interface().(requests.CreateWithDrawRequest)

		matchingContract, _ := dependency.contractRepo.FindById(request.ContractId, []string{"WithDrawRequests"})
		if matchingContract == nil {
			sl.ReportError(request.ContractId, "ContractId", "ContractId", "not-found", "not found contract")
		}

		cableStock, _ := matchingContract.CalCableStock()
		if request.CableAmount <= 0 || int(request.CableAmount) > cableStock {
			sl.ReportError(request.CableAmount, "CableAmount", "CableAmount", "invalid", "invalid cable amount")
		}

		matchingContractor, _ := dependency.userRepo.FindById(request.ContractorId)
		if matchingContract == nil || matchingContractor.Role != constants.ContractorRole {
			sl.ReportError(request.ContractorId, "ContractorId", "ContractorId", "not-found", "not found contractor")
		}
	}
}
