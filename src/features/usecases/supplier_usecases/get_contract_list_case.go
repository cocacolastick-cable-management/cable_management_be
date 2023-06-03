package supplier_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
)

type IGetContractListCase interface {
	Handle(accessToken string) ([]*responses.SupplierContractResponse, error)
}

type GetContractListCase struct {
	contractRepo       repositories.IContractRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
}

func NewGetContractListCase(contractRepo repositories.IContractRepository, makeSureAuthorized helpers.IMakeSureAuthorized) *GetContractListCase {
	return &GetContractListCase{contractRepo: contractRepo, makeSureAuthorized: makeSureAuthorized}
}

func (gcl GetContractListCase) Handle(accessToken string) ([]*responses.SupplierContractResponse, error) {

	claims, err := gcl.makeSureAuthorized.Handle(accessToken, constants.SupplierRole)
	if err != nil {
		return nil, err
	}

	contractList, _ := gcl.contractRepo.FindManyBySupplierId(claims.AccountId, []string{"WithDrawRequests"})

	response := make([]*responses.SupplierContractResponse, len(contractList))
	for i, contract := range contractList {
		contractRes, _ := helpers.MapSupplierContractResponse(contract)
		response[i] = contractRes
	}

	return response, nil
}
