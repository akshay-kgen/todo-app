package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/repo"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/akshay-kgen/todo-app/types"
	"github.com/go-chi/chi/v5"
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

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {

	todoId := chi.URLParam(r, "id")
	if todoId == "" {
		customErr := helpers.NewCustomError(errors.New("missing todoId in the request path"), "MISSING_TODO_ID")
		helpers.SendHandlerCustomErrResponse(w, customErr, http.StatusBadRequest)
		return
	}

	userId, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		customErr := helpers.NewCustomError(err, "UNAUTHORIZED_ACCESS")
		helpers.SendHandlerCustomErrResponse(w, customErr, http.StatusUnauthorized)
		return
	}

	todo, err := h.todoService.GetTodo(r.Context(), userId, todoId)
	if err != nil {
		if errors.Is(err, repo.ErrTodoNotFound) {
			customErr := helpers.NewCustomError(errors.New("todo not found"), "TODO_NOT_FOUND")
			helpers.SendHandlerCustomErrResponse(w, customErr, http.StatusNotFound)
		} else {
			customErr := helpers.NewCustomError(errors.New("failed to fetch todo"), "FETCH_TODO_ERROR")
			helpers.SendHandlerCustomErrResponse(w, customErr, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")
	if todoId == "" {
		helpers.SendHandlerErrResponse(w, "Missing todoId in the request path", http.StatusBadRequest)
		return
	}

	userId, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var todoRequestModel types.UpdateTodoReqModel
	err = json.NewDecoder(r.Body).Decode(&todoRequestModel)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	validationError := todoRequestModel.Validate()
	if validationError != nil {
		helpers.SendHandlerErrResponse(w, validationError.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedTodo, err := h.todoService.UpdateTodo(r.Context(), userId, todoId, &todoRequestModel)
	if err != nil {
		if err == repo.ErrTodoNotFound {
			helpers.SendHandlerErrResponse(w, "Todo not found", http.StatusNotFound)
		} else {
			helpers.SendHandlerErrResponse(w, "Failed to update todo", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}
