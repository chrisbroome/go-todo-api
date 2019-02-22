package todos

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/chrisbroome/go-todo-api/db"
)

type RemoveHandler struct {
	db db.Db
}

func NewRemoveHandler(db db.Db) *RemoveHandler {
	return &RemoveHandler{db: db}
}

func (this *RemoveHandler) Handle(r *http.Request) *ApiResponse {
	id := chi.URLParam(r, "id")
	found, err := this.db.DeleteTodo(id)
	if err != nil {
		return NewErrorApiResponse(err)
	}

	status := http.StatusNoContent
	if !found {
		status = http.StatusNotFound
	}
	return NewApiResponse(nil).WithStatusCode(status)
}
