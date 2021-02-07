package service

import (
	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

// Authorization ...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

// TodoList ...
type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, listID int) (todo.TodoList, error)
	DeleteByID(userID, listID int) error
	Update(userID, listID int, input todo.UpdateListInput) error
}

// TodoItem ...
type TodoItem interface {
	Create(userID, listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	DeleteByID(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

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
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
