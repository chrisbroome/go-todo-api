package todos

import (
	"net/http"

	"github.com/chrisbroome/go-todo-api/db"
)

type createTodoRequest struct {
	Label string `json:"label"`
}

type CreateHandler struct {
	db     db.Db
	parser RequestParser
}

func NewCreateHandler(db db.Db, parser RequestParser) *CreateHandler {
	return &CreateHandler{db: db, parser: parser}
}

func (this *CreateHandler) Handle(r *http.Request) *ApiResponse {
	req := createTodoRequest{}
	err := this.parser.ParseRequest(r, &req)
	if err != nil {
		return NewErrorResponse(err)
	}

	todo, err := this.db.CreateTodo(req.Label)
	if err != nil {
		return NewErrorResponse(err)
	}

	return NewApiResponse(toTodoResponse(todo)).WithStatusCode(http.StatusCreated)
}
