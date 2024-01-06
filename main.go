// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tabintel/todo-api/controllers"
	"github.com/tabintel/todo-api/services"
	"github.com/tabintel/todo-api/repositories"
	"github.com/tabintel/todo-api/utils"
)

func main() {
	// Initialize logger
	logger := utils.SetupLogger()

	// Initialize MongoDB connection
	client, err := repositories.InitMongoDB()
	if err != nil {
		logger.LogError("Failed to connect to MongoDB", err)
		return
	}

	// Inject logger and MongoDB client into repositories, services, and controllers
	todoRepository := repositories.NewTodoRepository(logger, client)
	todoService := services.NewTodoService(logger, todoRepository)
	todoController := controllers.NewTodoController(logger, todoService)

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.Use(RequestIDMiddleware()) // Implement RequestIDMiddleware
	r.POST("/todos", todoController.CreateTodo)
	r.GET("/todos", todoController.GetTodos)
	r.GET("/todos/:id", todoController.GetTodoByID)
	r.PUT("/todos/:id", todoController.UpdateTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)

	// Run the server
	r.Run(":8080")
}

// RequestIDMiddleware generates a unique request ID for each incoming request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := utils.GenerateUUID()
		c.Set("RequestID", requestID)
		c.Next()
	}
}
