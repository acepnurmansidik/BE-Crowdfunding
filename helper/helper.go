package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) *Response {
	meta := Meta{
		Message: message,
		Code: code,
		Status: status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return &jsonResponse
}

func FormatValidationError(err error) []string{
	// var utk menampung error
	var errors []string

	// ubah error menjadi error validatior
	for _, e := range err.(validator.ValidationErrors){
		// simpan setiap error string validator ke dalam slice errors
		errors = append(errors, e.Error())
	}

	return errors
}