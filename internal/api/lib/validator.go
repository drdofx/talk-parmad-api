package lib

import "github.com/go-playground/validator/v10"

func ValidatorInit() *validator.Validate {
	validate := validator.New()
	validate.RegisterAlias("req-email", "required,email")
	validate.RegisterAlias("req-alphanum", "required,alphanum")
	validate.RegisterAlias("req-alpha", "required,alpha")
	validate.RegisterAlias("req-numeric", "required,numeric")
	validate.RegisterAlias("req-dive", "required,dive")
	return validate
}
