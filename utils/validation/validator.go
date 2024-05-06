package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

func NewEchoValidator() echo.Validator {
	return &EchoValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}
