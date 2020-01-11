package util

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Errors represents the API errors
type Errors struct {
	Message string              `json:"message"`
	Status  int                 `json:"status"`
	Errors  []map[string]string `json:"errors"`
}

// Error represents the simple API error
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// BuildErrorsFromValidationErrors build error from the ValidatorErrors returned by the validator
func BuildErrorsFromValidationErrors(err error, trans *ut.Translator) []map[string]string {
	errors := make([]map[string]string, len(err.(validator.ValidationErrors)))

	for index, err := range err.(validator.ValidationErrors) {
		errors[index] = map[string]string{
			"field":   strings.ToLower(err.Field()),
			"message": err.Translate(*trans),
		}
	}

	return errors
}
