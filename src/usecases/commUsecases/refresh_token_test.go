package commUsecases

import (
	"github.com/cable_management/cable_management_be/src/services"
	"reflect"
	"testing"
)

func TestRefreshToken_Handle(t *testing.T) {
	type fields struct {
		tokenService services.IAuthTokenService
	}
	type args struct {
		refreshToken string
	}

	dependencies := fields{
		tokenService: services.NewAuthTokenService(),
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
				refreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc1MTIzNTEsIlJvbGUiOiJhZG1pbiIsIlR5cGUiOiJyZWZyZXNoLXRva2VuIiwiQWNjb3VudElkIjoiNTE2Njg5MjgtMWIyOC00NDQyLTlkNTctZmVjMzdmNmY2Y2ZkIn0.LjaftAOV1ZjU7VT6PM-MRgGGGFZecno06zk_qiFRghI",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := RefreshToken{
				tokenService: tt.fields.tokenService,
			}
			got, err := rt.Handle(tt.args.refreshToken)
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
