package services

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/domain/errs"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"time"
)

type INotifFactory interface {
	CreateNotifList(senderId uuid.UUID, action, objectType string, objectId uuid.UUID) ([]*entities.Notification, error)
	createNotifListCaseWithDraw(senderId, withDrawId uuid.UUID, action string) ([]*entities.Notification, error)
}

type NotifFactory struct {
	userRepo     repositories.IUserRepository
	withDrawRepo repositories.IWithDrawRequestRepository
}

func NewNotifFactory(userRepo repositories.IUserRepository, withDrawRepo repositories.IWithDrawRequestRepository) *NotifFactory {
	return &NotifFactory{userRepo: userRepo, withDrawRepo: withDrawRepo}
}

func (n NotifFactory) CreateNotifList(senderId uuid.UUID, action, objectType string, objectId uuid.UUID) ([]*entities.Notification, error) {

	_, err := n.userRepo.FindById(senderId)
	if err != nil {
		return nil, errs.ErrNotFoundUser
	}

	switch objectType {
	case constants.WithDrawReqObjectType:
		return n.createNotifListCaseWithDraw(senderId, objectId, action)
	default:
		return nil, errs.ErrInvalidObjectType
	}
}

func (n NotifFactory) createNotifListCaseWithDraw(senderId, withDrawId uuid.UUID, action string) ([]*entities.Notification, error) {

	receiverIdList := make([]uuid.UUID, 0)

	matchWithDraw, _ := n.withDrawRepo.FindById(withDrawId, []string{"Contract"})
	if matchWithDraw == nil {
		return nil, errs.ErrNotFoundWithDraw
	}

	receiverIdList = append(receiverIdList, matchWithDraw.ContractorId, matchWithDraw.Contract.SupplierId)

	notReceiverIdList := []uuid.UUID{senderId, matchWithDraw.ContractorId, matchWithDraw.Contract.SupplierId}

	receiverList, _ := n.userRepo.FindManyActiveByRoles([]string{constants.PlannerRole}, nil)

	var notifList []*entities.Notification
	for _, receiver := range receiverList {
		if slices.Contains(notReceiverIdList, receiver.Id) {
			continue
		}
		notifList = append(notifList, entities.NewNotification(action, false, constants.WithDrawReqObjectType, withDrawId, time.Now(), senderId, receiver.Id))
	}

	return notifList, nil
}
