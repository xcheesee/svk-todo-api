package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"svk-todo-api/pkg/domain"
)

type Handler struct {
	Svc domain.TodoSvc
}

func NewHandler(svc domain.TodoSvc) *Handler {
	return &Handler{
		Svc: svc,
	}
}

func (h *Handler) AddTodo(w http.ResponseWriter, req *http.Request) {
	var newTodo domain.Todo
	json.NewDecoder(req.Body).Decode(&newTodo)

	if err := h.Svc.Create(&newTodo); err != nil {
		log.Println(err)
		w.Write([]byte("An error has ocurred"))
		return
	}
	w.Write([]byte("Operation Successful"))
}

func (h *Handler) AllTodos(w http.ResponseWriter, req *http.Request) {
	todos, err := h.Svc.List()
	if err != nil {
		w.Write([]byte("An error has ocurred"))
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) GetTodo(w http.ResponseWriter, req *http.Request) {
	todoQueries := req.URL.Query()
	todoId, err := strconv.Atoi(todoQueries.Get("id"))
	if err != nil {
		w.Write([]byte("Valor de id invalido"))
		return
	}
	todo, err := h.Svc.Get(todoId)
	if err != nil {
		w.Write([]byte("An error has ocurred"))
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) EditTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "todo Put test")
}

func (h *Handler) DelTodo(w http.ResponseWriter, req *http.Request) {
	todoQueries := req.URL.Query()
	todoId, err := strconv.Atoi(todoQueries.Get("id"))
	if err != nil {
		w.Write([]byte("Valor de id invalido"))
		return
	}
	if err = h.Svc.Delete(todoId); err != nil {
		w.Write([]byte("An error has ocurred"))
		return
	}
	w.Write([]byte("Operation Successful"))
}
