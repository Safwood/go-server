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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db), 
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}