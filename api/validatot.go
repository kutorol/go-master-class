package api

import (
	"backend-master-class/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if cur, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(cur)
	}
	return false
}
