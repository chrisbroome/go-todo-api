package todos

import "net/http"

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
