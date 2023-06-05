package errs

import "errors"

var (
	ErrDuplicatedEmail = errors.New("email is already in use")

	ErrAuthFailed     = errors.New("authenticate failed")
	ErrUnAuthorized   = errors.New("unauthorized")
	ErrDisableAccount = errors.New("account is disable")

	ErrInvalidJwtToken       = errors.New("invalid jwt-token")
	ErrInvalidRole           = errors.New("invalid role")
	ErrInvalidPasswordFormat = errors.New("invalid Password format")
	ErrInvalidCableAmount    = errors.New("invalid cable amount")
	ErrInvalidObjectType     = errors.New("invalid object type")

	ErrNotFoundContract   = errors.New("not found contract")
	ErrNotFoundContractor = errors.New("not found contractor")
	ErrNotFoundWithDraw   = errors.New("not found contractor")
	ErrNotFound           = errors.New("not found")
	ErrNotFoundUser       = errors.New("not found")

	ErrNotIncludeRelationship = errors.New("not include relationship")
)
