package utils

type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Payload any    `json:"payload"`
	Errors  any    `json:"errors"`
}
