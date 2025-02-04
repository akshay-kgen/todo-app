package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/services"
	"github.com/akshay-kgen/todo-app/types"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReqModel types.CreateUserReqModel

	if err := json.NewDecoder(r.Body).Decode(&userReqModel); err != nil {
		helpers.SendHandlerErrResponse(w, "failed to decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userReqModel.ToNewUser()
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	customError := h.authService.Register(r.Context(), user)
	if customError != nil {
		helpers.SendHandlerCustomErrResponse(w, customError, http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
