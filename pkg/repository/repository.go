package repository

import (
	web "github.com/Yosh11/exemple_gin"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUser(user web.User) (int, error)
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
