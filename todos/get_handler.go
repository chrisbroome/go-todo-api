package todos

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/chrisbroome/go-todo-api/db"
)

type GetHandler struct {
	db db.Db
}

func NewGetHandler(db db.Db) *GetHandler {
	return &GetHandler{db: db}
}

func (this *GetHandler) Handle(r *http.Request) *ApiResponse {
	id := chi.URLParam(r, "id")
	todo, err := this.db.GetTodoById(id)
	if err != nil {
		return NewErrorApiResponse(err)
	}
	return NewApiResponse(toTodoResponse(todo))
}
