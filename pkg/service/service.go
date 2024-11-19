package service

import "github.com/phn00dev/go-todo-app/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{}
}
