package main

import (
	"log"

	todo "github.com/rwrrioe/todolist"
)

func main() {
	server := new(todo.Server)
	if err := server.Run("8080"); err != nil {
		log.Fatalf("error ocurred while running http server:%s", err.Error())
	}
}
