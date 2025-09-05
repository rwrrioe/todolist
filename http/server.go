package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	router.POST("/tasks", s.HTTPHandlers.HandlerCreateTask)
	router.GET("/tasks/:title", s.HTTPHandlers.HandlerGetTask)
	router.GET("/tasks", s.HTTPHandlers.HandlerGetAllTasks)
	router.PATCH("/tasks/:title", s.HTTPHandlers.HandlerCompleteTask)
	router.DELETE("/tasks/:title", s.HTTPHandlers.HandlerDeleteTask)

	return http.ListenAndServe(":9091", router)
}
