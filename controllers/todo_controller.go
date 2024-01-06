// controllers/todo_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabintel/todo-api/models"
	"github.com/tabintel/todo-api/services"
	"github.com/tabintel/todo-api/utils"
)

// TodoController struct
type TodoController struct {
	logger      *utils.Logger
	todoService *services.TodoService
}

// NewTodoController creates a new instance of TodoController
func NewTodoController(logger *utils.Logger, todoService *services.TodoService) *TodoController {
	return &TodoController{
		logger:      logger,
		todoService: todoService,
	}
}

// CreateTodo handles the creation of a new todo
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	requestID := ctx.GetString("RequestID")
	c.logger.LogRequest(requestID, "CreateTodo")

	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		c.logger.LogError("Failed to bind JSON", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := c.todoService.CreateTodo(&todo); err != nil {
		c.logger.LogError("Failed to create todo", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

// GetTodos retrieves all todos
func (c *TodoController) GetTodos(ctx *gin.Context) {
	requestID := ctx.GetString("RequestID")
	c.logger.LogRequest(requestID, "GetTodos")

	todos, err := c.todoService.GetTodos()
	if err != nil {
		c.logger.LogError("Failed to get todos", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// GetTodoByID retrieves a todo by its ID
func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	requestID := ctx.GetString("RequestID")
	c.logger.LogRequest(requestID, "GetTodoByID")

	id := ctx.Param("id")

	todo, err := c.todoService.GetTodoByID(id)
	if err != nil {
		c.logger.LogError("Failed to get todo by ID", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// UpdateTodo updates a todo
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	requestID := ctx.GetString("RequestID")
	c.logger.LogRequest(requestID, "UpdateTodo")

	id := ctx.Param("id")
	var updatedTodo models.Todo

	if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
		c.logger.LogError("Failed to bind JSON", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := c.todoService.UpdateTodo(id, &updatedTodo); err != nil {
		c.logger.LogError("Failed to update todo", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	ctx.JSON(http.StatusOK, updatedTodo)
}

// DeleteTodo removes a todo
func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	requestID := ctx.GetString("RequestID")
	c.logger.LogRequest(requestID, "DeleteTodo")

	id := ctx.Param("id")

	if err := c.todoService.DeleteTodo(id); err != nil {
		c.logger.LogError("Failed to delete todo", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
