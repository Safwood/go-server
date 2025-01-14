package repository

import (
	"fmt"

	todo "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error)  {
	// начало транзакций
	tx, err := r.db.Begin()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	todoListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(todoListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	if _, err := tx.Exec(usersListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
	// конец транзакций
}