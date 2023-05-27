package planner_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
	"time"
)

type IGetContractListCase interface {
	Handle(accessToken string, request requests.PaginationRequest) ([]*responses.ContractResponse, error)
}

type GetContractListCase struct {
	contractRepo       repositories.IContractRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
	validator          *validator.Validate
}

func NewGetContractListCase(contractRepo repositories.IContractRepository, makeSureAuthorized helpers.IMakeSureAuthorized, validator *validator.Validate) *GetContractListCase {
	return &GetContractListCase{contractRepo: contractRepo, makeSureAuthorized: makeSureAuthorized, validator: validator}
}

func (gcl GetContractListCase) Handle(accessToken string, request requests.PaginationRequest) ([]*responses.ContractResponse, error) {

	var err error

	_, err = gcl.makeSureAuthorized.Handle(accessToken, constants.PlannerRole)
	if err != nil {
		return nil, err
	}

	err = gcl.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	contractList, _ := gcl.contractRepo.GetList("Supplier", request.Page, request.Size, request.OrderBy, request.LastTimestamp)

	response := make([]*responses.ContractResponse, len(contractList))
	for i, contract := range contractList {
		response[i] = &responses.ContractResponse{
			Id:           contract.Id,
			SupplierId:   contract.SupplierId,
			SupplierName: contract.Supplier.DisplayName,
			StartDay:     time.Time{},
			EndDay:       time.Time{},
			CreatedAt:    time.Time{},
		}
	}

	return response, nil
}
