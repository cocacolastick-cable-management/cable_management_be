package helpers

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
)

func MapWithDrawToResponse(withDraw *entities.WithDrawRequest) (*responses.WithDrawResponse, error) {

	if withDraw.Contract == nil || withDraw.Contract.Supplier == nil || withDraw.Contractor == nil {
		return nil, errs.ErrNotIncludeRelationship
	}

	historyListRes := make([]*responses.HistoryResponse, len(withDraw.Histories))
	for i, history := range withDraw.Histories {
		historyListRes[i], _ = MapHistoryResponse(history)
	}

	return &responses.WithDrawResponse{
		Id:             withDraw.Id,
		SupplierId:     withDraw.Contract.SupplierId,
		SupplierName:   withDraw.Contract.Supplier.DisplayName,
		ContractorId:   withDraw.ContractorId,
		ContractorName: withDraw.Contractor.DisplayName,
		ContractId:     withDraw.ContractId,
		CableAmount:    withDraw.CableAmount,
		Status:         withDraw.Status,
		CreatedAt:      withDraw.CreatedAt.UTC(),
		Histories:      historyListRes,
	}, nil
}
