package repository

import (
	"TODO"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TODO.User) (int, error)
	GetUser(username, password string) (TODO.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
