package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/internal/server"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
	"github.com/caarlos0/env/v10"
	slogenv "github.com/cbrewster/slog-env"
	"github.com/go-chi/chi/v5"
	"github.com/phsym/console-slog"
)

func init() {
	sqlite.Init()
}

type config struct {
	Port           string   `env:"PORT" envDefault:"8000"`
	DBPath         string   `env:"DB_PATH" envDefault:"agricoladb.sqlite"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:","`
	MaxComplexity  int      `env:"MAX_COMPLEXITY" envDefault:"100"`
	GoEnv          string   `env:"GO_ENV" envDefault:"development"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	var logger *slog.Logger
	if cfg.GoEnv == "production" {
		logger = slog.New(
			// refer to GO_LOG env
			slogenv.NewHandler(slog.NewJSONHandler(os.Stderr, nil)),
		)
	} else {
		logger = slog.New(
			// refer to GO_LOG env
			slogenv.NewHandler(console.NewHandler(os.Stderr, nil)),
		)
	}
	slog.SetDefault(logger)

	dsn := fmt.Sprintf("file:%s?cache=shared&mode=ro", cfg.DBPath)
	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		slog.Error("failed opening connection to sqlite", err)
		os.Exit(1)
	}
	defer client.Close()

	router := chi.NewRouter()
	s := &server.Server{
		AllowedOrigins: cfg.AllowedOrigins,
		Client:         client,
		MaxComplexity:  cfg.MaxComplexity,
	}
	s.Install(router)

	slog.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port))
	if err := http.ListenAndServe("localhost:"+cfg.Port, router); err != nil {
		slog.Error("failed starting server", err)
		os.Exit(1)
	}
}
