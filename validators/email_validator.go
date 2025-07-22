package validators

import (
	"fmt"

	"github.com/aaronsisler/services.email/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateEmail(email models.Email) ([]ValidationError, error) {
	err := validate.Struct(email)
	if err == nil {
		return nil, nil
	}

	var errs []ValidationError
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			errs = append(errs, ValidationError{
				Field: fe.Namespace(),
				Error: fmt.Sprintf("validation for '%s' failed on the '%s' tag", fe.Field(), fe.Tag()),
			})
		}
		return errs, nil
	}

	return nil, err
}
