package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"svk-todo-api/pkg/db"
	pkghttp "svk-todo-api/pkg/http"
)

func main() {
	db, err := db.NewSqlConn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	h := pkghttp.NewHandler(db)
	mux := http.NewServeMux()

	f, err := os.OpenFile("logs", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		h.AllTodos(w, req)
	})
	mux.HandleFunc("/todo", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodPost {
			h.AddTodo(w, req)
		} else if req.Method == http.MethodGet {
			h.GetTodo(w, req)
		} else if req.Method == http.MethodDelete {
			h.DelTodo(w, req)
		} else if req.Method == http.MethodPut {
			h.EditTodo(w, req)
		}
	})
	http.ListenAndServe(":8080", mux)
}
