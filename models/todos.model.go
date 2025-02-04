package models

type TodoModel struct {
	TodoID      string `json:"todo_id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTodo(todoId, userId, title, description, status string) *TodoModel {
	return &TodoModel{
		TodoID:      todoId,
		UserID:      userId,
		Title:       title,
		Description: description,
		Status:      status,
	}
}
