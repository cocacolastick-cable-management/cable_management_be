package common_usecases

import (
	"github.com/cable_management/cable_management_be/src/constants"
	"github.com/cable_management/cable_management_be/src/infras/repositories"
	"github.com/cable_management/cable_management_be/src/instances/db"
	"github.com/cable_management/cable_management_be/src/instances/valider"
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

func TestSignInCase_Handle(t *testing.T) {
	type fields struct {
		authService services.IAuthService
		validator   *validator.Validate
	}
	type args struct {
		request *SignInRequest
	}

	dependencies := fields{
		authService: services.NewAuthService(
			repositories.NewUserRepository(db.DB),
			services.NewPasswordHashService(),
			services.NewAuthTokenService()),
		validator: valider.Validator,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *services.AuthData
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				request: &SignInRequest{
					Role:     constants.AdminRole,
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
			sic := SignInCase{
				authService: tt.fields.authService,
				validator:   tt.fields.validator,
			}
			got, err := sic.Handle(tt.args.request)
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
