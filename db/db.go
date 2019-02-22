package db

import "github.com/chrisbroome/go-todo-api/entities"

type (
	Db interface {
		CreateTodo(label string) (*entities.Todo, error)
		DeleteTodo(id string) error
		UpdateTodo(id string, input *TodoUpdateInput) (*entities.Todo, error)
		GetTodoById(id string) (*entities.Todo, error)
	}

	TodoUpdateInput struct {
		Label     string
		Completed bool
	}
)
