package domain

type Todo struct {
	Id          int    `json:"id"`
	UsuarioId   int    `json:"usuarioId"`
	CategoriaId int    `json:"categoriaId"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Created_at  string `json:"created_at"`
}

type Usuario struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type Categoria struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type TodoSvc interface {
	Get(id int) (*Todo, error)
	List() ([]*Todo, error)
	Create(t *Todo) error
	Delete(id int) error
}

type TodoDB interface {
	Get(id int) (*Todo, error)
	List() ([]*Todo, error)
	Create(t *Todo) error
	Delete(id int) error
}
