package api

import (
	"github.com/SohamKanji/simple-bank-project/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.ValidCurrency(currency)
	}
	return false
}
