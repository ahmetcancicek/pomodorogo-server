package utils

import (
	"github.com/go-playground/validator/v10"
	"math/rand"
	"strings"
)

func PayloadValidator(model interface{}) error {
	validate := validator.New()
	validateError := validate.Struct(model)
	if validateError != nil {
		return validateError
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString generate a string of random characters of given length
func GenerateRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}
