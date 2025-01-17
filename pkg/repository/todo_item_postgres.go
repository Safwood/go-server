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

