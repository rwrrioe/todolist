package main

import (
	"github.com/rwrrioe/todolist/http"
	"github.com/rwrrioe/todolist/todo"
)

func main() {
	task1 := todo.NewTask("homework", "spend 2h preparing homework for tomorrow's SAT")
	task2 := todo.NewTask("go work a walk", "go for a walk with your friends")
	task3 := todo.NewTask("judo training", "spend 2h training in the dojo")

	todolist := todo.NewList()
	todolist.AddTask(task1)
	todolist.AddTask(task2)
	todolist.AddTask(task3)

	handlers := http.NewHTTPHandlers(todolist)
	server := http.NewHTTPServer(handlers)
	server.StartServer()
}
