package main

import (
	"embed"
	_ "embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tahminator/go-example/graph"
	"github.com/tahminator/go-example/utils"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

//go:embed static/*
var content embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	utils.ValidateEnv([]string{"DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME", "DATABASE_USER", "DATABASE_PASSWORD", "ENV", "ALLOWED_ORIGINS"})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowCredentials: true,
		Debug:            os.Getenv("ENV") == "production",
	})

	http.Handle("/query", corsConfig.Handler(srv))

	if os.Getenv("ENV") != "production" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	} else {
		subFS, err := fs.Sub(content, "static")
		if err != nil {
			panic(err)
		}

		fs := http.FileServer(http.FS(subFS))

		http.Handle("/", fs)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
