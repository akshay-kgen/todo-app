package routes

import (
	"github.com/akshay-kgen/todo-app/handlers"
	"github.com/akshay-kgen/todo-app/repo"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	userRepo := repo.NewUserRepo(ddb)
	userService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	return func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	}
}
