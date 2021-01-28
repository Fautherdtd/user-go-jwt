package service

import (
	"github.com/fautherdtd/user-restapi/entities"
	"github.com/fautherdtd/user-restapi/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUserWithPhone(user entities.User) (int, error)
	ConfirmUserByCode(id int) error
	GenerateSmsCode(id int, phone string) (bool, error)
	GenerateToken(user entities.User) (string, error)
	CheckUserVerification(id int) (bool, error)
}

// Users ...
type Users interface {
}

// Service ...
type Service struct {
	Authorization
	Users
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUserService(repos.Users),
	}
}
