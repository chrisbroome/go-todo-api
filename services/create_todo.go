package services

import (
	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type CreateTodoService struct {
	db db.Db
}

func NewCreateTodoService(db db.Db) *CreateTodoService {
	return &CreateTodoService{db: db}
}

func (this *CreateTodoService) CreateTodo(label string) (*entities.Todo, error) {
	return this.db.CreateTodo(label)
}
