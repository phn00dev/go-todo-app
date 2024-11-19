package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/phn00dev/go-todo-app"
	"github.com/phn00dev/go-todo-app/pkg/repository"
)

const solt = "asdad&^2312556t&hbdfs4)sdagd*$Y"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
