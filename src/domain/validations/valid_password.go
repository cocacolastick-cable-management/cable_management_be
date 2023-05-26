package validations

import "regexp"

func ValidatePassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasEnoughLength := len(password) >= 8

	return hasNumber && hasUppercase && hasEnoughLength
}
