package validator

import (
	"samsamoohooh-go-api/internal/infra/exception"

	stdvalidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	engine *stdvalidator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		engine: stdvalidator.New(stdvalidator.WithRequiredStructEnabled()),
	}
}

func (v Validator) Validate(out any) error {
	err := v.engine.Struct(out)
	if err != nil {
		return exception.New(
			err,
			exception.ErrWebServerValidate,
			exception.StatusBadRequest,
			"Validation failed",
		)
	}

	return nil
}
