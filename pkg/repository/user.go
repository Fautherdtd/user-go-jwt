package repository

import (
	"github.com/jmoiron/sqlx"
)

// User ...
type User struct {
	db *sqlx.DB
}

// NewUser ...
func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}
