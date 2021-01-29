package service

import (
	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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
