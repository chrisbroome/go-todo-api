package services_test

import (
	"testing"

	"github.com/chrisbroome/go-todo-api/db/memory"
	"github.com/chrisbroome/go-todo-api/idgenerator/random"
	"github.com/chrisbroome/go-todo-api/services"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

type createTodoTest struct {
	idg        *random.IdGenerator
	db         *memory.Db
	interactor *services.CreateTodoService
	*assertions.Assertion
}

func newCreateTodoTest(t *testing.T) *createTodoTest {
	idg := random.NewIdGenerator()
	memDb := memory.NewDb(idg)
	return &createTodoTest{
		idg:        idg,
		db:         memDb,
		interactor: services.NewCreateTodoService(memDb),
		Assertion:  assertions.New(t),
	}
}

func TestWhenGivenALabelShouldCreateTheTodo(t *testing.T) {
	test := newCreateTodoTest(t)
	label := "my todo"
	todo, err := test.interactor.CreateTodo(label)
	if test.So(err, should.BeNil) {
		test.So(todo.Label, should.Equal, label)
	}
}
