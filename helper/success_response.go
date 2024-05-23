package helper

import "github.com/dimassfeb-09/smart-library-be/entity"

func SuccessResponseWithData(code int, message string, data any) *entity.ResponseWebWithData {
	return &entity.ResponseWebWithData{
		Error:   false,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponseWithoutData(code int, message string) *entity.ResponseWebWithoutData {
	return &entity.ResponseWebWithoutData{
		Error:   false,
		Code:    code,
		Message: message,
	}
}
