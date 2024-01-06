#### Todo API

A simple RESTful Todo list API.

##### Endpoints

- `POST /todos`: Create a new todo.
- `GET /todos`: Retrieve all todos.
- `GET /todos/:id`: Retrieve a todo by ID.
- `PUT /todos/:id`: Update a todo by ID.
- `DELETE /todos/:id`: Delete a todo by ID.


##### Setup and Run

1. Clone the Repository
```bash
git clone https://github.com/tabintel/todo-api.git
cd todo-api
```
2. Run this command to install all the dependencies
```bash
go mod tidy
```

3. Run the application with this command
```bash
go run main.go

```

Here is the [Postman API Documentation,](https://documenter.getpostman.com/view/31909794/2s9YsJAXoT) for testing the API endpoints.
