package repository

import (
	sights "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user sights.User) (int, error)
	GetUser(username string, password string) (sights.User, error)
}

type TodoList interface {
	Create(userId int, list sights.TodoList) (int, error)
	GetAllLists(userId int) ([]sights.TodoList, error)
	GetListById(userId int, listId int) (sights.TodoList, error)
}

type TodoItem interface {
	CreateItem(listId int, item sights.TodoItem) (int, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
	Park
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db), 
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
		Park: NewParkPostgres(db),
	}
}