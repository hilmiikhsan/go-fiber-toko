package common

import (
	"github.com/go-playground/validator/v10"
)

func Validate(modelValidate interface{}) []map[string]interface{} {
	validate := validator.New()
	var messages []map[string]interface{}
	err := validate.Struct(modelValidate)
	if err != nil {
		// var messages model.GeneralResponse
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}
	}

	return messages
}
