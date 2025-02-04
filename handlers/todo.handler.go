package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/akshay-kgen/todo-app/types"
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
	userId, contextError := helpers.GetUserIDFromContext(ctx)
	if contextError != nil {
		helpers.SendHandlerErrResponse(w, contextError.Error(), http.StatusUnauthorized)
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

func (h *TodoHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {

	userId, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	todos, err := h.todoService.GetAllTodo(r.Context(), userId)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
