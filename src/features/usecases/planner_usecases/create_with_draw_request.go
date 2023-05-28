package planner_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
)

type ICreateWithDrawCase interface {
	Handle(accessToken string, request requests.CreateWithDrawRequest) (*responses.WithDrawResponse, error)
}

type CreateWithDrawCase struct {
	withDrawFac        services.IWithDrawRequestFactory
	withDrawRepo       repositories.IWithDrawRequestRepository
	contractRepo       repositories.IContractRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
	validator          *validator.Validate
}

func NewCreateWithDrawCase(withDrawFac services.IWithDrawRequestFactory, withDrawRepo repositories.IWithDrawRequestRepository, contractRepo repositories.IContractRepository, makeSureAuthorized helpers.IMakeSureAuthorized, validator *validator.Validate) *CreateWithDrawCase {
	return &CreateWithDrawCase{withDrawFac: withDrawFac, withDrawRepo: withDrawRepo, contractRepo: contractRepo, makeSureAuthorized: makeSureAuthorized, validator: validator}
}

func (cwd CreateWithDrawCase) Handle(accessToken string, request requests.CreateWithDrawRequest) (*responses.WithDrawResponse, error) {

	var err error

	_, err = cwd.makeSureAuthorized.Handle(accessToken, constants.PlannerRole)
	if err != nil {
		return nil, err
	}

	err = cwd.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	newWithDraw, _ := cwd.withDrawFac.CreateRequest(request.CableAmount, request.ContractId, request.ContractorId)
	_ = cwd.withDrawRepo.Insert(newWithDraw)

	return &responses.WithDrawResponse{
		Id:           newWithDraw.Id,
		SupplierId:   newWithDraw.Contract.SupplierId,
		ContractorId: newWithDraw.ContractorId,
		ContractId:   newWithDraw.ContractId,
		CableAmount:  newWithDraw.CableAmount,
		Status:       newWithDraw.Status,
		CreatedAt:    newWithDraw.CreatedAt,
	}, nil
}
