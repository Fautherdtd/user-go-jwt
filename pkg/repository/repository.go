package repository

import (
	"github.com/fautherdtd/user-restapi/entities"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUserWithPhone(user entities.User) (int, error)
	GetUserByID(id int) (entities.User, error)
	// GetUserByPhone(string phone) (entities.User, error)
	ConfirmUserByCode(id int) error
	CheckUserVerification(id int) (bool, error)
}

// Users ...
type Users interface {
}

// Repository ...
type Repository struct {
	Authorization
	Users
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		Users:         NewUser(db),
	}
}
