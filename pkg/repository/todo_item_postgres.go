package repository

import (
	"fmt"
	"strings"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/jmoiron/sqlx"
)

// TodoItemPostgres ...
type TodoItemPostgres struct {
	db *sqlx.DB
}

// NewTodoItemPostgres ...
func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

// Create ...
func (r *TodoItemPostgres) Create(listID int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, listID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemID, tx.Commit()
}

// GetAll ...
func (r *TodoItemPostgres) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON
                            li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1
                                AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

// GetByID ...
func (r *TodoItemPostgres) GetByID(userID, itemID int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON
                            li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1
                                AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemID, userID); err != nil {
		return item, err
	}

	return item, nil
}

// DeleteByID ...
func (r *TodoItemPostgres) DeleteByID(userID, itemID int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
                            WHERE ti.id = li.item_id AND li.list_id = ul.list_id
                                AND ul.user_id = $1 AND ti.id = $2`, todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userID, itemID)
	return err
}

// Update ...
func (r *TodoItemPostgres) Update(userID, itemID int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
    WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argID, argID+1)

	args = append(args, userID, itemID)

	_, err := r.db.Exec(query, args...)
	return err
}
