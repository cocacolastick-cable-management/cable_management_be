package requests

type CreateUserRequest struct {
	Role        string
	DisplayName string
	Email       string
	Password    string
}
