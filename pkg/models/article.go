package models

import (
	"github.com/go-playground/validator/v10"
)

// Article holds our database fields
type Article struct {
	Title   string `json:"title" db:"title" validate:"required,min=3,max=32"`
	Content string `json:"content" db:"content" validate:"required,min=3,max=32"`
	Author  string `json:"author" db:"author" validate:"required,min=3,max=32"`
}

// ErrorResponse contains our struct error validations
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// Validate checks struct data requirements
func (a *Article) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
