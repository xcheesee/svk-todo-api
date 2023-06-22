package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"svk-todo-api/models"
)

var TController TodoController

func (t TodoController) Add(w http.ResponseWriter, req *http.Request) {
	todos := Response{
		Todos: []models.Todo{
			{
				Id:          1,
				CategoriaId: 1,
				UsuarioId:   2,
				Titulo:      "teste",
				Descricao:   "Teste",
				Data:        "21/06/2023",
			},
			{
				Id:          2,
				CategoriaId: 2,
				UsuarioId:   3,
				Titulo:      "pog",
				Descricao:   "champ",
				Data:        "21/06/2023",
			},
		},
	}
	var newTodo models.Todo
	json.NewDecoder(req.Body).Decode(&newTodo)
	todos.Todos = append(todos.Todos, newTodo)
	json.NewEncoder(w).Encode(todos)
}

func (t TodoController) All(w http.ResponseWriter, req *http.Request) {
	todos := Response{
		Todos: []models.Todo{
			{
				Id:          1,
				CategoriaId: 1,
				UsuarioId:   2,
				Titulo:      "teste",
				Descricao:   "Teste",
				Data:        "21/06/2023",
			},
			{
				Id:          2,
				CategoriaId: 2,
				UsuarioId:   3,
				Titulo:      "pog",
				Descricao:   "champ",
				Data:        "21/06/2023",
			},
		},
	}
	data, err := json.Marshal(todos)
	if err != nil {
		panic("encode error")
	}
	w.Write(data)
}

func (t TodoController) Get(w http.ResponseWriter, req *http.Request) {
	todoQueries := req.URL.Query()
	todoId, err := strconv.Atoi(todoQueries.Get("id"))
	if err != nil {
		w.Write([]byte("Valor de id invalido"))
		return
	}
	todos := Response{
		Todos: []models.Todo{
			{
				Id:          1,
				CategoriaId: 1,
				UsuarioId:   2,
				Titulo:      "teste",
				Descricao:   "Teste",
				Data:        "21/06/2023",
			},
			{
				Id:          2,
				CategoriaId: 2,
				UsuarioId:   3,
				Titulo:      "pog",
				Descricao:   "champ",
				Data:        "21/06/2023",
			},
		},
	}
	for _, todo := range todos.Todos {
		if todoId == todo.Id {
			json.NewEncoder(w).Encode(todo)
		}
	}
	w.Write([]byte("Todo nao encontrado"))
}

func (t TodoController) Edit(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "todo Put test")
}

func (t TodoController) Del(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "todo Del test")
}
