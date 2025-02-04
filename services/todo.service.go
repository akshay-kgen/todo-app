package services

import (
	"context"
	"errors"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/akshay-kgen/todo-app/repo"
)

type TodoService struct {
	todoRepo *repo.TodoRepo
}

func NewTodoService(repo *repo.TodoRepo) *TodoService {
	return &TodoService{
		todoRepo: repo,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, todo *models.TodoModel) (*models.TodoModel, error) {
	if todo.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	err := s.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
