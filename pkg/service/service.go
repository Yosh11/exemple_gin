package service

import (
	web "github.com/Yosh11/exemple_gin"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUser(user web.User) (int, error)
}

// TodoList ...
type TodoList interface{}

// TodoItem ...
type TodoItem interface{}

// Service ...
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
