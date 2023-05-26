package services

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestAuthTokenService_GenerateAuthData(t *testing.T) {
	type args struct {
		role      string
		accountId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthData
		wantErr bool
	}{
		{
			name: "case-1",
			args: args{
				role:      AdminRole,
				accountId: uuid.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := AuthTokenService{}
			got, err := ats.GenerateAuthData(tt.args.role, tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAuthData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthTokenService_IsAccessTokenValid(t *testing.T) {
	type args struct {
		accessToken string
		claims      *AuthTokenClaims
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ4NjMxMDcsIlJvbGUiOiJhZG1pbiIsIlR5cGUiOiJhY2Nlc3MtdG9rZW4iLCJBY2NvdW50SWQiOiI0ZTIwMGUwYy0xZDNiLTQ2ZmMtYjY1OS00MDg1OTljMTcwODEifQ.ZQUe8Qwtrl_M0xzJfzNM61UcgmTIPnNKqKQqKtSv6yU",
				claims:      nil,
			},
			want: true,
		},
		{
			name: "case-2",
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc0NTQ5ODcsIlJvbGUiOiJhZG1pbiIsIlR5cGUiOiJyZWZyZXNoLXRva2VuIiwiQWNjb3VudElkIjoiNGUyMDBlMGMtMWQzYi00NmZjLWI2NTktNDA4NTk5YzE3MDgxIn0.C1jgrX7Hjpp7reanM3SDIqJOXJngs2FhjhrdnC_ZpVA",
				claims:      nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := AuthTokenService{}
			if got := ats.IsAccessTokenValid(tt.args.accessToken, tt.args.claims); got != tt.want {
				t.Errorf("IsAccessTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthTokenService_IsRefreshTokenValid(t *testing.T) {
	type args struct {
		refreshToken string
		claims       *AuthTokenClaims
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := AuthTokenService{}
			if got := ats.IsRefreshTokenValid(tt.args.refreshToken, tt.args.claims); got != tt.want {
				t.Errorf("IsRefreshTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthTokenService_ParseToClaims(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthTokenClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := AuthTokenService{}
			got, err := ats.ParseToClaims(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthTokenService(t *testing.T) {
	tests := []struct {
		name string
		want *AuthTokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthTokenService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthTokenService() = %v, want %v", got, tt.want)
			}
		})
	}
}
