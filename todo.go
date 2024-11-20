package todo

type TodoList struct {
	ID          int    `joson:"id" db:"id"`
	Title       string `json:"title"  db:"title"`
	Description string `json:"description"  db:"description"`
}

type UsersList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	ID     int
	UserID int
	ListID int
}
