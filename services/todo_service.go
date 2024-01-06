// services/todo_service.go
package services

import (
	"github.com/tabintel/todo-api/models"
	"github.com/tabintel/todo-api/repositories"
	"github.com/tabintel/todo-api/utils"
)

// TodoService struct
type TodoService struct {
	logger        *utils.Logger
	todoRepository *repositories.TodoRepository
}

// NewTodoService creates a new instance of TodoService
func NewTodoService(logger *utils.Logger, todoRepository *repositories.TodoRepository) *TodoService {
	return &TodoService{
		logger:        logger,
		todoRepository: todoRepository,
	}
}

// CreateTodo adds a new todo
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	s.logger.LogRequest(todo.ID, "CreateTodo")
	return s.todoRepository.CreateTodo(todo)
}

// GetTodos retrieves all todos
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	s.logger.LogRequest("", "GetTodos")
	return s.todoRepository.GetTodos()
}

// GetTodoByID retrieves a todo by its ID
func (s *TodoService) GetTodoByID(id string) (*models.Todo, error) {
	s.logger.LogRequest(id, "GetTodoByID")
	return s.todoRepository.GetTodoByID(id)
}

// UpdateTodo updates a todo
func (s *TodoService) UpdateTodo(id string, updatedTodo *models.Todo) error {
	s.logger.LogRequest(id, "UpdateTodo")
	return s.todoRepository.UpdateTodo(id, updatedTodo)
}

// DeleteTodo removes a todo
func (s *TodoService) DeleteTodo(id string) error {
	s.logger.LogRequest(id, "DeleteTodo")
	return s.todoRepository.DeleteTodo(id)
}
