package main

import (
	"myproject/todo/pkg/configs"
	repository "myproject/todo/pkg/repositories/todos"
)

func main() {
	config := &configs.Config{
		Port: ":3000",
	}

	app := NewApplication(*config, &repository.LocalTodoRepository{})

	app.StartServer()
}
