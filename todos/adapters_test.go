package todos

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestApiResponseWriterAdapterFixture(t *testing.T) {
	gunit.Run(new(ApiResponseWriterAdapterFixture), t)
}

type ApiResponseWriterAdapterFixture struct {
	*gunit.Fixture
}

func (this *ApiResponseWriterAdapterFixture) Setup() {
}

func (this *ApiResponseWriterAdapterFixture) Test() {

}
