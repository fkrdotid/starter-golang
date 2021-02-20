package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status string `json:"status"`
	Code int `json:"code"`
	Success bool `json:"success"`
}

func JsonResponse(message string, code int, status string, success bool, data interface{} ) Response{
	meta := Meta {
		Message: message,
		Code: code,
		Success: success,
		Status: status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}


func FormatValidationError(err error) []string {
	var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		return errors

}

