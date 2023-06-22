package controllers

import "svk-todo-api/models"

type Response struct {
	Todos []models.Todo `json:"todos"`
}

type Controller interface {
	All()
	Get()
	Add()
	Edit()
	Del()
}

type TodoController struct{}
