package helpers

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/features/dtos/responses"
)

func MapUserResponse(user *entities.User) (*responses.UserResponse, error) {
	return &responses.UserResponse{
		Id:          user.Id,
		DisplayName: user.DisplayName,
		Role:        user.Role,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt.UTC(),
	}, nil
}
