package routes

import (
	"github.com/akshay-kgen/todo-app/handlers"
	"github.com/akshay-kgen/todo-app/middlewares"
	"github.com/akshay-kgen/todo-app/repo"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	userRepo := repo.NewUserRepo(ddb)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	return func(r chi.Router) {
		r.Use(middlewares.Authenticate)
		r.Get("/me", userHandler.GetMe)
	}
}
