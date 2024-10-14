package validator

import stdvalidator "github.com/go-playground/validator/v10"

type Validator struct {
	validate *stdvalidator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: stdvalidator.New(stdvalidator.WithRequiredStructEnabled()),
	}
}

func (v *Validator) Validate(out any) error {
	return v.validate.Struct(out)
}
