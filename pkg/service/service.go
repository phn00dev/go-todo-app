package service

import (
	"github.com/phn00dev/go-todo-app"
	"github.com/phn00dev/go-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
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

func NewServices(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
