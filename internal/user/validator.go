package user

import "github.com/go-playground/validator/v10"

type Validator struct {
	v *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		v: validator.New(),
	}
}

func (v *Validator) Validate(u *User) error {
	return v.v.Struct(u)
}
