package main

import (
	"log"

	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/handler"
)

func main()  {
	srv := new(todo.Server)
	handler := new(handler.Handler)
	

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Ошибка подключения сервера: %s", err.Error())
	}
}