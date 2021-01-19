package service

import "github.com/Yosh11/exemple_gin/pkg/repository"

// Authorization ...
type Authorization interface{}

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
	return &Service{}
}
