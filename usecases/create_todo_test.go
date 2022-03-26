package usecases

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"

	"github.com/chrisbroome/go-todo-api/db/memory"
)

func TestCreateTodoInteractorFixture(t *testing.T) {
	gunit.Run(new(CreateTodoInteractorFixture), t)
}

type CreateTodoInteractorFixture struct {
	*gunit.Fixture
	interactor *CreateTodoInteractor
}

func (this *CreateTodoInteractorFixture) Setup() {
	this.interactor = NewCreateTodoInteractor(memory.NewDb())
}

func (this *CreateTodoInteractorFixture) TestWhenGivenALabelShouldCreateTheTodo() {
	label := "my todo"
	todo, err := this.interactor.Execute(label)
	if this.So(err, should.BeNil) {
		this.So(todo.Label, should.Equal, label)
	}
}
