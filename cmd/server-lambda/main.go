package main

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"

	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/internal/server"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
)

func init() {
	sqlite.Init()
}

type config struct {
	Port           string   `env:"PORT" envDefault:"8000"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:","`
	MaxComplexity  int      `env:"MAX_COMPLEXITY" envDefault:"100"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	dns := "file:/agricoladb.sqlite?cache=shared&mode=ro"
	client, err := ent.Open(dialect.SQLite, dns)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	router := chi.NewRouter()
	s := &server.Server{
		AllowedOrigins: cfg.AllowedOrigins,
		Client:         client,
		MaxComplexity:  cfg.MaxComplexity,
	}
	s.Install(router)

	lambda.Start(httpadapter.NewV2(router).ProxyWithContext)
}
