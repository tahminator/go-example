package graph

import (
	"log"

	"github.com/tahminator/go-example/database"
	repository "github.com/tahminator/go-example/database/service/todo"
)

func (r *Resolver) databaseBuilder() {
	if r.db == nil {
		err := database.Connect()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db, err := database.GetPool()
		if err != nil {
			log.Fatalf("Failed to get database pool: %v", err)
		}
		r.db = db
	}
}

func (r *Resolver) repositoryBuilder() {
	if r.todoRepository == nil {
		r.todoRepository = repository.NewPostgresTodoRepository(r.db)
	}
}
