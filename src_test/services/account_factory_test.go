package services

import (
	"github.com/cable_management/cable_management_be/src_test/entities"
	"github.com/cable_management/cable_management_be/src_test/infra/database"
	"github.com/cable_management/cable_management_be/src_test/infra/database/repositories"
	"reflect"
	"testing"
)

func TestAccountFactory_CreateNewAccount(t *testing.T) {

	type fields struct {
		passwordHashService IPasswordHashService
		adminRepo           repositories.IAdminRepository
	}
	type args struct {
		role     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.IAbstractAccount
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := AccountFactory{
				passwordHashService: tt.fields.passwordHashService,
				adminRepo:           tt.fields.adminRepo,
			}
			got, err := af.CreateNewAccount(tt.args.role, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNewAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestAccountFactory_createAdmin(t *testing.T) {
//
//	type fields struct {
//		passwordHashService IPasswordHashService
//		adminRepo           repositories.IAdminRepository
//	}
//	type args struct {
//		email    string
//		password string
//	}
//
//	dependencies := fields{
//		passwordHashService: NewPasswordHashService(),
//		adminRepo:           repositories.NewAdminRepository(database.DB),
//	}
//
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		{
//			name:   "case-1",
//			fields: dependencies,
//			args: args{
//				email:    "vupham@gmail.com",
//				password: "123456",
//			},
//			wantErr: true,
//		},
//		{
//			name:   "case-2",
//			fields: dependencies,
//			args: args{
//				email:    "vupham2fsdsf@gmail.com",
//				password: "123456",
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			af := AccountFactory{
//				passwordHashService: tt.fields.passwordHashService,
//				adminRepo:           tt.fields.adminRepo,
//			}
//			_, err := af.createAdmin(tt.args.email, tt.args.password)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("createAdmin() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			//if !reflect.DeepEqual(got, tt.want) {
//			//	t.Errorf("createAdmin() got = %v, want %v", got, tt.want)
//			//}
//		})
//	}
//}

func TestAccountFactory_createSupplier(t *testing.T) {
	type fields struct {
		passwordHashService IPasswordHashService
		adminRepo           repositories.IAdminRepository
		supplierRepo        repositories.ISupplierRepository
	}
	type args struct {
		email    string
		password string
	}

	dependencies := fields{
		passwordHashService: NewPasswordHashService(),
		adminRepo:           repositories.NewAdminRepository(database.DB),
		supplierRepo:        repositories.NewSupplierRepository(database.DB),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: dependencies,
			args: args{
				email:    "vupham@gmail.com",
				password: "123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := AccountFactory{
				passwordHashService: tt.fields.passwordHashService,
				adminRepo:           tt.fields.adminRepo,
				supplierRepo:        tt.fields.supplierRepo,
			}
			_, err := af.createSupplier(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSupplier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
