package tools

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notEmptyString", notEmptyStringValidator)
	}
}

func notEmptyStringValidator(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	trimmedString := strings.TrimSpace(field)
	return len(trimmedString) > 0
}
