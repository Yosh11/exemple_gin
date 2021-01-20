package repository

import "github.com/jmoiron/sqlx"

// Authorization ...
type Authorization interface{}

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
	return &Repository{}
}
