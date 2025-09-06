package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rwrrioe/todolist/todo"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func respondError(c *gin.Context, err error, status int) {
	errorDTO := newErrorDTO(err, time.Now())
	c.JSON(status, errorDTO)
}

/*
pattern: tasks
methon: POST
info: JSON in HTTP request body

succeed:
	-status code 201 Created
	-response body: JSON represents created task
failed
	- status code: 400,409,500...
	-response body: JSON is error, time
*/

func (h *HTTPHandlers) HandlerCreateTask(c *gin.Context) {
	var taskDTO TaskDTO
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		respondError(c, err, 400)
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		respondError(c, err, 400)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		respondError(c, err, 500)
		return

	}
	c.JSON(http.StatusCreated, todoTask)
}

/*
pattern: tasks/{title}
method: GET
info: pattern

succeed :
	- status code 200 OK
	- response body: JSON represents created task

failed:
	-statuc code 400, 404, 500...
	- responce body: JSON with error, time
*/

func (h *HTTPHandlers) HandlerGetTask(c *gin.Context) {
	title := c.Param("title")

	task, err := h.todoList.GetTask(title)
	if err != nil {
		respondError(c, err, 404)
		return
	}

	c.JSON(200, task)
}

/*
pattern: tasks/ tasks/?completed=false
method: GET
info: pattern

succeed :
	- status code 200 OK
	- response body: JSON represents created tasks

failed:
	-statuc code 400, 404, 500...
	- responce body: JSON with error, time
*/

func (h *HTTPHandlers) HandlerGetAllTasks(c *gin.Context) {
	isCompleted := c.Query("completed")

	if isCompleted != "" {
		if isCompleted, err := strconv.ParseBool(isCompleted); err != nil {
			respondError(c, err, 400)
			return
		} else if !isCompleted {
			tasks := h.todoList.ListUncompletedTasks()
			c.JSON(200, tasks)
			return
		}
	}

	tasks := h.todoList.ListTasks()
	c.JSON(200, tasks)
}

/*
pattern: tasks/{title}
method: PATCH
info: pattern + JSON in the request body

succeed :
  - status code 200 OK
  - response body: JSON represents changed tasks

failed:

	-statuc code 400, 409, 500...
	-responce body: JSON with error, time
*/
func (h *HTTPHandlers) HandlerCompleteTask(c *gin.Context) {
	completeDTO := CompleteDTO{false}
	if err := c.ShouldBindJSON(&completeDTO); err != nil {
		respondError(c, err, 400)
		return
	}

	title := c.Param("title")

	if completeDTO.Complete {
		if err := h.todoList.CompleteTask(title); err != nil {
			respondError(c, err, 404)
			return
		}
	} else {
		if err := h.todoList.UncompleteTask(title); err != nil {
			respondError(c, err, 404)
			return
		}
	}

	task, err := h.todoList.GetTask(title)
	if err != nil {
		respondError(c, err, 404)
		return
	}

	c.JSON(200, task)
}

/*
pattern: tasks/{title}
method: DELETE
info: pattern

succeed :
  - status code 204 No Content
  - response body: -

failed:

	-statuc code 400, 404,500...
	-responce body: JSON with error, time
*/
func (h *HTTPHandlers) HandlerDeleteTask(c *gin.Context) {
	title := c.Param("title")
	if err := h.todoList.DeleteTask(title); err != nil {
		respondError(c, err, 404)
		return
	}

	c.Status(204)
}
