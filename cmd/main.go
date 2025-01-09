package main

import (
	"log"

	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/handler"
	"github.com/Safwood/go-server/pkg/repository"
	"github.com/Safwood/go-server/pkg/service"
	"github.com/spf13/viper"
)

func main()  {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка подключения конфиг файла %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Ошибка подключения сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}