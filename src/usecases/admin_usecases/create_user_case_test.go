package admin_usecases

import (
	"github.com/cable_management/cable_management_be/src/constants"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
	"github.com/cable_management/cable_management_be/src/instances/db"
	"github.com/cable_management/cable_management_be/src/instances/valider"
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/cable_management/cable_management_be/src/usecases/_commons/requests"
	"github.com/cable_management/cable_management_be/src/usecases/_commons/responses"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

func TestCreateUserCase_Handle(t *testing.T) {
	type fields struct {
		tokenService services.IAuthTokenService
		userFac      services.IUserFactory
		userRepo     repositories.IUserRepository
		validator    *validator.Validate
	}
	type args struct {
		accessToken string
		request     *requests.CreateUserRequest
	}

	userRepo := repositories.NewUserRepository(db.DB)
	dependencies := fields{
		tokenService: services.NewAuthTokenService(),
		userFac:      services.NewUserFactory(userRepo, services.NewPasswordHashService()),
		userRepo:     userRepo,
		validator:    valider.Validator,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *responses.UserResponse
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwOTc3MzgsIlJvbGUiOiJhZG1pbiIsIlR5cGUiOiJhY2Nlc3MtdG9rZW4iLCJBY2NvdW50SWQiOiIwNjVjYTNlZi1iNWVkLTQ0YTctYTRjMC05MTc3MDc2Yjk2NDcifQ.1U9sYm4k5MBWun-CtExfsoe9J1WwjmN47MIA3l-f6TM",
				request: &requests.CreateUserRequest{
					Role:     constants.PlannerRole,
					Email:    "vupham@gmail.com",
					Password: "123456Vv",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cac := NewCreateUserCase(
				tt.fields.tokenService,
				tt.fields.userFac,
				tt.fields.userRepo,
				tt.fields.validator,
			)
			got, err := cac.Handle(tt.args.accessToken, tt.args.request)
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
