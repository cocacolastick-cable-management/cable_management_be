package common_usecases

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/cable_management/cable_management_be/src/domain/services"
	"github.com/cable_management/cable_management_be/src/features/dtos/requests"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type IUpdateWithDrawStatusCase interface {
	Handle(accessToken string, withDrawId uuid.UUID, request requests.UpdateWithDrawStatusRequest) (*responses.WithDrawResponse, error)
}

type UpdateWithDrawStatusCase struct {
	makeSureAuthorized helpers.IMakeSureAuthorized
	validator          *validator.Validate
	withDrawRepo       repositories.IWithDrawRequestRepository
	historyRepo        repositories.IWithDrawRequestHistoryRepository
	notifRepo          repositories.INotificationRepository
	notifFac           services.INotifFactory
}

func NewUpdateWithDrawStatusCase(makeSureAuthorized helpers.IMakeSureAuthorized, validator *validator.Validate, withDrawRepo repositories.IWithDrawRequestRepository, historyRepo repositories.IWithDrawRequestHistoryRepository, notifRepo repositories.INotificationRepository, notifFac services.INotifFactory) *UpdateWithDrawStatusCase {
	return &UpdateWithDrawStatusCase{makeSureAuthorized: makeSureAuthorized, validator: validator, withDrawRepo: withDrawRepo, historyRepo: historyRepo, notifRepo: notifRepo, notifFac: notifFac}
}

func (uws UpdateWithDrawStatusCase) Handle(accessToken string, withDrawId uuid.UUID, request requests.UpdateWithDrawStatusRequest) (*responses.WithDrawResponse, error) {

	claims, err := uws.makeSureAuthorized.Handle(accessToken, constants.SupplierRole, constants.PlannerRole, constants.ContractorRole)
	if err != nil {
		return nil, err
	}

	matchWithDraw, err := uws.withDrawRepo.FindById(withDrawId, []string{"Histories", "Histories.Creator", "Contract", "Contract.Supplier", "Contractor"})
	if err != nil {
		return nil, errs.ErrNotFound
	}

	// TODO validate request

	action := ""
	switch {
	case request.NewStatus == constants.WD_CanceledStatus && claims.Role == constants.PlannerRole && matchWithDraw.Status == constants.WD_NewStatus:
		action = constants.WD_CancelAction
		break
	case request.NewStatus == constants.WD_ReadyStatus && claims.Role == constants.SupplierRole && matchWithDraw.Status == constants.WD_NewStatus:
		action = constants.WD_UpdateAction
		break
	case request.NewStatus == constants.WD_CollectedStatus && claims.Role == constants.ContractorRole && matchWithDraw.Status == constants.WD_ReadyStatus:
		action = constants.WD_UpdateAction
		break
	default:
		response, _ := helpers.MapWithDrawToResponse(matchWithDraw)
		return response, errs.ErrUnAuthorized
	}

	matchWithDraw.Status = request.NewStatus
	history := entities.NewWithDrawRequestHistory(action, time.Now(), request.NewStatus, claims.AccountId, matchWithDraw.Id)
	matchWithDraw.Histories = append(matchWithDraw.Histories, history)

	_ = uws.withDrawRepo.Save(matchWithDraw)

	go func() {
		notifList, _ := uws.notifFac.CreateNotifList(claims.AccountId, action, constants.WithDrawReqObjectType, matchWithDraw.Id)
		_ = uws.notifRepo.InsertMany(notifList)
	}()

	// TODO send email

	matchWithDraw, _ = uws.withDrawRepo.FindById(withDrawId, []string{"Histories", "Histories.Creator", "Contract", "Contract.Supplier", "Contractor"})

	response, _ := helpers.MapWithDrawToResponse(matchWithDraw)
	return response, nil
}
