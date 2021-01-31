package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

// NewpostgresDB ...
func NewpostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
