package utils

import (
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/pt_BR"

	ut "github.com/go-playground/universal-translator"
	br_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func ValidateStruct(entity interface{}) []*ErrorResponse {
	pt_BR := pt_BR.New()
	uni = ut.New(pt_BR, pt_BR)
	trans, _ := uni.GetTranslator("pt_BR")
	validate = validator.New()

	br_translations.RegisterDefaultTranslations(validate, trans)

	var errors []*ErrorResponse
	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Translate(trans)
			errors = append(errors, &element)
		}
	}
	return errors
}
