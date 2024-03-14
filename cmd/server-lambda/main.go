package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/internal/logger"
	"github.com/AgricolaDevJP/agricoladb-server/internal/server"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5"
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

	logger := logger.NewLogger(&logger.LoggerOption{})
	slog.SetDefault(logger)

	dns := fmt.Sprintf("file:%s?cache=shared&mode=ro", cfg.DBPath)
	client, err := ent.Open(dialect.SQLite, dns)
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

	lambda.Start(httpadapter.NewV2(router).ProxyWithContext)
}
