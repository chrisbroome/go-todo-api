package usecases

import (
	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type CreateTodoInteractor struct {
	db db.Db
}

func NewCreateTodoInteractor(db db.Db) *CreateTodoInteractor {
	return &CreateTodoInteractor{db: db}
}

func (this *CreateTodoInteractor) Execute(label string) (*entities.Todo, error) {
	return this.db.CreateTodo(label)
}
