package service

import (
	sights "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user sights.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list sights.TodoList) (int, error)
	GetAllLists(userId int) ([]sights.TodoList, error)
	GetListById(userId int, listId int) (sights.TodoList, error)
}

type TodoItem interface {
	CreateItem(userId, listId int, todoList sights.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]sights.TodoItem, error)
	GetItemById(userId, listId int) (sights.TodoItem, error)
	DeleteItem(userId, listId int) (error)
	Update(userId, itemId int, input sights.UpdateItemInput) error
}

type Park interface {
	CreatePark(userId int, park sights.Park) (int, error)
	GetAllParks(userId int) ([]sights.Park, error)
	GetParkById(userId, parkId int) (sights.GetParkOutput, error)
	DeletePark(userId, parkId int) (error)
	UpdatePark(userId, parkId int, input sights.UpdateParkInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	Park
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: newTodoListService(repos.TodoList),
		TodoItem: newTodoItemService(repos.TodoItem, repos.TodoList),
		Park: newParkService(repos.Park),
	}
}