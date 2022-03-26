package usecases

import (
	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type GetTodoInteractor struct {
	db db.TodoGetter
}

func NewGetTodoInteractor(db db.TodoGetter) *GetTodoInteractor {
	return &GetTodoInteractor{db: db}
}

func (this *GetTodoInteractor) Execute(id string) (*entities.Todo, error) {
	return this.db.GetTodoById(id)
}
