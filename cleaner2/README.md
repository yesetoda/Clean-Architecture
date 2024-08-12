### README

# Cleaner Task Management API

This is a simple Task Management API built using GoLang, Gin, and MongoDB, following the principles of Clean Architecture. The API allows users to manage tasks and user accounts, providing authentication and authorization features using JWT.

## Features

- **User Management**: 
  - Signup: Create a new user.
  - Login: Authenticate a user and generate a JWT token.
  - Get All Users: Retrieve all users (admin only).
  - Get User by Username: Retrieve a specific user by their username (admin only).
  - Delete User: Delete a user by their username (admin only).
  - Promote User: Promote a user to admin (admin only).

- **Task Management**:
  - Create Task: Create a new task (admin only).
  - Get All Tasks: Retrieve all tasks (user and admin).
  - Get Task by ID: Retrieve a specific task by its ID (user and admin).
  - Update Task: Update an existing task (admin only).
  - Delete Task: Delete a task by its ID (admin only).
  - Filter Tasks: Filter tasks based on criteria like status, due date, etc. (user and admin).

## Project Structure

- `domain/`: Contains the core domain models (User and Task).
- `usecases/`: Contains the use case logic for tasks and users.
- `infrastructure/`: Contains infrastructure code for database interaction, JWT handling, and hashing.
- `delivery/`: The main entry point of the application.

## Prerequisites

- Go 1.18+
- MongoDB
- Git
- Docker (optional)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yesetoda/Clean-Architecture
    cd cleaner2
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up environment variables:

    Create a `.env` file in the root directory and add the following variables:

    ```bash
    MongodbUri=<Your MongoDB URI>
    MongodbName=<Your Database Name>
    TaskCollectionName=tasks
    UserCollectionName=users
    JWT_KEY=<Your JWT Secret Key>
    ```

4. Run the application:

    ```bash
    go run /cmd/taskmanager/main.go
    ```
