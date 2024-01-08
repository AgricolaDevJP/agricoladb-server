package main

import (
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/internal/server"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi"
)

func init() {
	sqlite.Init()
}

type config struct {
	Port           string   `env:"PORT" envDefault:"8000"`
	DBPath         string   `env:"DB_PATH" envDefault:"agricoladb.sqlite"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:","`
	MaxComplexity  int      `env:"MAX_COMPLEXITY" envDefault:"100"`
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

	router := chi.NewRouter()
	s := &server.Server{
		AllowedOrigins: cfg.AllowedOrigins,
		Client:         client,
		MaxComplexity:  cfg.MaxComplexity,
	}
	s.Install(router)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
