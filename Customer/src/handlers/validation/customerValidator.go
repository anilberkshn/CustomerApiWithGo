package validation

import "github.com/go-playground/validator/v10"

type CustomerValidator struct {
	validator *validator.Validate
}

func (c *CustomerValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}
