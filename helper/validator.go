package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(v any) *entity.ErrorResponseWithErrors {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field: %s, Error: %s", strings.ToLower(err.Field()), err.Tag()))
		}
		return &entity.ErrorResponseWithErrors{
			Error:   true,
			Code:    http.StatusBadRequest,
			Message: "Bad Request with Payload",
			Errors:  errors,
		}
	}
	return nil
}
