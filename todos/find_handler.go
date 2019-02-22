package todos

import (
	"net/http"

	"github.com/chrisbroome/go-todo-api/db"
)

type FindHandler struct {
	db db.Db
}

func NewFindHandler(db db.Db) *FindHandler {
	return &FindHandler{db: db}
}

func (this *FindHandler) Handle(r *http.Request) *ApiResponse {
	todos, err := this.db.ListTodos()
	if err != nil {
		return NewErrorApiResponse(err)
	}

	ret := make([]*TodoResponse, len(todos))
	for i, todo := range todos {
		ret[i] = toTodoResponse(todo)
	}

	return NewApiResponse(ret)
}
