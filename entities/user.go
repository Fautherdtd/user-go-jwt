package entities

import "time"

// User ...
type User struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Email    string    `json:"email"`
	Gender   string    `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
}

// Confirm ...
type Confirm struct {
	UserID int    `json:"user_id" binding:"required"`
	Phone  string `json:"phone"`
	Code   int    `json:"code"`
}

// SignUp ...
type SignUp struct {
	Token   string `json:"token"`
	Confirm bool   `json:"confirm"`
}
