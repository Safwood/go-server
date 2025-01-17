package repository

import (
	"fmt"

	todo "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}

func (r *TodoItemPostgres) CreateItem(listId int, item todo.TodoItem) (int, error)  {
	// начало транзакций
	tx, err := r.db.Begin()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var itemId int
	todoItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES ($1, $2, $3) RETURNING id", todoItemTable)
	row := tx.QueryRow(todoItemQuery, item.Title, item.Description, item.Done)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	listItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemTable)
	if _, err := tx.Exec(listItemQuery, listId, itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
	// конец транзакций
}

func (r *TodoItemPostgres) GetAllItems(userId, listId int) ([]todo.TodoItem, error)  {
	var items []todo.TodoItem

	query := fmt.Sprintf("SELECT tit.title, tit.description, tit.done FROM %s tit INNER JOIN %s lit on tit.id = lit.item_id INNER JOIN %s ul on ul.list_id = lit.list_id WHERE lit.list_id = $1 AND ul.user_id = $2", todoItemTable, listItemTable, usersListsTable)
	err := r.db.Select(&items, query, listId, userId)

	return items, err
}

func (r *TodoItemPostgres) GetItemById(userId, itemId int) (todo.TodoItem, error)  {
	var item todo.TodoItem

	query := fmt.Sprintf("SELECT tit.title, tit.description, tit.done FROM %s tit INNER JOIN %s lit on tit.id = lit.item_id INNER JOIN %s ul on ul.list_id = lit.list_id WHERE lit.item_id = $1 AND ul.user_id = $2", todoItemTable, listItemTable, usersListsTable)
	err := r.db.Get(&item, query, itemId, userId)

	return item, err

}

