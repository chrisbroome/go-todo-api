package memory

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type Db struct {
	todos map[string]*entities.Todo
}

func NewDb() *Db {
	return &Db{
		todos: make(map[string]*entities.Todo),
	}
}

func (this *Db) nextId() string {
	buf := make([]byte, 20)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

func (this *Db) CreateTodo(label string) (*entities.Todo, error) {
	id := this.nextId()
	todo := &entities.Todo{
		Id:    id,
		Label: label,
	}
	this.todos[id] = todo
	return todo, nil
}

func (this *Db) DeleteTodo(id string) error {
	delete(this.todos, id)
	return nil
}

func (this *Db) GetTodoById(id string) (*entities.Todo, error) {
	todo := this.todos[id]
	return todo, nil
}

func (this *Db) UpdateTodo(id string, input *db.TodoUpdateInput) (*entities.Todo, error) {
	todo, _ := this.GetTodoById(id)
	if todo == nil {
		return nil, nil
	}

	updatedTodo := &entities.Todo{
		Id:        id,
		Label:     input.Label,
		Completed: input.Completed,
	}
	this.todos[id] = updatedTodo
	return updatedTodo, nil
}