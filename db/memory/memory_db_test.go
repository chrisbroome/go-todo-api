package memory

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"

	"github.com/chrisbroome/go-todo-api/db"
)

func TestMemoryDbFixture(t *testing.T) {
	gunit.Run(new(MemoryDbFixture), t)
}

type MemoryDbFixture struct {
	*gunit.Fixture
}

func (this *MemoryDbFixture) Setup() {
}

func (this *MemoryDbFixture) TestNewDb() {
	db := NewDb()

	this.So(db.todos, should.BeEmpty)
}

func (this *MemoryDbFixture) TestCreateTodo() {
	memoryDb := NewDb()

	const label = "a"
	todo, err := memoryDb.CreateTodo(label)

	this.So(memoryDb.todos, should.HaveLength, 1)
	this.So(todo, should.NotBeNil)
	this.So(err, should.BeNil)
	this.So(todo.Label, should.Equal, label)
}

func (this *MemoryDbFixture) TestGetTodo() {
	memoryDb := NewDb()
	createdTodo, _ := memoryDb.CreateTodo("Pick up stuff")
	todo, err := memoryDb.GetTodoById(createdTodo.Id)

	this.So(err, should.BeNil)
	this.So(createdTodo, should.Resemble, todo)
}

func (this *MemoryDbFixture) TestDeleteTodo() {
	memoryDb := NewDb()
	createdTodo, _ := memoryDb.CreateTodo("This will be deleted")
	lengthBefore := len(memoryDb.todos)
	err := memoryDb.DeleteTodo(createdTodo.Id)
	lengthAfter := len(memoryDb.todos)
	if this.So(err, should.BeNil) {
		this.So(lengthAfter, should.Equal, lengthBefore-1)
	}
}

func (this *MemoryDbFixture) TestUpdateTodoWhenTodoExists() {
	memoryDb := NewDb()
	createdTodo, _ := memoryDb.CreateTodo("This will be updated")
	updatesToMake := &db.TodoUpdateInput{
		Label:     "Updated Label",
		Completed: true,
	}
	updatedTodo, err := memoryDb.UpdateTodo(createdTodo.Id, updatesToMake)
	if this.So(err, should.BeNil) {
		this.So(updatedTodo.Label, should.Equal, updatesToMake.Label)
		this.So(updatedTodo.Completed, should.Equal, updatesToMake.Completed)
	}
}

func (this *MemoryDbFixture) TestUpdateTodoWhenTodoDoesNotExist() {
	memoryDb := NewDb()
	updatesToMake := &db.TodoUpdateInput{
		Label:     "Updated Label",
		Completed: true,
	}
	updatedTodo, err := memoryDb.UpdateTodo("id that doesn't exist", updatesToMake)
	this.So(err, should.BeNil)
	this.So(updatedTodo, should.BeNil)
}

func (this *MemoryDbFixture) TestNextIdDifferentFromPreviousId() {
	memoryDb := NewDb()
	id1 := memoryDb.nextId()
	id2 := memoryDb.nextId()
	this.So(id1, should.NotEqual, id2)
}
