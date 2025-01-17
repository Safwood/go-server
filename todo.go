package todo

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
	Id int `json:"id"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

type ListItem struct {
	Id int `json:"id"`
	ListId string `json:"list_id"`
	ItemId string `json:"item_id"`
}