package repository

import (
	"fmt"

	sights "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db}
}

func (r *TodoListPostgres) Create(userId int, list sights.TodoList) (int, error)  {
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

func (r *TodoListPostgres) GetAllLists(userId int) ([]sights.TodoList, error) {
	var lists []sights.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1`, todoListsTable, usersListsTable)
    err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetListById(userId int, listId int) (sights.TodoList, error) {
	var list sights.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListsTable, usersListsTable)
    err := r.db.Get(&list, query, userId, listId)

	return list, err
}