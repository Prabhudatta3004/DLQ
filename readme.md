Here's the `README.md` file you can copy and add directly to your project:

```markdown
# DLQ Service

A robust Dead Letter Queue (DLQ) service built with Go, using PostgreSQL for persistent storage and Gin Gonic for the web framework. This service allows you to add, retrieve, delete, and clear messages in a dead letter queue through a RESTful API.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Project Structure](#project-structure)
4. [Installation and Setup](#installation-and-setup)
   - [Clone the Repository](#clone-the-repository)
   - [Database Setup](#database-setup)
   - [Environment Configuration](#environment-configuration)
5. [Running the Service](#running-the-service)
6. [API Documentation](#api-documentation)
   - [Add a Message](#add-a-message)
   - [Get a Message by ID](#get-a-message-by-id)
   - [Get a Message by MessageID](#get-a-message-by-messageid)
   - [Get All Messages](#get-all-messages)
   - [Delete a Message by ID](#delete-a-message-by-id)
   - [Delete a Message by MessageID](#delete-a-message-by-messageid)
   - [Clear All Messages](#clear-all-messages)
7. [Error Handling](#error-handling)
8. [Logging](#logging)
9. [Configuration Details](#configuration-details)
10. [Testing the Service](#testing-the-service)
11. [Deployment Considerations](#deployment-considerations)
12. [Additional Enhancements](#additional-enhancements)
13. [Conclusion](#conclusion)

---

## Introduction

The DLQ Service provides a reliable and efficient way to manage messages that could not be processed successfully in your application. By storing these messages in a dedicated dead letter queue, you can inspect, reprocess, or handle them as needed without losing important data.

---

## Prerequisites

Ensure you have the following installed on your system:

- **Go** (version 1.16 or higher)
- **PostgreSQL** (version 10 or higher)
- **Git**

---

## Project Structure

```
dlq_service/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── controllers/
│   └── dlq_controller.go
├── models/
│   └── message.go
├── repository/
│   └── message_repository.go
├── routers/
│   └── routers.go
├── services/
│   └── dlq_service.go
├── utils/
│   ├── logger.go
│   └── response.go
├── .env.example
├── go.mod
└── go.sum
```

---

## Installation and Setup

### Clone the Repository

```bash
git clone https://github.com/Prabhudatta3004/DLQ.git
cd DLQ
```

### Database Setup

#### Install PostgreSQL

If PostgreSQL is not installed, you can install it using Homebrew on macOS:

```bash
brew install postgresql
brew services start postgresql
```

#### Create Database and User

Access the PostgreSQL shell:

```bash
psql postgres
```

Create a new database and user:

```sql
CREATE DATABASE dlq_service;
CREATE USER dlq_user WITH PASSWORD 'dlq_password';
GRANT ALL PRIVILEGES ON DATABASE dlq_service TO dlq_user;
\q
```

### Environment Configuration

Copy the `.env.example` file to `.env`:

```bash
cp .env.example .env
```

Update the `.env` file with your database credentials:

```env
# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=dlq_user
DB_PASSWORD=dlq_password
DB_NAME=dlq_service

