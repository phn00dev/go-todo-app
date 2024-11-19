package main

import (
	"log"

	"github.com/phn00dev/go-todo-app"
	"github.com/phn00dev/go-todo-app/pkg/handler"
	"github.com/phn00dev/go-todo-app/pkg/repository"
	"github.com/phn00dev/go-todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
