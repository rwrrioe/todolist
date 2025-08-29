package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func HandleError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	if errors.Is(err, todo.ErrTaskNotFound) {
		status = http.StatusConflict
	} else if errors.Is(err, todo.ErrTaskNotFound) {
		status = http.StatusNotFound
	}

	errDTO := newErrorDTO(err, time.Now())
	http.Error(w, errDTO.ToString(), status)

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
		errDTO := newErrorDTO(err, time.Now())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := newErrorDTO(err, time.Now())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		HandleError(w, err)
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
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		HandleError(w, err)
		return
	}
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
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
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	uncompletedTasks := h.todoList.ListUncompletedTasks()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	CompleteDTO := CompleteDTO{false}
	if err := json.NewDecoder(r.Body).Decode(&CompleteDTO); err != nil {
		errDTO := newErrorDTO(err, time.Now())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	title := mux.Vars(r)["title"]

	if CompleteDTO.Complete {
		if err := h.todoList.CompleteTask(title); err != nil {
			HandleError(w, err)
			return
		} else {
			if err := h.todoList.UncompleteTask(title); err != nil {
				HandleError(w, err)
				return
			}
		}

	}
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
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
