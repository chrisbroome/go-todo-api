package todos

import (
	"github.com/chrisbroome/go-todo-api/entities"
)

type TodoResponse struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	Completed bool   `json:"completed"`
}

func toTodoResponse(todo *entities.Todo) *TodoResponse {
	if todo == nil {
		return nil
	}
	return &TodoResponse{
		Id:        todo.Id,
		Label:     todo.Label,
		Completed: todo.Completed,
	}
}
