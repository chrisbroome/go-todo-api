package memory_test

import (
	"testing"

	"github.com/chrisbroome/go-todo-api/db/memory"
	"github.com/chrisbroome/go-todo-api/idgenerator/random"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"

	"github.com/chrisbroome/go-todo-api/db"
)

type memoryDbTest struct {
	idg *random.IdGenerator
	db  *memory.Db
	*assertions.Assertion
}

func newMemoryDbTest(t *testing.T) *memoryDbTest {
	idg := random.NewIdGenerator()
	return &memoryDbTest{
		idg:       idg,
		db:        memory.NewDb(idg),
		Assertion: assertions.New(t),
	}
}

func TestNewDb(t *testing.T) {
	test := newMemoryDbTest(t)
	todos, err := test.db.ListTodos()
	test.So(err, should.BeNil)
	test.So(todos, should.BeEmpty)
}

func TestCreateTodo(t *testing.T) {
	test := newMemoryDbTest(t)

	const label = "a"
	todo, err := test.db.CreateTodo(label)
	test.So(todo, should.NotBeNil)
	test.So(err, should.BeNil)
	test.So(todo.Label, should.Equal, label)

	todos, err := test.db.ListTodos()
	test.So(err, should.BeNil)
	test.So(todos, should.HaveLength, 1)
}

func TestGetTodo(t *testing.T) {
	test := newMemoryDbTest(t)
	createdTodo, _ := test.db.CreateTodo("Pick up stuff")
	todo, err := test.db.GetTodoById(createdTodo.Id)

	test.So(err, should.BeNil)
	test.So(createdTodo, should.Resemble, todo)
}

func TestDeleteWhenItemExistsFoundShouldReturnTrue(t *testing.T) {
	test := newMemoryDbTest(t)
	createdTodo, err := test.db.CreateTodo("This will be deleted")
	test.So(err, should.BeNil)
	test.So(createdTodo, should.NotBeNil)

	todos, err := test.db.ListTodos()
	test.So(err, should.BeNil)
	test.So(todos, should.NotBeEmpty)
	lengthBefore := len(todos)

	found, err := test.db.DeleteTodo(createdTodo.Id)
	if test.So(err, should.BeNil) {
		test.So(found, should.BeTrue)

		todos, err = test.db.ListTodos()
		test.So(err, should.BeNil)
		test.So(todos, should.BeEmpty)
		lengthAfter := len(todos)
		test.So(lengthAfter, should.Equal, lengthBefore-1)
	}
}

func TestDeleteWhenItemDoesNotExistFoundShouldReturnFalse(t *testing.T) {
	test := newMemoryDbTest(t)
	todos, _ := test.db.ListTodos()
	lengthBefore := len(todos)
	found, err := test.db.DeleteTodo("")
	todos, _ = test.db.ListTodos()
	lengthAfter := len(todos)
	test.So(found, should.BeFalse)
	test.So(err, should.BeNil)
	test.So(lengthAfter, should.Equal, lengthBefore)
}

func TestUpdateTodoWhenTodoExists(t *testing.T) {
	test := newMemoryDbTest(t)
	createdTodo, _ := test.db.CreateTodo("This will be updated")
	updatesToMake := &db.TodoUpdateInput{
		Label:     "Updated Label",
		Completed: true,
	}
	updatedTodo, err := test.db.UpdateTodo(createdTodo.Id, updatesToMake)
	if test.So(err, should.BeNil) {
		test.So(updatedTodo.Label, should.Equal, updatesToMake.Label)
		test.So(updatedTodo.Completed, should.Equal, updatesToMake.Completed)
	}
}

func TestUpdateTodoWhenTodoDoesNotExist(t *testing.T) {
	test := newMemoryDbTest(t)
	updatesToMake := &db.TodoUpdateInput{
		Label:     "Updated Label",
		Completed: true,
	}
	updatedTodo, err := test.db.UpdateTodo("id that doesn't exist", updatesToMake)
	test.So(err, should.BeNil)
	test.So(updatedTodo, should.BeNil)
}

// func TestNextIdDifferentFromPreviousId(t *testing.T) {
// 	test := newMemoryDbTest(t)
// 	id1 := test.db.nextId()
// 	id2 := test.db.nextId()
// 	test.So(id1, should.NotEqual, id2)
// }
