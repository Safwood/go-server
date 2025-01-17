package service

import (
	todo "github.com/Safwood/go-server"
	"github.com/Safwood/go-server/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func newTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo, listRepo}
}

func (s *TodoItemService) CreateItem(userId, listId int, item todo.TodoItem) (int, error) {
	// проверяем что такой список и пользователь существуют
	if _, err := s.listRepo.GetListById(userId, listId); err != nil {
		return 0, err
	}
	return s.repo.CreateItem(listId, item)
}

func (s *TodoItemService) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId, listId int) (todo.TodoItem, error) {
	return s.repo.GetItemById(userId, listId)
}

func (s *TodoItemService) DeleteItem(userId, listId int) (error) {
	return s.repo.DeleteItem(userId, listId)
}

func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}