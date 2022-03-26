package mock

import "github.com/chrisbroome/go-todo-api/entities"

type Db struct {
	Called     bool
	CalledWith string
	Todo       *entities.Todo
	Err        error
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) WithTodo(todo *entities.Todo) *Db {
	db.Todo = todo
	return db
}

func (db *Db) WithError(err error) *Db {
	db.Err = err
	return db
}

func (db *Db) GetTodoById(id string) (*entities.Todo, error) {
	db.Called = true
	db.CalledWith = id
	return db.Todo, db.Err
}
