package main

import (
	"fmt"
	"go-validation-translator/validator"
)

type Request struct {
	Name     string `json:"name" validate:"required,min=1,max=50" trans:"名前"`
	Email    string `json:"email" validate:"required,min=1,max=50" trans:"メアド"`
	Password string `json:"password" validate:"required,min=1,max=50" trans:"パスワード"`
}

func main() {
	validationTranslator := validator.NewValidationTranslator()
	request := &Request{
		Name:     "",
		Email:    "",
		Password: "",
	}

	errors := validationTranslator.Validate(request)
	fmt.Println(errors)
}
