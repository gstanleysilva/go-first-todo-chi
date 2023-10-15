package main

import (
	"fmt"
	"myproject/todo/pkg/configs"
	"myproject/todo/pkg/handlers"
	"myproject/todo/pkg/types"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config   *configs.Config
	TodoRepo *types.TodoRepository
}

func NewApplication(config configs.Config, repo types.TodoRepository) *Application {
	return &Application{
		Config:   &config,
		TodoRepo: &repo,
	}
}

func (a *Application) StartServer() {

	chi := chi.NewRouter()
	chi.Use(middleware.Logger)

	chi.Route("/todos", a.TodosRoutes)

	server := http.Server{
		Addr:    a.Config.Port,
		Handler: chi,
	}

	fmt.Println("Server running on port", a.Config.Port)

	server.ListenAndServe()
}

func (a *Application) TodosRoutes(r chi.Router) {

	handler := handlers.NewTodoHandler(*a.TodoRepo)

	r.Get("/", handler.List)
	r.Post("/", handler.Create)
	r.Delete("/{id}", handler.Delete)
	r.Put("/{id}", handler.Update)
	r.Get("/{id}", handler.ReadByID)
}
