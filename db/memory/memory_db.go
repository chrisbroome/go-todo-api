package memory

import (
	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
)

type Db struct {
	idg   entities.IdGenerator
	todos map[string]*entities.Todo
}

func NewDb(idg entities.IdGenerator) *Db {
	return &Db{
		idg:   idg,
		todos: make(map[string]*entities.Todo),
	}
}

func (db *Db) nextId() string {
	return db.idg.GenerateID()
}

func (db *Db) CreateTodo(label string) (*entities.Todo, error) {
	id := db.nextId()
	todo := &entities.Todo{
		Id:    id,
		Label: label,
	}
	db.todos[id] = todo
	return todo, nil
}

func (db *Db) DeleteTodo(id string) (bool, error) {
	_, found := db.todos[id]
	delete(db.todos, id)
	return found, nil
}

func (db *Db) GetTodoById(id string) (*entities.Todo, error) {
	todo := db.todos[id]
	return todo, nil
}

func (db *Db) UpdateTodo(id string, input *db.TodoUpdateInput) (*entities.Todo, error) {
	todo, _ := db.GetTodoById(id)
	if todo == nil {
		return nil, nil
	}

	updatedTodo := &entities.Todo{
		Id:        id,
		Label:     input.Label,
		Completed: input.Completed,
	}
	db.todos[id] = updatedTodo
	return updatedTodo, nil
}

func (db *Db) ListTodos() ([]*entities.Todo, error) {
	ret := make([]*entities.Todo, len(db.todos))
	i := 0
	for _, value := range db.todos {
		ret[i] = value
		i++
	}
	return ret, nil
}
