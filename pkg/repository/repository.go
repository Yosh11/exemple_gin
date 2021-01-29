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
type TodoList interface{}

// TodoItem ...
type TodoItem interface{}

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
	}
}
