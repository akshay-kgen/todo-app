package models

import "time"

type TodoModel struct {
	TodoID      string    `json:"todoId"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodo(todoId, userId, title, description, status string) *TodoModel {
	return &TodoModel{
		TodoID:      todoId,
		UserID:      userId,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
