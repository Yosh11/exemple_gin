package service

import (
	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

// TodoListService ...
type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService ...
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create ...
func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

// GetAll ...
func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

// GetByID ...
func (s *TodoListService) GetByID(userID, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

// DeleteByID ...
func (s *TodoListService) DeleteByID(userID, listID int) error {
	return s.repo.DeleteByID(userID, listID)
}

// Update ...
func (s *TodoListService) Update(userID, listID int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return nil
	}
	return s.repo.Update(userID, listID, input)
}
