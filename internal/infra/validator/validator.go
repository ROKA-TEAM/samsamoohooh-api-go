package validator

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/pkg/box"

	stdvalidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	engine *stdvalidator.Validate
}

func (v Validator) Validate(out any) error {
	err := v.engine.Struct(out)
	if err != nil {
		return box.Wrap(domain.ErrValidation, err.Error())
	}

	return nil
}

func New() *Validator {
	return &Validator{
		engine: stdvalidator.New(stdvalidator.WithRequiredStructEnabled()),
	}
}
