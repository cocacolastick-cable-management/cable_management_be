package helpers

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
)

func MapHistoryResponse(history *entities.WithDrawRequestHistory) (*responses.HistoryResponse, error) {

	if history.Creator == nil {
		return nil, errs.ErrNotIncludeRelationship
	}

	return &responses.HistoryResponse{
		Id:           history.Id,
		RequestId:    history.RequestId,
		CreatorId:    history.CreatorId,
		CreatorName:  history.Creator.DisplayName,
		CreatorEmail: history.Creator.Email,
		CreatorRole:  history.Creator.Role,
		CreatedAt:    history.CreatedAt.UTC(),
		Status:       history.Status,
		Action:       history.Action,
	}, nil
}
