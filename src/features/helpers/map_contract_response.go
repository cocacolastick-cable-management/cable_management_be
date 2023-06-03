package helpers

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
)

func MapContractResponse(contract *entities.Contract) (*responses.ContractResponse, error) {

	stock, err := contract.CalCableStock()
	if err != nil {
		return nil, err
	}

	if contract.Supplier == nil {
		return nil, errs.ErrNotIncludeRelationship
	}

	response := &responses.ContractResponse{
		Id:            contract.Id,
		SupplierId:    contract.SupplierId,
		SupplierName:  contract.Supplier.DisplayName,
		CableAmount:   contract.CableAmount,
		Stock:         stock,
		UniqueName:    contract.UniqueName,
		SupplierEmail: contract.Supplier.Email,
		StartDay:      contract.StartDay.UTC(),
		EndDay:        contract.EndDay.UTC(),
		CreatedAt:     contract.CreatedAt.UTC(),
	}

	return response, nil
}

func MapSupplierContractResponse(contract *entities.Contract) (*responses.SupplierContractResponse, error) {

	stock, err := contract.CalCableStock()
	if err != nil {
		return nil, err
	}

	response := &responses.SupplierContractResponse{
		Id:          contract.Id,
		UniqueName:  contract.UniqueName,
		CableAmount: contract.CableAmount,
		Stock:       stock,
		StartDay:    contract.StartDay.UTC(),
		EndDay:      contract.EndDay.UTC(),
		CreatedAt:   contract.CreatedAt.UTC(),
	}

	return response, nil
}
