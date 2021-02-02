package service

import (
	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

// TodoItemService ...
type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

// NewTodoItemService ...
func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

// Create ...
func (s *TodoItemService) Create(userID, listID int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list does not exists
		return 0, err
	}

	return s.repo.Create(listID, item)
}

// GetAll ...
func (s *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}

// GetByID ...
func (s *TodoItemService) GetByID(userID, itemID int) (todo.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

// DeleteByID ...
func (s *TodoItemService) DeleteByID(userID, itemID int) error {
	return s.repo.DeleteByID(userID, itemID)
}

// Update ...
func (s *TodoItemService) Update(userID, itemID int, input todo.UpdateItemInput) error {
	return s.repo.Update(userID, itemID, input)
}
