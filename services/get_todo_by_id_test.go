package services_test

import (
	"testing"

	"github.com/chrisbroome/go-todo-api/entities"
	"github.com/chrisbroome/go-todo-api/idgenerator/random"
	"github.com/chrisbroome/go-todo-api/mock"
	"github.com/chrisbroome/go-todo-api/services"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

type getTodoByIDTest struct {
	idg *random.IdGenerator
	db  *mock.Db
	svc *services.GetTodoByIDService
	*assertions.Assertion
}

func newGetTodoByIDTest(t *testing.T) *getTodoByIDTest {
	idg := random.NewIdGenerator()
	db := mock.NewDb()
	svc := services.NewGetTodoByIDService(db)
	return &getTodoByIDTest{
		idg:       idg,
		db:        db,
		svc:       svc,
		Assertion: assertions.New(t),
	}
}

func TestShouldFetchTheTodoFromTheDatabase(t *testing.T) {
	test := newGetTodoByIDTest(t)
	id := test.idg.GenerateID()
	todo := entities.NewTodo("label")
	todo.Id = id
	test.db.WithTodo(todo)

	todo, err := test.svc.GetTodoByID(id)
	test.So(err, should.BeNil)
	test.So(test.db.Called, should.BeTrue)
	test.So(test.db.CalledWith, should.Equal, id)
	test.So(todo.Id, should.Equal, id)
}
