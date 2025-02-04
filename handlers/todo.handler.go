package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/middlewares"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/akshay-kgen/todo-app/types"
	"github.com/golang-jwt/jwt/v4"
)

type TodoHandler struct {
	todoService *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: service,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	decodedUser, ok := ctx.Value(middlewares.UserContextKey).(jwt.MapClaims)
	if !ok {
		helpers.SendHandlerErrResponse(w, "No user found", http.StatusUnauthorized)
		return
	}

	userId, ok := decodedUser["userId"].(string)
	if !ok {
		helpers.SendHandlerErrResponse(w, "Invalid user data", http.StatusUnauthorized)
		return
	}

	var todoRequestModel *types.CreateTodoReqModel
	err := json.NewDecoder(r.Body).Decode(&todoRequestModel)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	todo, err := todoRequestModel.ToNewTodo()
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	todo.UserId = userId

	createdTodo, err := h.todoService.CreateTodo(ctx, todo)
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}
