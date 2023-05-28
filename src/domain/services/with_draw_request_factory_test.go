package services

import (
	"github.com/cable_management/cable_management_be/src/domain/contracts/db/repositories"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db"
	repositories2 "github.com/cable_management/cable_management_be/src/infra/db/repositories"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestWithDrawRequestFactory_CreateRequest(t *testing.T) {
	type fields struct {
		contractRepo repositories.IContractRepository
		userRepo     repositories.IUserRepository
	}
	type args struct {
		cableAmount  uint
		contractId   uuid.UUID
		contractorId uuid.UUID
	}

	contractId, _ := uuid.Parse("065ca3ef-b5ed-44a7-a4c0-9177076b9647")
	contractorId, _ := uuid.Parse("5cd7f608-1ef4-4cbc-9d38-9a6bd7c0797a")

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.WithDrawRequest
		wantErr bool
	}{
		{
			name: "case-1",
			fields: fields{
				contractRepo: repositories2.NewContractRepository(db.DB),
				userRepo:     repositories2.NewUserRepository(db.DB),
			},
			args: args{
				cableAmount:  10,
				contractId:   contractId,
				contractorId: contractorId,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wdf := WithDrawRequestFactory{
				contractRepo: tt.fields.contractRepo,
				userRepo:     tt.fields.userRepo,
			}
			got, err := wdf.CreateRequest(tt.args.cableAmount, tt.args.contractId, tt.args.contractorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
