package todos

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/entities"
	"github.com/chrisbroome/go-todo-api/marshal"
)

type ApiResponseWriter struct {
	writer    http.ResponseWriter
	marshaler marshal.Marshaller
}

func NewApiResponseWriter(writer http.ResponseWriter, marshaler marshal.Marshaller) *ApiResponseWriter {
	return &ApiResponseWriter{
		writer:    writer,
		marshaler: marshaler,
	}
}

func (this *ApiResponseWriter) Header() http.Header {
	return this.writer.Header()
}

func (this *ApiResponseWriter) Write(bytes []byte) (int, error) {
	return this.writer.Write(bytes)
}

func (this *ApiResponseWriter) WriteHeader(statusCode int) {
	this.writer.WriteHeader(statusCode)
}

func (this *ApiResponseWriter) WriteError(err error) {
	res := ErrorResponse{
		Err:     err,
		Message: err.Error(),
	}

	marshalErr := this.marshaler.Marshal(this.writer, res)
	if marshalErr != nil {

	}
}

func (this *ApiResponseWriter) WriteResponse(value interface{}) {
	this.marshaler.Marshal(this.writer, value)
}

var internalServerError = []byte(`{"error": "internal server error"}`)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"error"`
	Message    string `json:"message"`
}

type TodoResponse struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	Completed bool   `json:"completed"`
}

type CreateTodoHandler struct {
	Db db.Db
	marshal.Unmarshaller
	marshal.Marshaller
}

type createTodoRequest struct {
	label string
}

func (h *CreateTodoHandler) ParseRequest(r *http.Request) (string, *ErrorResponse) {
	req := createTodoRequest{}
	err := h.Unmarshaller.Unmarshal(r.Body, &req)
	if err != nil {
		return "", &ErrorResponse{
			Err:        err,
			Message:    "Error parsing body",
			StatusCode: http.StatusBadRequest,
		}
	}

	return req.label, nil
}

func (h *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	label, errRes := h.ParseRequest(r)
	if errRes != nil {
		w.WriteHeader(errRes.StatusCode)
		err := h.Marshal(w, errRes)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		return
	}
	todo, err := h.Db.CreateTodo(label)
	if err != nil {
		h.writeError(w, &ErrorResponse{})
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.writeResponse(w, toTodoResponse(todo), http.StatusCreated)
}

func (h *CreateTodoHandler) writeResponse(w http.ResponseWriter, value *TodoResponse, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	h.write(w, value, statusCode)
}

func (h *CreateTodoHandler) writeError(w http.ResponseWriter, value *ErrorResponse) {
	h.write(w, value, value.StatusCode)
}

func (h *CreateTodoHandler) write(w http.ResponseWriter, value interface{}, statusCode int) {
	writer := NewApiResponseWriter(w, h.Marshaller)
	writer.WriteHeader(statusCode)
	writer.WriteResponse(value)
}

func toTodoResponse(todo *entities.Todo) *TodoResponse {
	return &TodoResponse{
		Id:        todo.Id,
		Label:     todo.Label,
		Completed: todo.Completed,
	}
}

type GetTodoHandler struct {
	Db db.Db
	marshal.Marshaller
}

func (h *GetTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo, err := h.Db.GetTodoById(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := TodoResponse{
		Id:        todo.Id,
		Label:     todo.Label,
		Completed: todo.Completed,
	}

	h.write(w, res, http.StatusOK)
}

func (h *GetTodoHandler) write(w http.ResponseWriter, value interface{}, statusCode int) {
	writer := NewApiResponseWriter(w, h.Marshaller)
	writer.WriteHeader(statusCode)
	writer.WriteResponse(value)
}
