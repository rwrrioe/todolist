package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func newHTTPServer(HTTPHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: HTTPHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.HTTPHandlers.HandlerCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.HTTPHandlers.HandlerGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.HTTPHandlers.HandlerGetAllTasks)
	router.Path("/tasks").Methods("GET").Queries("complited", "false").HandlerFunc(s.HTTPHandlers.HandlerGetAllUncompletedTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.HTTPHandlers.HandlerCompleteTask)
	router.Path("tasks/{title}").Methods("DELETE").HandlerFunc(s.HTTPHandlers.HandlerDeleteTask)

	return http.ListenAndServe(":9091", router)
}
