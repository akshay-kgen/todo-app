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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq types.LoginReq

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := loginReq.Validate()
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, customError := h.authService.Login(r.Context(), loginReq.Email, loginReq.Password)
	if customError != nil {
		helpers.SendHandlerCustomErrResponse(w, customError, http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