# Server configuration
SERVER_PORT=8080
```

---

## Running the Service

### Install Dependencies

```bash
go mod tidy
```

### Run the Application

```bash
go run ./cmd/server/main.go
```

You should see output indicating the server is running on port `8080`.

---

## API Documentation

### Add a Message

- **Endpoint**: `POST /dlq/message`
- **Description**: Adds a new message to the DLQ.
- **Request Body**:

  ```json
  {
    "message_id": "unique-message-id",
    "payload": "Your message content"
  }
  ```

- **Response**:

  ```json
  {
    "message": "Message added to DLQ",
    "data": {
      "id": 1,
      "message_id": "unique-message-id",
      "payload": "Your message content",
      "created_at": "2024-11-14T21:20:00Z"
    }
  }
  ```

### Get a Message by ID

- **Endpoint**: `GET /dlq/message/id/:id`
- **Description**: Retrieves a message from the DLQ by its numeric ID.
- **Response**:

  ```json
  {
    "id": 1,
    "message_id": "unique-message-id",
    "payload": "Your message content",
    "created_at": "2024-11-14T21:20:00Z"
  }
  ```

### Get a Message by MessageID

- **Endpoint**: `GET /dlq/message/message_id/:message_id`
- **Description**: Retrieves a message from the DLQ by its `message_id`.
- **Response**:

  ```json
  {
    "id": 1,
    "message_id": "unique-message-id",
    "payload": "Your message content",
    "created_at": "2024-11-14T21:20:00Z"
  }
  ```

### Get All Messages

- **Endpoint**: `GET /dlq/messages`
- **Description**: Retrieves all messages in the DLQ.
- **Response**:

  ```json
  [
    {
      "id": 1,
      "message_id": "unique-message-id",
      "payload": "Your message content",
      "created_at": "2024-11-14T21:20:00Z"
    },
    {
      "id": 2,
      "message_id": "another-message-id",
      "payload": "Another message content",
      "created_at": "2024-11-14T21:25:00Z"
    }
  ]
  ```

### Delete a Message by ID

- **Endpoint**: `DELETE /dlq/message/id/:id`
- **Description**: Deletes a message from the DLQ by its numeric ID.
- **Response**:

  ```json
  {
    "message": "Message deleted from DLQ"
  }
  ```

### Delete a Message by MessageID

- **Endpoint**: `DELETE /dlq/message/message_id/:message_id`
- **Description**: Deletes a message from the DLQ by its `message_id`.
- **Response**:

  ```json
  {
    "message": "Message deleted from DLQ"
  }
  ```

### Clear All Messages

- **Endpoint**: `DELETE /dlq/messages`
- **Description**: Deletes all messages from the DLQ.
- **Response**:

  ```json
  {
    "message": "All messages cleared from DLQ"
  }
  ```

---

## Error Handling

The service provides consistent error responses with appropriate HTTP status codes.

- **400 Bad Request**: Invalid request data or parameters.
- **404 Not Found**: Requested resource not found.
- **500 Internal Server Error**: Server encountered an unexpected condition.

**Error Response Format**:

```json
{
  "error": "Detailed error message"
}
```

---

## Logging

The service uses Logrus for structured logging.

- **Info Logs**: Informational messages about the service operation.
- **Error Logs**: Errors encountered during request processing.

Logs are outputted in JSON format to `stdout`.

---

## Configuration Details

Configuration is managed via environment variables, loaded from a `.env` file.

### Configurable Parameters

- **DB_HOST**: Database host (default: `localhost`).
- **DB_PORT**: Database port (default: `5432`).
- **DB_USER**: Database user (default: `postgres`).
- **DB_PASSWORD**: Database password.
- **DB_NAME**: Database name (default: `dlq_service`).
- **SERVER_PORT**: Port on which the server runs (default: `8080`).

### Changing Configuration

Modify the `.env` file or set environment variables directly.

---

## Testing the Service

You can test the service using tools like **cURL** or **Postman**.

### Example cURL Commands

#### Add a Message

```bash
curl -X POST http://localhost:8080/dlq/message \
  -H "Content-Type: application/json" \
  -d '{"message_id":"test123","payload":"Hello, World!"}'
```

#### Get All Messages

```bash
curl http://localhost:8080/dlq/messages
```

#### Get a Message by ID

```bash
curl http://localhost:8080/dlq/message/id/1
```

#### Delete a Message by MessageID

```bash
curl -X DELETE http://localhost:8080/dlq/message/message_id/test123
```

---

## Deployment Considerations

### Switching to Release Mode

For production, set Gin to release mode to enhance performance:

```go
import "github.com/gin-gonic/gin"

func main() {
    gin.SetMode(gin.ReleaseMode)
    // Rest of your code...
}
```

### Configuring Trusted Proxies

If deploying behind a proxy, configure trusted proxies:

```go
router := gin.Default()
router.SetTrustedProxies([]string{"<PROXY_IP_ADDRESS>"})
```

---

## Additional Enhancements

### Authentication and Authorization

Implement token-based authentication (e.g., JWT) to secure endpoints.

### Rate Limiting

Use middleware to limit the number of requests from a client.

### HTTPS Support

Configure TLS certificates to serve the API over HTTPS.

### Dockerization

Create a `Dockerfile` to containerize the application for easier deployment.

### Continuous Integration/Deployment

Set up CI/CD pipelines using tools like GitHub Actions or Jenkins.

---

## Conclusion

You've successfully set up and run the DLQ Service. This service provides a robust API for managing a Dead Letter Queue with persistent storage using PostgreSQL.

---

## Need Further Assistance?

Feel free to reach out if you need help with:

- **Extending Functionality**: Adding new features or modifying existing ones.
- **Performance Tuning**: Optimizing the service for better performance.
- **Security Enhancements**: Implementing advanced security measures.
- **Deployment Strategies**: Best practices for deploying in different environments.

Happy coding!
```