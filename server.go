package main

import (
	"log"
	"net/http"
	"os"
	database "github.com/sumana2001/hackernews/internal/pkg/db/mysql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sumana2001/hackernews/graph"
	"github.com/sumana2001/hackernews/graph/generated"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
