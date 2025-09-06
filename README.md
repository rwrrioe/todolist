# Todo List (Gin version)

This is a refactored version of my Todo List project built with the [Gin](https://github.com/gin-gonic/gin) web framework.

## Features
- Create a task (`POST /tasks`)
- Get a task by title (`GET /tasks/:title`)
- Get all tasks or filter by completion (`GET /tasks?completed=true|false`)
- Mark task as complete/incomplete (`PATCH /tasks/:title`)
- Delete a task (`DELETE /tasks/:title`)
- Error responses are returned as JSON with message and timestamp

## Run the server
```bash
go run main.go
```

The server will start at [http://localhost:9091](http://localhost:9091).

## Example requests

### Create a task
```bash
curl -X POST http://localhost:9091/tasks -H "Content-Type: application/json" -d '{"title": "learn Gin", "description": "practice with middleware"}'
```

### Get all tasks
```bash
curl http://localhost:9091/tasks
```

### Get only uncompleted tasks
```bash
curl http://localhost:9091/tasks?completed=false
```

### Mark a task as complete
```bash
curl -X PATCH http://localhost:9091/tasks/learn%20Gin -H "Content-Type: application/json" -d '{"complete": true}'
```

### Delete a task
```bash
curl -X DELETE http://localhost:9091/tasks/learn%20Gin
```

---

## Branch info
- This branch (`gin-todo`) contains the Gin implementation.  
- The `main` branch contains the version using `net/http` and `mux`.  
