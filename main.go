package main

import (
	//"encoding/json"
	//"fmt"
	"net/http"
	"svk-todo-api/controllers"
	//"net/http"
)

func main() {
	todoController := controllers.TController

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		todoController.All(w, req)
	})
	http.HandleFunc("/todo/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == "POST" {
			todoController.Add(w, req)
		} else if req.Method == "GET" {
			todoController.Get(w, req)
		} else if req.Method == "DELETE" {
			todoController.Del(w, req)
		} else if req.Method == "PUT" {
			todoController.Edit(w, req)
		}
	})
	http.ListenAndServe(":8080", nil)
}
