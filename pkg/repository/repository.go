package repository

import (
	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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
	Create(listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	DeleteByID(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

// Repository ...
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
