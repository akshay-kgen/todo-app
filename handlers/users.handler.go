package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akshay-kgen/todo-app/helpers"
	"github.com/akshay-kgen/todo-app/serializers"
	"github.com/akshay-kgen/todo-app/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {

	userId, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializedUser := serializers.SerializeUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serializedUser)
}
