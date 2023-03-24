package main

import (
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	defaultPort   = "8080"
	maxComplexity = 100
)

func init() {
	sqlite.Init()
}

func main() {
	// ent client
	client, err := ent.Open(dialect.SQLite, "file:agricoladb.sqlite?cache=shared")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))
	srv.Use(extension.FixedComplexityLimit(maxComplexity))

	http.Handle("/", playground.Handler("AgricolaDB playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
