package main

import (
	"os"

	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/handler"
	"github.com/Safwood/go-server/pkg/repository"
	"github.com/Safwood/go-server/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main()  {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	
	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка подключения конфиг файла %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка подключения env файла %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("No connection with DB", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Ошибка подключения сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}