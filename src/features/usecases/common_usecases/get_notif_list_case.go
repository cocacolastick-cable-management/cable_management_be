package common_usecases

import (
	"fmt"
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
	"github.com/cable_management/cable_management_be/src/features/helpers"
)

type IGetNotifListCase interface {
	Handle(accessToken string) ([]*responses.NotifResponse, error)
}

type GetNotifListCase struct {
	notifRepo          repositories.INotificationRepository
	makeSureAuthorized helpers.IMakeSureAuthorized
}

func NewGetNotifListCase(notifRepo repositories.INotificationRepository, makeSureAuthorized helpers.IMakeSureAuthorized) *GetNotifListCase {
	return &GetNotifListCase{notifRepo: notifRepo, makeSureAuthorized: makeSureAuthorized}
}

func (g GetNotifListCase) Handle(accessToken string) ([]*responses.NotifResponse, error) {

	claims, err := g.makeSureAuthorized.Handle(accessToken, constants.ContractorRole, constants.PlannerRole, constants.SupplierRole)
	if err != nil {
		return nil, err
	}

	notifList, _ := g.notifRepo.FindManyByReceiverId(claims.AccountId, nil)
	fmt.Print(notifList)

	//TODO implement me
	panic("implement me")
}
