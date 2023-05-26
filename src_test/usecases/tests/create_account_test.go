package tests

import (
	"github.com/cable_management/cable_management_be/src_test/infra/database"
	"github.com/cable_management/cable_management_be/src_test/infra/database/repositories"
	"github.com/cable_management/cable_management_be/src_test/infra/valider"
	"github.com/cable_management/cable_management_be/src_test/services"
	"github.com/cable_management/cable_management_be/src_test/usecases/adminUsecases"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

func TestCreateAccount_Handle(t *testing.T) {
	type fields struct {
		validator      *validator.Validate
		accountFac     services.IAccountFactory
		supplierRepo   repositories.ISupplierRepository
		plannerRepo    repositories.IPlannerRepository
		contractorRepo repositories.IContractorRepository
		tokenService   services.IAuthTokenService
	}
	type args struct {
		accessToken   string
		createRequest *adminUsecases.CreateAccountRequestDto
	}

	dependencies := fields{
		validator: valider.Validator,
		accountFac: services.NewAccountFactory(
			services.NewPasswordHashService(),
			repositories.NewSupplierRepository(database.DB),
			repositories.NewPlannerRepository(database.DB),
			repositories.NewContractorRepository(database.DB)),
		supplierRepo:   repositories.NewSupplierRepository(database.DB),
		plannerRepo:    repositories.NewPlannerRepository(database.DB),
		contractorRepo: repositories.NewContractorRepository(database.DB),
		tokenService:   services.NewAuthTokenService(),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *adminUsecases.CreateAccountResponseDto
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ5MzY5MDgsIlJvbGUiOiJhZG1pbiIsIlR5cGUiOiJhY2Nlc3MtdG9rZW4iLCJBY2NvdW50SWQiOiI1MTY2ODkyOC0xYjI4LTQ0NDItOWQ1Ny1mZWMzN2Y2ZjZjZmQifQ.n0EJ7ViA9F40T8cEAnjcd_u0Z7gq-v19adRHgyje38I",
				createRequest: &adminUsecases.CreateAccountRequestDto{
					Role:     services.ContractorRole,
					Email:    "vupham2@gmail.com",
					Password: "123456@Vv",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := adminUsecases.NewCreateAccount(
				tt.fields.validator,
				tt.fields.accountFac,
				tt.fields.supplierRepo,
				tt.fields.plannerRepo,
				tt.fields.contractorRepo,
				tt.fields.tokenService)
			got, err := ca.Handle(tt.args.accessToken, tt.args.createRequest)
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
