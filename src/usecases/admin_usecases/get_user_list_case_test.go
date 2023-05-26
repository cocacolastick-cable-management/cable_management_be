package admin_usecases

import (
	"github.com/cable_management/cable_management_be/infras/db"
	"github.com/cable_management/cable_management_be/infras/db/repositories"
	"github.com/cable_management/cable_management_be/infras/valider"
	"github.com/cable_management/cable_management_be/src/dtos/requests"
	"github.com/cable_management/cable_management_be/src/dtos/responses"
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

func TestGetUserListCase_Handle(t *testing.T) {
	type fields struct {
		tokenService services.IAuthTokenService
		validator    *validator.Validate
		userRepo     repositories.IUserRepository
	}
	type args struct {
		accessToken string
		request     requests.PaginationRequest
	}

	dependencies := fields{
		tokenService: services.NewAuthTokenService(),
		validator:    valider.Validator,
		userRepo:     repositories.NewUserRepository(db.DB),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []responses.UserResponse
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				accessToken: "123",
				request: requests.PaginationRequest{
					Page:          0,
					Size:          5,
					LastTimestamp: nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gul := GetUserListCase{
				tokenService: tt.fields.tokenService,
				validator:    tt.fields.validator,
				userRepo:     tt.fields.userRepo,
			}
			got, err := gul.Handle(tt.args.accessToken, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
