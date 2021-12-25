package api

import (
	"order-demo/util"

	"github.com/go-playground/validator/v10"
)

var validEmail validator.Func = func(fl validator.FieldLevel) bool {
	if email, ok := fl.Field().Interface().(string); ok {
		return util.EmailRegex().MatchString(email)
	}
	return false
}
