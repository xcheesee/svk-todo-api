package models

type Todo struct {
	Id          int    `json:"id"`
	UsuarioId   int    `json:"usuarioId"`
	CategoriaId int    `json:"categoriaId"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Created_at  string `json:"created_at"`
}
