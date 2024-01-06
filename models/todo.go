// models/todo.go
package models

// Todo struct represents a todo item
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
