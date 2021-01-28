package service

import (
	"github.com/fautherdtd/user-restapi/pkg/repository"
)

// UserService ...
type UserService struct {
	repo repository.Users
}

// NewUserService ...
func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}
