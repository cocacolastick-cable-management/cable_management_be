package services

import (
	"github.com/cable_management/cable_management_be/src/infra/database"
	"github.com/cable_management/cable_management_be/src/infra/database/repositories"
	"reflect"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	type fields struct {
		passwordHashService IPasswordHashService
		tokenService        IAuthTokenService
		adminRepo           repositories.IAdminRepository
		plannerRepo         repositories.IPlannerRepository
		supplierRepo        repositories.ISupplierRepository
		contractorRepo      repositories.IContractorRepository
	}
	type args struct {
		role     string
		email    string
		password string
	}

	dependencies := fields{
		passwordHashService: NewPasswordHashService(),
		tokenService:        NewAuthTokenService(),
		adminRepo:           repositories.NewAdminRepository(database.DB),
		plannerRepo:         repositories.NewPlannerRepository(database.DB),
		supplierRepo:        repositories.NewSupplierRepository(database.DB),
		contractorRepo:      repositories.NewContractorRepository(database.DB),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthData
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				role:     AdminRole,
				email:    "vupham@gmail.com",
				password: "123456",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := AuthService{
				passwordHashService: tt.fields.passwordHashService,
				tokenService:        tt.fields.tokenService,
				adminRepo:           tt.fields.adminRepo,
				plannerRepo:         tt.fields.plannerRepo,
				supplierRepo:        tt.fields.supplierRepo,
				contractorRepo:      tt.fields.contractorRepo,
			}
			got, err := as.Authenticate(tt.args.role, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
