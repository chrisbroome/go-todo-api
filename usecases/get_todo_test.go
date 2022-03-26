package usecases

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"

	"github.com/chrisbroome/go-todo-api/entities"
)

func TestGetTodoFixture(t *testing.T) {
	gunit.Run(new(GetTodoFixture), t)
}

type GetTodoFixture struct {
	*gunit.Fixture
	interactor *GetTodoInteractor
	db         *SpyDb
	id         string
	todo       *entities.Todo
}

func (this *GetTodoFixture) Setup() {
	todo := entities.NewTodo("label")
	todo.Id = "some id"
	this.db = NewSpyDb().WithTodo(todo)
	this.interactor = NewGetTodoInteractor(this.db)
}

func (this *GetTodoFixture) TestShouldFetchTheTodoFromTheDatabase() {
	id := "some id"
	todo, _ := this.interactor.Execute(id)
	this.So(this.db.Called, should.BeTrue)
	this.So(this.db.CalledWith, should.Equal, id)
	this.So(todo.Id, should.Equal, id)
}

type SpyDb struct {
	Called     bool
	CalledWith string
	Todo       *entities.Todo
	Err        error
}

func NewSpyDb() *SpyDb {
	return &SpyDb{}
}

func (this *SpyDb) WithTodo(todo *entities.Todo) *SpyDb {
	this.Todo = todo
	return this
}

func (this *SpyDb) WithError(err error) *SpyDb {
	this.Err = err
	return this
}

func (this *SpyDb) GetTodoById(id string) (*entities.Todo, error) {
	this.Called = true
	this.CalledWith = id
	return this.Todo, this.Err
}
