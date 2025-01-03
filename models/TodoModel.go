package models

type Todo struct {
	TodoID      int    `json:"todo_id" db:"todo_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserName    string `json:"name" db:"name"` // Assuming this is the user's name
	CreatedOn   string `json:"created_on" db:"created_on"`
	Target      string `json:"target" db:"target"`
	TypeID      int    `json:"type_id" db:"type_id"`
	TypeName    string `json:"type_name" db:"type_name"`
}
