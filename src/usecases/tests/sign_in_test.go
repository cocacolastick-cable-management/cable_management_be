package tests

import (
	"github.com/cable_management/cable_management_be/src/infra/database"
	"github.com/cable_management/cable_management_be/src/infra/database/repositories"
	"github.com/cable_management/cable_management_be/src/infra/valider"
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/cable_management/cable_management_be/src/usecases/commUsecases"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

func TestSignIn_Handle(t *testing.T) {
	type fields struct {
		authService services.IAuthService
		validator   *validator.Validate
	}
	type args struct {
		signInRequest *commUsecases.SignInRequestDto
	}

	dependencies := fields{
		authService: services.NewAuthService(
			services.NewPasswordHashService(),
			services.NewAuthTokenService(),
			repositories.NewAdminRepository(database.DB),
			repositories.NewPlannerRepository(database.DB),
			repositories.NewSupplierRepository(database.DB),
			repositories.NewContractorRepository(database.DB)),
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
				signInRequest: &commUsecases.SignInRequestDto{
					Role:     "admin",
					Email:    "vupham@gmail.com",
					Password: "123456",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := commUsecases.NewSignIn(tt.fields.authService, tt.fields.validator)
			got, err := si.Handle(tt.args.signInRequest)
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
