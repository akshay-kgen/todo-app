package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/akshay-kgen/todo-app/repo"
	"github.com/akshay-kgen/todo-app/types"
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

func (s *TodoService) GetAllTodo(ctx context.Context, userId string) ([]*models.TodoModel, error) {
	todos, err := s.todoRepo.GetAllTodo(userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetTodo(ctx context.Context, userId, todoId string) (*models.TodoModel, error) {
	todo, err := s.todoRepo.GetTodo(userId, todoId)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, userId, todoId string, todoRequestModel *types.UpdateTodoReqModel) (*models.TodoModel, error) {
	existingTodo, err := s.todoRepo.GetTodo(userId, todoId)
	if err != nil {
		return nil, err
	}

	if todoRequestModel.Title != nil {
		existingTodo.Title = *todoRequestModel.Title
	}
	if todoRequestModel.Description != nil {
		existingTodo.Description = *todoRequestModel.Description
	}
	if todoRequestModel.Status != nil {
		existingTodo.Status = *todoRequestModel.Status
	}
	existingTodo.UpdatedAt = time.Now()

	err = s.todoRepo.UpdateTodo(existingTodo)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	return existingTodo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, userId, todoId string) error {

	err := s.todoRepo.DeleteTodo(userId, todoId)
	if err != nil {
		return err
	}
	return nil
}
