package util

import (
	"github.com/go-playground/validator/v10"
)

func PayloadValidator(model interface{}) error {
	validate := validator.New()
	validateError := validate.Struct(model)
	if validateError != nil {
		return validateError
	}
	return nil
}
