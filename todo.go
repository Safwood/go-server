package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id int  `json:"id" db:"id"`
	UserId string  `json:"user_id" db:"user_id"`
	ListId string  `json:"list_id" db:"list_id"`
}

type TodoItem struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done bool `json:"done" db:"done"`
}

type ListItem struct {
	Id int `json:"id" db:"id"`
	ListId string `json:"list_id" db:"list_id"`
	ItemId string `json:"item_id" db:"item_id"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}