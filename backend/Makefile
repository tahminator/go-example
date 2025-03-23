build:
	go build server.go

run:
	./server

dev:
	go run server.go

generate:
	go run github.com/99designs/gqlgen generate

migrate:
	@dotenvx run -- sh -c 'migrate -source file://database/migration -database "pgx5://$${DATABASE_USER}:$${DATABASE_PASSWORD}@$${DATABASE_HOST}:$${DATABASE_PORT}/$${DATABASE_NAME}" -verbose up'

