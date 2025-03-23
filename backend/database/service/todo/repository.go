package repository

import (
	"context"

	"github.com/tahminator/go-example/graph/model"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.NewTodo) (*model.Todo, error)
	FindTodos(ctx context.Context) ([]*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.InputTodo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, todoId string) (*model.Todo, error)
}
