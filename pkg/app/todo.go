package app

import (
	"svk-todo-api/pkg/domain"
)

type todoSvc struct {
	DB domain.TodoDB
}

func NewTodoSvc(db domain.TodoDB) domain.TodoSvc {
	return todoSvc{
		DB: db,
	}
}

func (s todoSvc) Get(id int) (*domain.Todo, error) {
	return nil, nil
}

func (s todoSvc) List() ([]*domain.Todo, error) {
	return nil, nil
}

func (s todoSvc) Create(t *domain.Todo) error {
	return nil
}

func (s todoSvc) Delete(id int) error {
	return nil
}
