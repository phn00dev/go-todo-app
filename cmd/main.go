package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db : %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
