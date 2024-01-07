package main

import (
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	maxComplexity = 100
)

func init() {
	sqlite.Init()
}

type config struct {
	Port           string   `env:"PORT" envDefault:"8000"`
	DBPath         string   `env:"DB_PATH" envDefault:"agricoladb.sqlite"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:","`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	dsn := fmt.Sprintf("file:%s?cache=shared&mode=ro", cfg.DBPath)

	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))
	srv.Use(extension.FixedComplexityLimit(maxComplexity))

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: true,
	}).Handler)

	router.Handle("/", playground.Handler("AgricolaDB playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
