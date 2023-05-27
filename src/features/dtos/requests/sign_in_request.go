package requests

type SignInRequest struct {
	Role     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
