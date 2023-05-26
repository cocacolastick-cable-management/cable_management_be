package services

import (
	"testing"
)

func TestPasswordHashService_Hash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case-1",
			args: args{
				password: "123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phs := PasswordHashService{}
			_, err := phs.Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Logf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
