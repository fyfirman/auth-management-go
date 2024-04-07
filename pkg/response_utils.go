package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func PrepareValidationErrors(err error) []map[string]string {
	var errors []map[string]string
	for _, err := range err.(validator.ValidationErrors) {
		var element = make(map[string]string)
		element["field"] = err.Field()
		element["message"] = fieldErrorMessage(err)
		errors = append(errors, element)
	}
	return errors
}

func fieldErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email address", e.Field())
	case "min":
		// e.Param() contains the parameter of the validation tag, for min it's the minimum length/number
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters", e.Field(), e.Param())
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s", e.Field(), e.Param())
	case "alphanum":
		return fmt.Sprintf("%s must contain alphanumeric characters only", e.Field())
	// Add more cases as needed based on your validation rules
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}
