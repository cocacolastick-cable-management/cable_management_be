package repositories

import (
	"github.com/cable_management/cable_management_be/src_test/entities"
	"github.com/cable_management/cable_management_be/src_test/infra/database"
	"gorm.io/gorm"
	"testing"
)

func TestAdminRepository_Insert(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		entity *entities.Admin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "case-1",
			fields: fields{db: database.DB},
			args: args{
				entity: entities.NewAdmin("vupham@gmail.com", "123456"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := AdminRepository{
				db: tt.fields.db,
			}
			if err := ar.Insert(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func TestNewAdminRepository(t *testing.T) {
//	type args struct {
//		db *gorm.DB
//	}
//	tests := []struct {
//		name string
//		args args
//		want *AdminRepository
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewAdminRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewAdminRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestAdminRepository_FindByEmail(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		email string
	}

	dependencies := fields{
		db: database.DB,
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
				email: "vupham@gmail.com",
			},
			wantErr: false,
		},
		{
			name:   "case-2",
			fields: dependencies,
			args: args{
				email: "vupham2@gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := AdminRepository{
				db: tt.fields.db,
			}
			_, err := ar.FindByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FindByEmail() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
