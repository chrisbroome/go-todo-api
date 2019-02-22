package todos

import (
	"net/http"

	"github.com/chrisbroome/go-todo-api/marshal"
)

type ApiHandler interface {
	Handle(r *http.Request) *ApiResponse
}

type RequestParser interface {
	ParseRequest(r *http.Request, dest interface{}) error
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	err        error  `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Error:      err.Error(),
		err:        err,
	}
}

func (e *ErrorResponse) WithStatusCode(statusCode int) *ErrorResponse {
	r := NewErrorResponse(e.err)
	r.StatusCode = statusCode
	return r
}

func NewErrorApiResponse(err error) *ApiResponse {
	code := http.StatusInternalServerError
	return &ApiResponse{
		statusCode: code,
		err:        err,
		body: &ErrorResponse{
			StatusCode: code,
			err:        err,
			Error:      err.Error(),
		},
	}
}

type UnmarshallingRequestParser struct {
	Unmarshaller marshal.Unmarshaller
}

func (this *UnmarshallingRequestParser) ParseRequest(r *http.Request, value interface{}) error {
	return this.Unmarshaller.Unmarshal(r.Body, value)
}
