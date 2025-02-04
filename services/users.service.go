package services

import (
	"fmt"

	"github.com/akshay-kgen/todo-app/models"
	"github.com/akshay-kgen/todo-app/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) GetUserById(userId string) (*models.UserModel, error) {

	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user details: %w", err)
	}

	return user, nil
}
