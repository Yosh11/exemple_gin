package web

// TodoList ...
type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UsersList ...
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItems ...
type TodoItems struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListsItem ...
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
