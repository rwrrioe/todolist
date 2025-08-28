# Golang REST API: To-Do List Example (Educational Project)

This repository serves as a **hands-on educational resource**, showcasing how to build and structure a simple REST API in Go (Golang). It is based on the “Chapter 2: Productive Golang and REST API” section from a full Go course.

## Purpose & Scope

- **Educational Focus**: Designed for learners exploring REST API development in Go. Step-by-step code illustrates best practices, common patterns, and idiomatic Go.
- **Hands-On Practice**: You can clone the project, run the API locally, and interact with its endpoints using tools like `curl` or Postman.
- **Key Concepts Demonstrated**:
  - Routing using `github.com/gorilla/mux`
  - Handling HTTP methods (GET, POST, DELETE, PATCH)
  - Parsing path variables and query parameters
  - JSON encoding/decoding with `encoding/json`
  - Basic error handling and status codes
  - Simple in-memory task storage

## Features Overview

- **GET /tasks** — retrieve a list of all tasks
- **GET /tasks/{title}** — fetch a specific task by title
- **POST /tasks** — create a new task
- **DELETE /tasks/{title}** — remove a task by title
- **PATCH /tasks/{title}/complete** — mark a task as completed

Each endpoint responds with a JSON object or an error message, following REST principles. The project also includes error handling examples, returning appropriate HTTP status codes when tasks are not found or a request is invalid.

## Getting Started

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd your-repo-directory
   ```

2. Build and run the server:
   ```bash
   go run main.go
   ```

3. Interact with the API:
   ```bash
   # Create a new task
   curl -X POST http://localhost:9091/tasks         -H "Content-Type: application/json"         -d '{"title":"Buy milk"}'

   # Get all tasks
   curl http://localhost:8080/tasks

   # Complete a task
   curl -X PATCH http://localhost:9091/tasks/Buy%20milk/complete
   ```

4. Observe status codes like `201 Created`, `404 Not Found`, and `200 OK`.

## Why This Project Is Useful

- **Beginner-Friendly**: Clear structure and comments make it easy to follow.
- **Extendable**: Use as a base to add features like persistent storage, user authentication, or middleware.
- **Solid Foundation**: Reinforces HTTP fundamentals and Go idiomatic patterns.

---

### Note

This project is **purely for learning purposes**, not production-ready. It uses in-memory storage, lacks input validation, authentication, and other security or scalability features. Nevertheless, it’s a great starting point for hands-on practice and experimentation.

Happy coding!
