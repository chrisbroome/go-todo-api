package todos

import (
	"errors"
	"net/http"

	"github.com/chrisbroome/go-todo-api/marshal"
)

type ApiHandler interface {
	Handle(r *http.Request) *ApiResponse
}

type ApiResponse struct {
	statusCode int
	body       interface{}
	err        error
}

func NewApiResponse(body interface{}) *ApiResponse {
	return &ApiResponse{
		statusCode: http.StatusOK,
		body:       body,
		err:        nil,
	}
}

func (res *ApiResponse) WithStatusCode(statusCode int) *ApiResponse {
	return &ApiResponse{
		statusCode: statusCode,
		body:       res.body,
		err:        res.err,
	}
}

func (res *ApiResponse) WithErr(err error) *ApiResponse {
	return &ApiResponse{
		statusCode: res.statusCode,
		body:       res.body,
		err:        err,
	}
}


var internalServerError = []byte(`{"error": "internal server error"}`)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"error"`
	Message    string `json:"message"`
}

func NewErrorResponse(err error) *ApiResponse {
	code := http.StatusInternalServerError
	return &ApiResponse {
		statusCode: code,
		err: err,
		body: &ErrorResponse{
			StatusCode: code,
			Err: err,
			Message: err.Error(),
		},
	}
}

type RequestParser interface {
	ParseRequest(r *http.Request, dest interface{}) error
}

type BasicApiHandler struct {
	handler ApiHandler
	marshaller marshal.HttpMarshaller
	unmarshaller marshal.Unmarshaller
}

func NewBasicApiHandler(handler ApiHandler) *BasicApiHandler {
	jsonMarshaller := &marshal.JSONMarshaller{}
	return &BasicApiHandler{handler: handler, marshaller: jsonMarshaller, unmarshaller: jsonMarshaller}
}

func (this *BasicApiHandler) WithMarshaller(marshaller marshal.HttpMarshaller) *BasicApiHandler {
	r := NewBasicApiHandler(this.handler)
	r.marshaller = marshaller
	return r
}

func (this *BasicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := this.handler.Handle(r)
	w.Header().Set("content-type", this.marshaller.ContentType())
	if res.err != nil {
		w.WriteHeader(res.statusCode)
		_ = this.marshaller.Marshal(w, NewErrorResponse(res.err).WithStatusCode(res.statusCode).body)
		return
	}

	if res.body == nil {
		status := http.StatusNotFound
		w.WriteHeader(status)
		_ = this.marshaller.Marshal(w, NewErrorResponse(errors.New("not found")).WithStatusCode(status).body)
		return
	}

	w.WriteHeader(res.statusCode)
	_ = this.marshaller.Marshal(w, res.body)
}

type UnmarshallingRequestParser struct {
	Unmarshaller marshal.Unmarshaller
}

func (this *UnmarshallingRequestParser) ParseRequest(r *http.Request, value interface{}) error {
	return this.Unmarshaller.Unmarshal(r.Body, value)
}
