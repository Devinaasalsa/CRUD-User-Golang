package validators

import (
	"fmt"
	"strings"
	// "myapp-me/models"
	// "github.com/go-playground/validator/v10"
)

func GetValidationErrorMessage(tag, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Field %s is required", field)
	case "unique":
		return fmt.Sprintf("Duplicate entry 'value' for key %s", field)
	case "email":
		return fmt.Sprintf("Invalid email format for %s", field)
	case "min":
		return fmt.Sprintf("Field %s must have a minimum length of %s", field, extractParam(tag))
	case "max":
		return fmt.Sprintf("Field %s must have a maximum length of %s", field, extractParam(tag))
	case "gte":
		return fmt.Sprintf("Field %s must be greater than or equal to %s", field, extractParam(tag))
	case "lte":
		return fmt.Sprintf("Field %s must be less than or equal to %s", field, extractParam(tag))
	case "eqfield":
		return fmt.Sprintf("Field %s must be equal to %s", field, extractParam(tag))
	default:
		return fmt.Sprintf("Invalid value for %s", field)
	}
}

func extractParam(tag string) string {
	parts := strings.SplitN(tag, "=", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
