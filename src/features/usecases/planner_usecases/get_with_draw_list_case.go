package planner_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
)

type IGetWithDrawListCase interface {
	Handle(accessToken string) ([]*responses.WithDrawResponse, error)
}

type GetWithDrawListCase struct {
	makeSureAuthorized helpers.IMakeSureAuthorized
	withDrawRepo       repositories.IWithDrawRequestRepository
}

func NewGetWithDrawListCase(makeSureAuthorized helpers.IMakeSureAuthorized, withDrawRepo repositories.IWithDrawRequestRepository) *GetWithDrawListCase {
	return &GetWithDrawListCase{makeSureAuthorized: makeSureAuthorized, withDrawRepo: withDrawRepo}
}

func (gwd GetWithDrawListCase) Handle(accessToken string) ([]*responses.WithDrawResponse, error) {

	var err error

	_, err = gwd.makeSureAuthorized.Handle(accessToken, constants.PlannerRole)
	if err != nil {
		return nil, err
	}

	withDrawList, _ := gwd.withDrawRepo.GetActiveList([]string{"Histories", "Histories.Creator", "Contract", "Contract.Supplier", "Contractor"})

	response := make([]*responses.WithDrawResponse, len(withDrawList))
	for i, withDraw := range withDrawList {
		withDrawRes, _ := helpers.MapWithDrawToResponse(withDraw)
		response[i] = withDrawRes
	}

	return response, nil
}
