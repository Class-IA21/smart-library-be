package helper

import "github.com/dimassfeb-09/smart-library-be/entity"

func ErrorResponse(code int, message string) *entity.ErrorResponse {
	return &entity.ErrorResponse{
		Error:   true,
		Code:    code,
		Message: message,
	}
}
