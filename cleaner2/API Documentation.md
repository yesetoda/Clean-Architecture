
## API Documentation

### Authentication

- **Signup**

  - **Endpoint**: `/signup`
  - **Method**: `POST`
  - **Description**: Create a new user account.
  - **Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Successfully added the user"
    }
    ```

- **Login**

  - **Endpoint**: `/login`
  - **Method**: `POST`
  - **Description**: Authenticate a user and generate a JWT token.
  - **Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "token": "JWT_TOKEN"
    }
    ```

### Task Management

- **Create Task**

  - **Endpoint**: `/task`
  - **Method**: `POST`
  - **Description**: Create a new task (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Body**:
    ```json
    {
      "id": "int",
      "title": "string",
      "description": "string",
      "duedate": "string",
      "status": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Successfully added the task"
    }
    ```

- **Get All Tasks**

  - **Endpoint**: `/task`
  - **Method**: `GET`
  - **Description**: Retrieve all tasks (User and Admin).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    [
      {
        "id": "int",
        "title": "string",
        "description": "string",
        "duedate": "string",
        "status": "string"
      }
    ]
    ```

- **Get Task by ID**

  - **Endpoint**: `/task/:id`
  - **Method**: `GET`
  - **Description**: Retrieve a task by its ID (User and Admin).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "int",
      "title": "string",
      "description": "string",
      "duedate": "string",
      "status": "string"
    }
    ```

- **Update Task**

  - **Endpoint**: `/task/:id`
  - **Method**: `PATCH`
  - **Description**: Update a task by its ID (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Body**:
    ```json
    {
      "title": "string",
      "description": "string",
      "duedate": "string",
      "status": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Update successful"
    }
    ```

- **Delete Task**

  - **Endpoint**: `/task/:id`
  - **Method**: `DELETE`
  - **Description**: Delete a task by its ID (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Delete successful"
    }
    ```

- **Filter Tasks**

  - **Endpoint**: `/task/filter`
  - **Method**: `GET`
  - **Description**: Filter tasks based on criteria (User and Admin).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Query Parameters**:
    - `status`: Filter by task status.
    - `duedate`: Filter by task due date.
  - **Response**:
    ```json
    [
      {
        "id": "int",
        "title": "string",
        "description": "string",
        "duedate": "string",
        "status": "string"
      }
    ]
    ```

### User Management

- **Get All Users**

  - **Endpoint**: `/user`
  - **Method**: `GET`
  - **Description**: Retrieve all users (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    [
      {
        "username": "string",
        "role": "string"
      }
    ]
    ```

- **Get User by Username**

  - **Endpoint**: `/user/:username`
  - **Method**: `GET`
  - **Description**: Retrieve a user by their username (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    {
      "username": "string",
      "role": "string"
    }
    ```

- **Delete User**

  - **Endpoint**: `/user/:username`
  - **Method**: `DELETE`
  - **Description**: Delete a user by their username (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Delete successful"
    }
    ```

- **Promote User**

  - **Endpoint**: `/user/:username`
  - **Method**: `PATCH`
  - **Description**: Promote a user to admin (Admin only).
  - **Headers**:
    ```json
    {
      "Authorization": "Bearer JWT_TOKEN"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Promotion successful"
    }
    ```

