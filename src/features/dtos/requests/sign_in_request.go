package requests

type SignInRequest struct {
	Role     string `json:"role" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
