package types

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTodoDTO struct {
	Description string `json:"description"`
}

type UpdateTodoDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoRepository interface {
	Create(CreateTodoDTO) (*Todo, error)
	ReadByID(string) (*Todo, error)
	Update(UpdateTodoDTO) (*Todo, error)
	List() (*[]Todo, error)
	Delete(string) error
}
