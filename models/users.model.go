package models

type UserModel struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(userId, username, password string) *UserModel {
	return &UserModel{
		UserID:   userId,
		Username: username,
		Password: password,
	}
}
