package storage

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyInList = errors.New("todo already in list")

type DB interface {
	SaveTodo(ctx context.Context, name, description string) (Todo, error)
	GetTodoList(ctx context.Context) ([]Todo, error)
	GetTodoByName(ctx context.Context, name string) (Todo, error)
	EditTodo(ctx context.Context, id string, todo Todo) (Todo, error)
	DeleteTodo(ctx context.Context, id string) error
	ClearTodoList(ctx context.Context) error
}

type Todo struct {
	ID          string
	Name        string
	Description string
}
