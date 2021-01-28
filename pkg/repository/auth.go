package repository

import (
	"fmt"

	"github.com/fautherdtd/user-restapi/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Auth ...
type Auth struct {
	db *sqlx.DB
}

// NewAuth ...
func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

// CreateUserWithPhone ...
func (r *Auth) CreateUserWithPhone(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, phone) values ($1, $2) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Phone)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Error 'CreateUserWithPhone': %s", err.Error())
		return 0, err
	}

	return id, nil
}

// GetUserByID ...
func (r *Auth) GetUserByID(id int) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)

	err := r.db.Get(user, query, id)
	return user, err
}

// GetUserByPhone ...
func (r *Auth) GetUserByPhone(phone string) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=$1", userTable)

	err := r.db.Get(user, query, phone)
	return user, err
}

// ConfirmUserByCode ...
func (r *Auth) ConfirmUserByCode(id int) error {
	query := fmt.Sprintf("UPDATE %s SET verifed=true WHERE id=$1", userTable)
	_, err := r.db.Exec(query, id)
	return err
}

// CheckUserVerification ...
func (r *Auth) CheckUserVerification(id int) (bool, error) {
	var verifed bool
	query := fmt.Sprintf("SELECT verifed FROM %s WHERE id=$1", userTable)

	err := r.db.Get(&verifed, query, id)
	return verifed, err
}
