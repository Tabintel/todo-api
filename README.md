# Todo List API

## Overview

This project implements a simple Todo List API using the Gin-gonic framework, MongoDB for data storage, and follows the Controller-Service-Repository architecture. The API provides endpoints to perform CRUD operations on todo items.

## Project Structure

The project follows a standard Go project structure:

- `controllers`: Contains the controllers handling HTTP requests.
- `models`: Defines the data models used in the application.
- `repositories`: Manages the interaction with the MongoDB database.
- `services`: Implements the business logic for todo-related operations.
- `tests`: Contains unit tests for the application.

## Technologies Used

- **Gin-gonic:** A web framework written in Go.
- **MongoDB:** Used as the database for storing todo items.
- **Uber Zap:** A logging library used for creating well-structured log files.

## Setup

1. **Install Go:**
   - [Download and Install Go](https://golang.org/doc/install)

2. **Install Dependencies:**
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u go.mongodb.org/mongo-driver/mongo
   go get -u go.uber.org/zap
