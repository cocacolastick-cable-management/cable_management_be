package planner_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
	"time"
)

type ICreateWithDrawCase interface {
	Handle(accessToken string, request requests.CreateWithDrawRequest) (*responses.WithDrawResponse, error)
}

type CreateWithDrawCase struct {
	withDrawFac        services.IWithDrawRequestFactory
	withDrawRepo       repositories.IWithDrawRequestRepository
	historyRepo        repositories.IWithDrawRequestHistoryRepository
	contractRepo       repositories.IContractRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
	validator          *validator.Validate
}

func NewCreateWithDrawCase(withDrawFac services.IWithDrawRequestFactory, withDrawRepo repositories.IWithDrawRequestRepository, historyRepo repositories.IWithDrawRequestHistoryRepository, contractRepo repositories.IContractRepository, makeSureAuthorized helpers.IMakeSureAuthorized, validator *validator.Validate) *CreateWithDrawCase {
	return &CreateWithDrawCase{withDrawFac: withDrawFac, withDrawRepo: withDrawRepo, historyRepo: historyRepo, contractRepo: contractRepo, makeSureAuthorized: makeSureAuthorized, validator: validator}
}

func (cwd CreateWithDrawCase) Handle(accessToken string, request requests.CreateWithDrawRequest) (*responses.WithDrawResponse, error) {

	var err error

	var claims *services.AuthTokenClaims
	claims, err = cwd.makeSureAuthorized.Handle(accessToken, constants.PlannerRole)
	if err != nil {
		return nil, err
	}

	err = cwd.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	newWithDraw, _ := cwd.withDrawFac.CreateRequest(request.CableAmount, request.ContractId, request.ContractorId)
	newHistory := entities.NewWithDrawRequestHistory(constants.WD_CreateAction, time.Now(), newWithDraw.Status, claims.AccountId, newWithDraw.Id)

	_ = cwd.withDrawRepo.Insert(newWithDraw)
	_ = cwd.historyRepo.Insert(newHistory)

	newWithDraw, _ = cwd.withDrawRepo.FindById(newWithDraw.Id, []string{"Histories", "Contract", "Contract.Supplier", "Contractor"})

	return helpers.MapWithDrawToResponse(newWithDraw)
}
