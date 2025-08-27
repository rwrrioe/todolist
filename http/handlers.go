package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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

func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
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

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: tasks
method: GET
info: pattern

succeed :
	- status code 200 OK
	- response body: JSON represents created tasks

failed:
	-statuc code 400, 404, 500...
	- responce body: JSON with error, time
*/

func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: tasks?completed=false
method: GET
info: query params

succeed :
  - status code 200 OK
  - response body: JSON represents created tasks

failed:

	-statuc code 400, 404, 500...
	- responce body: JSON with error, time
*/
func (h *HTTPHandlers) HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

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
func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {

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
func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {

}
