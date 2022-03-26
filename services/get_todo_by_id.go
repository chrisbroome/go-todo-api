package services

import (
	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type GetTodoByIDService struct {
	db db.TodoGetter
}

func NewGetTodoByIDService(db db.TodoGetter) *GetTodoByIDService {
	return &GetTodoByIDService{db: db}
}

func (this *GetTodoByIDService) GetTodoByID(id string) (*entities.Todo, error) {
	return this.db.GetTodoById(id)
}
