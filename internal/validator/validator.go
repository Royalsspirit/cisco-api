package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Val represents the Valdiator
type Val struct {
	Val   *validator.Validate
	Trans ut.Translator
}

// NewValidator create a new Validator
func NewValidator() *Val {
	eng := en.New()
	uni := ut.New(eng, eng)

	trans, _ := uni.GetTranslator("en")

	v := validator.New()

	err := en_translations.RegisterDefaultTranslations(v, trans)

	if err != nil {
		panic(err)
	}

	return &Val{Val: v, Trans: trans}
}
