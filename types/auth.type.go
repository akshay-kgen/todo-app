package types

import "github.com/go-playground/validator/v10"

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (model *LoginReq) Validate() error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}
