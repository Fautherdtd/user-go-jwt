package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Lists of tables
const (
	userTable = "users"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB ...
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	fmt.Print(err)

	if err != nil {
		return nil, err
	}
	return db, nil
}
