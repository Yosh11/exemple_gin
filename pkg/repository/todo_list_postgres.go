package repository

import (
	"fmt"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/jmoiron/sqlx"
)

// TodoListPostgres ...
type TodoListPostgres struct {
	db *sqlx.DB
}

// NewTodoListPostgres ...
func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

// Create ...
func (r *TodoListPostgres) Create(userID int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// GetAll ...
func (r *TodoListPostgres) GetAll(userID int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userID)

	return lists, err
}

// GetByID ...
func (r *TodoListPostgres) GetByID(userID, listID int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
        INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userID, listID)

	return list, err
}
