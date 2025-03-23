package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tahminator/go-example/graph/model"
)

type PostgresTodoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTodoRepository(db *pgxpool.Pool) *PostgresTodoRepository {
	return &PostgresTodoRepository{db: db}
}

func (repo *PostgresTodoRepository) FindTodos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo

	query := `
    SELECT
      id,
      text,
      done,
      "createdAt"
    FROM
      "Todo" t
    ORDER BY "createdAt" ASC
  `

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}

	for rows.Next() {
		todo := &model.Todo{}
		var rawID uuid.UUID

		err := rows.Scan(&rawID, &todo.Text, &todo.Done, &todo.CreatedAt)
		if err != nil {
			continue
		}
		todo.ID = rawID.String()
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *PostgresTodoRepository) CreateTodo(ctx context.Context, todo *model.NewTodo) (*model.Todo, error) {
	newTodo := &model.Todo{}
	query := `
    INSERT INTO "Todo"
      (text)
    VALUES
      ($1)
    RETURNING
      id, done, "createdAt"
  `

	var id uuid.UUID
	err := repo.db.QueryRow(ctx, query, todo.Text).Scan(&id, &newTodo.Done, &newTodo.CreatedAt)
	if err != nil || id.String() == "" {
		return nil, fmt.Errorf("failed to insert todo: %w", err)
	}

	newTodo.ID = id.String()
	newTodo.Text = todo.Text
	return newTodo, nil
}

func (repo *PostgresTodoRepository) UpdateTodo(ctx context.Context, todo *model.InputTodo) (*model.Todo, error) {
	query := `
    UPDATE
      "Todo"
    SET
      text = $1, done = $2
    WHERE
      id = $3
    RETURNING
      id,
      text,
      done,
      "createdAt"
  `

	updatedTodo := &model.Todo{}
	uuid, err := uuid.Parse(todo.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse todo ID into UUID type: %w", err)
	}
	row := repo.db.QueryRow(ctx, query, &todo.Text, &todo.Done, uuid)

	err = row.Scan(&updatedTodo.ID, &updatedTodo.Text, &updatedTodo.Done, &updatedTodo.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create new updated todo object: %w", err)
	}

	return updatedTodo, nil
}

func (repo *PostgresTodoRepository) DeleteTodo(ctx context.Context, todoId string) (*model.Todo, error) {
	query := `
    DELETE FROM
      "Todo"
    WHERE
      id = $1
    RETURNING
      id,
      text,
      done,
      "createdAt"
  `

	deletedTodo := &model.Todo{}
	uuid, err := uuid.Parse(todoId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse todo ID into UUID type: %w", err)
	}

	row := repo.db.QueryRow(ctx, query, &uuid)

	err = row.Scan(&deletedTodo.ID, &deletedTodo.Text, &deletedTodo.Done, &deletedTodo.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create deleted todo object: %w", err)
	}

	return deletedTodo, nil
}

// This enforces that it matches the repository.
var _ TodoRepository = new(PostgresTodoRepository)
