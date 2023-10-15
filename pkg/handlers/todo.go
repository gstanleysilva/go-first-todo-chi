package handlers

import (
	"encoding/json"
	"fmt"
	"myproject/todo/pkg/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	repo types.TodoRepository
}

func NewTodoHandler(repo types.TodoRepository) *TodoHandler {
	return &TodoHandler{
		repo: repo,
	}
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	returnSlice := make([]types.Todo, 0)

	todos, err := h.repo.List()
	if err != nil {
		fmt.Println("failed to list todos:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if todos != nil {
		returnSlice = append(returnSlice, *todos...)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnSlice)
}

func (h *TodoHandler) ReadByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := h.repo.ReadByID(id)
	if err != nil {
		fmt.Println("failed to read by id:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := &types.UpdateTodoDTO{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("failed to decode body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := h.repo.ReadByID(id); err != nil {
		fmt.Println("failed to read by id:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data.ID = id

	todo, err := h.repo.Update(*data)
	if err != nil {
		fmt.Println("failed to update todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	todo := &types.CreateTodoDTO{}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println("failed to decode body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdTodo, err := h.repo.Create(*todo)
	if err != nil {
		fmt.Println("failed to create:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTodo)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, err := h.repo.ReadByID(id); err != nil {
		fmt.Println("failed to read by id:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.repo.Delete(id)
	if err != nil {
		fmt.Println("failed to remove todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
