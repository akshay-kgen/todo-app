package routes

import (
	"github.com/akshay-kgen/todo-app/handlers"
	"github.com/akshay-kgen/todo-app/middlewares"
	"github.com/akshay-kgen/todo-app/repo"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
)

func TodoRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	todoRepo := repo.NewTodoRepo(ddb)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	return func(r chi.Router) {
		r.Use(middlewares.Authenticate)
		r.Post("/", todoHandler.CreateTodo)

	}
}
