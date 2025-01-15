package service

import (
	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
}

type TodoItem interface {
	
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: newTodoListService(repos.TodoList),
	}
}