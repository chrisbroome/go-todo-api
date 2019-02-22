package todos

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/chrisbroome/go-todo-api/marshal"
)

type BasicApiHandler struct {
	handler      ApiHandler
	marshaller   marshal.HttpMarshaller
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

	if res.statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		_, _ = w.Write([]byte(""))
		return
	}

	w.Header().Set("content-type", this.marshaller.ContentType())
	if res.err != nil {
		w.WriteHeader(res.statusCode)
		_ = this.marshaller.Marshal(w, NewErrorResponse(res.err).WithStatusCode(res.statusCode))
		return
	}

	if res.statusCode == http.StatusNotFound || reflect.ValueOf(res.body).IsNil() {
		status := http.StatusNotFound
		w.WriteHeader(status)
		notFoundRes := NewErrorResponse(errors.New("not found")).WithStatusCode(status)
		_ = this.marshaller.Marshal(w, notFoundRes)
		return
	}

	w.WriteHeader(res.statusCode)
	_ = this.marshaller.Marshal(w, res.body)
}
