package middleware

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func PwdValidation(field validator.FieldLevel) bool {
	inputPwd := field.Field().String()
	hasLower := false
	hasUpper := false
	hasNumber := false

	for _, ch := range inputPwd {
		if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsNumber(ch) {
			hasNumber = true
		}
	}

	return hasLower && hasUpper && hasNumber
}
