package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/chrisbroome/go-todo-api/db"
	"github.com/chrisbroome/go-todo-api/db/memory"
	"github.com/chrisbroome/go-todo-api/marshal"
	"github.com/chrisbroome/go-todo-api/todos"
)

func todoRoutes(r chi.Router, db db.Db) chi.Router {
	jsonMarshaller := &marshal.JSONMarshaller{}
	requestParser := &todos.UnmarshallingRequestParser{Unmarshaller: jsonMarshaller}
	r.Method("POST", "/", todos.NewBasicApiHandler(todos.NewCreateHandler(db, requestParser)))
	r.Method("GET", "/{id}", todos.NewBasicApiHandler(todos.NewGetHandler(db)))
	return r
}

type HttpApiApplication struct {
	port         int
	db           db.Db
	marshaller   marshal.HttpMarshaller
	unmarshaller marshal.Unmarshaller
	router       *chi.Mux
}

func NewHttpApiApplication(port int) *HttpApiApplication {
	jsonMarshaller := &marshal.JSONMarshaller{}
	db := memory.NewDb()
	return &HttpApiApplication{
		port:         port,
		db:           db,
		router:       chi.NewRouter(),
		marshaller:   jsonMarshaller,
		unmarshaller: jsonMarshaller,
	}
}

func (app *HttpApiApplication) Configure() *HttpApiApplication {
	router := app.router
	db := app.db
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// timeout after 60 seconds
	router.Use(middleware.Timeout(60 * time.Second))

	// root route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	router.Route("/todos", func(r chi.Router) {
		todoRoutes(r, db)
	})

	router.MethodNotAllowed(router.MethodNotAllowedHandler())
	router.NotFound(router.NotFoundHandler())

	return app
}

func (app *HttpApiApplication) Run() error {
	chi.Walk(app.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%10v %v\n", method, route)
		return nil
	})
	fmt.Printf("Listening on port %v\n", app.port)
	return http.ListenAndServe(fmt.Sprintf(":%v", app.port), app.router)
}

func main() {
	NewHttpApiApplication(3000).Configure().Run()
}
