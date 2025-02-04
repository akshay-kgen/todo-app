package serializers

import (
	"github.com/akshay-kgen/todo-app/models"
	"github.com/akshay-kgen/todo-app/types"
)

func SerializeUser(user *models.UserModel) *types.UserResponseModel {
	return &types.UserResponseModel{
		UserId:    user.UserId,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
