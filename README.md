# JWT Authentication System

This project is a basic JWT-based authentication system built using the **Go** programming language and **Gin** web framework. It includes endpoints for user registration and login, secured routes that require authentication, and token validation using **JSON Web Tokens (JWT)**. 

## Features

- **User Registration**: Allows users to register with a username and password.
- **Login**: Authenticates users and generates a JWT token upon successful login.
- **Protected Routes**: Access certain routes only with a valid JWT token.
- **Password Hashing**: Passwords are securely hashed using bcrypt before storage.
- **JWT Token Validation**: Validates tokens before granting access to protected routes.
- **Environment Configuration**: Secret keys and other sensitive data are managed using environment variables.

## Tech Stack

- **Go**: The primary programming language used for this application.
- **Gin**: A web framework for Go, designed for performance and ease of use.
- **JWT**: Used for generating and validating JSON Web Tokens.
- **Bcrypt**: For secure password hashing.
- **dotenv**: To manage environment variables.

## Getting Started

### Prerequisites

- **Go 1.16+** installed on your local machine.
- **Postman** or **cURL** for testing the API.
- **Git** to clone the repository.

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/jwt-auth-system.git
    cd jwt-auth-system
    ```

2. **Install dependencies**:

    In the project directory, run:

    ```bash
    go mod tidy
    ```

3. **Create a `.env` file**:

    Create a `.env` file in the root of the project and add your `JWT_SECRET_KEY`:

    ```bash
    JWT_SECRET_KEY=your-very-secure-secret-key
    ```

4. **Run the application**:

    Start the server by running:

    ```bash
    go run main.go
    ```

5. **Test the API**:

    You can use **Postman**, **cURL**, or any other API client to test the following endpoints.

## API Endpoints

### 1. **User Registration**

- **Endpoint**: `/register`
- **Method**: `POST`
- **Description**: Registers a new user by accepting a username and password.
- **Request Body**:

    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```

- **Response**:

    ```json
    {
        "message": "User registered successfully"
    }
    ```

### 2. **User Login**

- **Endpoint**: `/login`
- **Method**: `POST`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:

    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```

- **Response**:

    ```json
    {
        "token": "your-jwt-token"
    }
    ```

### 3. **Protected Route**

- **Endpoint**: `/api/dashboard`
- **Method**: `GET`
- **Description**: Access a protected dashboard. Requires a valid JWT token.
- **Headers**:

    ```json
    {
        "Authorization": "Bearer your-jwt-token"
    }
    ```

- **Response**:

    ```json
    {
        "message": "Welcome to the dashboard, [username]"
    }
    ```

## Token Refresh (Optional)

You can implement token refresh logic by adding an endpoint that generates a new JWT if the existing one is valid but expired.

## Environment Variables

This project uses environment variables to manage sensitive information. Create a `.env` file at the root of the project with the following variable:

```bash
JWT_SECRET_KEY=your-very-secure-secret-key
```

Make sure not to expose the `.env` file by adding it to `.gitignore`.

## Future Improvements

1. **Database Integration**: Replace the in-memory user store with a real database like PostgreSQL.
2. **Token Refresh**: Implement token refresh logic for expired tokens.
3. **Password Reset**: Add functionality for users to reset their password.

## Running Tests

You can create unit tests for the individual functions using Go's built-in testing tools. For example:

```bash
go test ./...
```

## License
Patent Pending B)
