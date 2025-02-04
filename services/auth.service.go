package services

import (
	"context"
	"fmt"
	"log"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/models"
	"github.com/akshay-kgen/todo-app/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(repo *repo.UserRepo) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, user *models.UserModel) *helpers.CustomError {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return helpers.NewCustomError(fmt.Errorf("internal error while hashing password: %v", err), "500")
	}

	user.Password = string(hashedPassword)

	if err := s.userRepo.CreateUser(user); err != nil {
		log.Printf("failed to create user in repository: %v", err)
		return helpers.NewCustomError(fmt.Errorf("error creating user: %v", err), "500")
	}

	return nil
}
