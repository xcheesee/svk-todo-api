package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"svk-todo-api/models"
)

var TController TodoController

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (h *BaseHandler) AddTodo(w http.ResponseWriter, req *http.Request) {
	todos := Response{
		Todos: []models.Todo{
			{
				Id:          1,
				CategoriaId: 1,
				UsuarioId:   2,
				Titulo:      "teste",
				Descricao:   "Teste",
				Created_at:  "21/06/2023",
			},
			{
				Id:          2,
				CategoriaId: 2,
				UsuarioId:   3,
				Titulo:      "pog",
				Descricao:   "champ",
				Created_at:  "21/06/2023",
			},
		},
	}
	var newTodo models.Todo
	json.NewDecoder(req.Body).Decode(&newTodo)
	todos.Todos = append(todos.Todos, newTodo)
	json.NewEncoder(w).Encode(todos)
}

func (h *BaseHandler) AllTodos(w http.ResponseWriter, req *http.Request, db *sql.DB) {
	var todos []models.Todo
	rows, err := h.db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.Id, &todo.UsuarioId, &todo.CategoriaId, &todo.Titulo, &todo.Descricao, &todo.Created_at); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *BaseHandler) GetTodo(w http.ResponseWriter, req *http.Request) {
	var todo models.Todo
	todoQueries := req.URL.Query()
	todoId, err := strconv.Atoi(todoQueries.Get("id"))
	if err != nil {
		w.Write([]byte("Valor de id invalido"))
		return
	}
	rows, err := h.db.Query("SELECT * FROM todos where id=?", todoId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&todo.Id,
			&todo.UsuarioId,
			&todo.CategoriaId,
			&todo.Titulo,
			&todo.Descricao,
			&todo.Created_at,
		); err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *BaseHandler) EditTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "todo Put test")
}

func (h *BaseHandler) DelTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "todo Del test")
}
