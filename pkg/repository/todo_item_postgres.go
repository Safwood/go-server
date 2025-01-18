package repository

import (
	"fmt"
	"strings"

	todo "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

	query := fmt.Sprintf("SELECT tit.id, tit.title, tit.description, tit.done FROM %s tit INNER JOIN %s lit on tit.id = lit.item_id INNER JOIN %s ul on ul.list_id = lit.list_id WHERE lit.list_id = $1 AND ul.user_id = $2", todoItemTable, listItemTable, usersListsTable)
	err := r.db.Select(&items, query, listId, userId)

	return items, err
}

func (r *TodoItemPostgres) GetItemById(userId, itemId int) (todo.TodoItem, error)  {
	var item todo.TodoItem

	query := fmt.Sprintf("SELECT tit.id, tit.title, tit.description, tit.done FROM %s tit INNER JOIN %s lit on tit.id = lit.item_id INNER JOIN %s ul on ul.list_id = lit.list_id WHERE lit.item_id = $1 AND ul.user_id = $2", todoItemTable, listItemTable, usersListsTable)
	err := r.db.Get(&item, query, itemId, userId)

	return item, err
}

func (r *TodoItemPostgres) DeleteItem(userId, itemId int) (error)  {
	query := fmt.Sprintf("DELETE FROM %s tit USING %s lit, %s ul WHERE lit.item_id = $1 AND ul.user_id = $2", todoItemTable, listItemTable, usersListsTable)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemTable, setQuery, listItemTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	logrus.Debugf(`query: %d`,query)
	logrus.Debugf(`query: %d`, args)

	_, err := r.db.Exec(query, args...)
	return err
}