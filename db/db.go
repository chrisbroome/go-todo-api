package db

import "github.com/chrisbroome/go-todo-api/entities"

type (
	TodoCreator interface {
		CreateTodo(label string) (*entities.Todo, error)
	}

	TodoGetter interface {
		GetTodoById(id string) (*entities.Todo, error)
	}

	Db interface {
		CreateTodo(label string) (*entities.Todo, error)
		DeleteTodo(id string) (bool, error)
		UpdateTodo(id string, input *TodoUpdateInput) (*entities.Todo, error)
		GetTodoById(id string) (*entities.Todo, error)
		ListTodos() ([]*entities.Todo, error)
	}

	TodoUpdateInput struct {
		Label     string
		Completed bool
	}
)
