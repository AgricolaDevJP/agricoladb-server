package main

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
)

func init() {
	sqlite.Init()
}

type config struct {
	Port           string   `env:"PORT" envDefault:"8000"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:","`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	client, err := ent.Open(dialect.SQLite, "file:/agricoladb.sqlite?cache=shared&mode=ro")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}})
	server := handler.NewDefaultServer(schema)

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: true,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", server)
	lambda.Start(httpadapter.NewV2(router).ProxyWithContext)
}
