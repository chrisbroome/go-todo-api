package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/marshal"
	"github.com/chrisbroome/go-todo-api/todos"
)

func todoRoutes(r chi.Router, db db.Db) chi.Router {
	jsonMarshaller := &marshal.JSONMarshaller{}
	createTodo := &todos.CreateTodoHandler{
		Db:           db,
		Unmarshaller: jsonMarshaller,
		Marshaller:   jsonMarshaller,
	}

	getTodo := &todos.GetTodoHandler{
		Db:         db,
		Marshaller: jsonMarshaller,
	}

	r.Route("/todos", func(r chi.Router) {
		r.Method("POST", "/", createTodo)
		r.Method("GET", "/{id}", getTodo)
	})

	return r
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// timeout after 60 seconds
	r.Use(middleware.Timeout(60 * time.Second))

	// root route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	r.Route("/todos", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {

		})
	})

	r.MethodNotAllowed(r.MethodNotAllowedHandler())
	r.NotFound(r.NotFoundHandler())

	http.ListenAndServe(":3000", r)
}
