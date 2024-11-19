package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/phn00dev/go-todo-app"
	"github.com/phn00dev/go-todo-app/pkg/handler"
	"github.com/phn00dev/go-todo-app/pkg/repository"
	"github.com/phn00dev/go-todo-app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Printf("error initializing configs : %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
