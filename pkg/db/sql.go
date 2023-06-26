package db

import (
	"database/sql"
	"log"
	"svk-todo-api/pkg/domain"

	"github.com/go-sql-driver/mysql"
)

type sqlConn struct {
	Client *sql.DB
}

func NewSqlConn() (domain.TodoDB, error) {
	var db *sql.DB
	var err error

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "svk_todo",
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return sqlConn{Client: db}, nil
}

func (c sqlConn) Get(id int) (*domain.Todo, error) {
	var todo domain.Todo
	todoId := id
	rows, err := c.Client.Query("SELECT * FROM todos where id=?", todoId)
	if err != nil {
		return nil, err
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
	return &todo, nil
}

func (c sqlConn) List() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	rows, err := c.Client.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(
			&todo.Id,
			&todo.UsuarioId,
			&todo.CategoriaId,
			&todo.Titulo,
			&todo.Descricao,
			&todo.Created_at,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (c sqlConn) Create(t *domain.Todo) error {
	_, err := c.Client.Exec(
		"INSERT INTO todos(usuarioId, categoriaId, titulo, descricao, created_at) VALUES(?, ?, ?, ?, ?)",
		t.UsuarioId, t.CategoriaId, t.Titulo, t.Descricao, t.Created_at,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c sqlConn) Delete(id int) error {
	_, err := c.Client.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
