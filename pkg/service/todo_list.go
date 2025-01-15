package service

import (
	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}