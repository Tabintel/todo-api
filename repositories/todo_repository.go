// repositories/todo_repository.go
package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/tabintel/todo-api/models"
	"github.com/tabintel/todo-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "todoDB"
	todoCollection = "todos"
)

// TodoRepository struct
type TodoRepository struct {
	logger *utils.Logger
	client *mongo.Client
}

// NewTodoRepository creates a new instance of TodoRepository
func NewTodoRepository(logger *utils.Logger, client *mongo.Client) *TodoRepository {
	return &TodoRepository{
		logger: logger,
		client: client,
	}
}

// CreateTodo adds a new todo to the database
func (r *TodoRepository) CreateTodo(todo *models.Todo) error {
	collection := r.client.Database(dbName).Collection(todoCollection)

	todo.ID = utils.GenerateUUID()
	todo.Completed = false

	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		r.logger.LogError("Failed to insert todo into database", err)
		return fmt.Errorf("failed to insert todo into database")
	}

	return nil
}

// GetTodos retrieves all todos from the database
func (r *TodoRepository) GetTodos() ([]models.Todo, error) {
	collection := r.client.Database(dbName).Collection(todoCollection)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		r.logger.LogError("Failed to fetch todos from database", err)
		return nil, fmt.Errorf("failed to fetch todos from database")
	}
	defer cursor.Close(context.Background())

	var todos []models.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		r.logger.LogError("Failed to decode todos", err)
		return nil, fmt.Errorf("failed to decode todos")
	}

	return todos, nil
}

// GetTodoByID retrieves a todo by its ID from the database
func (r *TodoRepository) GetTodoByID(id string) (*models.Todo, error) {
	collection := r.client.Database(dbName).Collection(todoCollection)

	var todo models.Todo
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&todo)
	if err != nil {
		r.logger.LogError("Failed to fetch todo by ID from database", err)
		return nil, fmt.Errorf("failed to fetch todo by ID from database")
	}

	return &todo, nil
}

// UpdateTodo updates a todo in the database
func (r *TodoRepository) UpdateTodo(id string, updatedTodo *models.Todo) error {
	collection := r.client.Database(dbName).Collection(todoCollection)

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{
				{"title", updatedTodo.Title},
				{"description", updatedTodo.Description},
				{"completed", updatedTodo.Completed},
			}},
		},
	)
	if err != nil {
		r.logger.LogError("Failed to update todo in database", err)
		return fmt.Errorf("failed to update todo in database")
	}

	return nil
}

// DeleteTodo removes a todo from the database
func (r *TodoRepository) DeleteTodo(id string) error {
	collection := r.client.Database(dbName).Collection(todoCollection)

	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		r.logger.LogError("Failed to delete todo from database", err)
		return fmt.Errorf("failed to delete todo from database")
	}

	return nil
}

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
