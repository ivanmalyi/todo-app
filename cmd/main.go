package main

import (
	"github.com/ivanmalyi/todo-app"
	"github.com/ivanmalyi/todo-app/pkg/handler"
	"github.com/ivanmalyi/todo-app/pkg/repository"
	"github.com/ivanmalyi/todo-app/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)

	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while run server: %v", err)
	}

}
