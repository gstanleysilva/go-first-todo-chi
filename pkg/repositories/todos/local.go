package repository

import (
	"errors"
	"myproject/todo/pkg/types"

	"github.com/google/uuid"
)

type LocalTodoRepository struct {
	Todos []types.Todo
}

func NewTodoRepository() *LocalTodoRepository {
	todos := make([]types.Todo, 0)
	return &LocalTodoRepository{
		Todos: todos,
	}
}

func (t *LocalTodoRepository) Create(data types.CreateTodoDTO) (*types.Todo, error) {
	if data.Description == "" {
		return nil, errors.New("description was not provided")
	}

	id := uuid.New()

	todo := &types.Todo{
		ID:          id.String(),
		Description: data.Description,
		Completed:   false,
	}

	t.Todos = append(t.Todos, *todo)

	return todo, nil
}

func (t *LocalTodoRepository) ReadByID(id string) (*types.Todo, error) {
	for _, todo := range t.Todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (t *LocalTodoRepository) Update(data types.UpdateTodoDTO) (*types.Todo, error) {
	for index := range t.Todos {
		if t.Todos[index].ID == data.ID {
			t.Todos[index].Description = data.Description
			t.Todos[index].Completed = data.Completed
			return &t.Todos[index], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (t *LocalTodoRepository) List() (*[]types.Todo, error) {
	return &t.Todos, nil
}

func (t *LocalTodoRepository) Delete(id string) error {
	indexToRemove := -1

	for index, todo := range t.Todos {
		if todo.ID == id {
			indexToRemove = index
		}
	}

	if indexToRemove == -1 {
		return errors.New("todo not found")
	}

	t.Todos = append(t.Todos[:indexToRemove], t.Todos[indexToRemove+1:]...)

	return nil
}
