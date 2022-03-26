package random_test

import (
	"testing"

	"github.com/chrisbroome/go-todo-api/idgenerator/random"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestIdGenerator_GenerateID(t *testing.T) {
	assert := assertions.New(t)
	idg := random.NewIdGenerator()
	id := idg.GenerateID()
	assert.So(id, should.NotBeBlank)
}
