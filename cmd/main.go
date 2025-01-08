package main

import (
	"log"

	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/handler"
	"github.com/Safwood/go-server/pkg/repository"
	"github.com/Safwood/go-server/pkg/service"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Ошибка подключения сервера: %s", err.Error())
	}
}