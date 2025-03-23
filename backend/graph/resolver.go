package graph

import (
	"github.com/jackc/pgx/v5/pgxpool"
	repository "github.com/tahminator/go-example/database/service/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db             *pgxpool.Pool
	todoRepository *repository.PostgresTodoRepository
}
