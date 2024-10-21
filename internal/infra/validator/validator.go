package validator

import stdvalidator "github.com/go-playground/validator/v10"

type Validator struct {
	engine *stdvalidator.Validate
}

func (v Validator) Validate(out any) error {
	return v.engine.Struct(out)
}

func New() *Validator {
	return &Validator{
		engine: stdvalidator.New(stdvalidator.WithRequiredStructEnabled()),
	}
}
