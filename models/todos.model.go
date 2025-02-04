package models

import (
	"time"

	"github.com/google/uuid"
)

type TodoModel struct {
	TodoId      string    `json:"todoId"`
	UserId      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodo(userId, title, description string, status string) *TodoModel {
	if status == "" {
		status = "created"
	}
	return &TodoModel{
		TodoId:      uuid.New().String(),
		UserId:      userId,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
