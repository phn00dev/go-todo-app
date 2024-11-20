package service

import (
	"github.com/phn00dev/go-todo-app"
	"github.com/phn00dev/go-todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {

	_, err := s.listRepo.GetOne(userId, listId)
	if err != nil {

		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, listId int, inputs todo.UpdateItemInput) error {
	if err := inputs.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, inputs)
}
