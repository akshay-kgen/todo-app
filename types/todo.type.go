package types

import (
	"fmt"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/go-playground/validator/v10"
)

type CreateTodoReqModel struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=500"`
}

func (model *CreateTodoReqModel) Validate() error {
	validate := validator.New()

	err := validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func (model *CreateTodoReqModel) ToNewTodo() (*models.TodoModel, error) {

	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("create todo request validation failed: %w", err)
	}

	todo := models.NewTodo("", model.Title, model.Description, "")

	return todo, nil
}

type UpdateTodoReqModel struct {
	Title       *string `json:"title" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,min=5,max=500"`
	Status      *string `json:"status" validate:"omitempty,oneof=created in-progress completed"`
}

func (model *UpdateTodoReqModel) Validate() error {
	validate := validator.New()

	err := validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}
