package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/rs/cors"
)

type Server struct {
	AllowedOrigins []string
	Client         *ent.Client
	MaxComplexity  int
}

func (s *Server) Install(r *chi.Mux) {
	httpLogger := &httplog.Logger{
		Logger: slog.Default(),
		Options: httplog.Options{
			LogLevel:         slog.LevelDebug,
			RequestHeaders:   true,
			MessageFieldName: "message",
			QuietDownRoutes: []string{
				"/robots.txt",
				"/favicon.ico",
			},
			QuietDownPeriod: 10 * time.Second,
		},
	}

	r.Use(httplog.RequestLogger(httpLogger))
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowCredentials: true,
	}).Handler)
	r.Use(middleware.Compress(5))

	r.Handle("/", s.handlerIndex())
	r.Handle("/graphql", s.handleGraphql())
	r.Get("/robots.txt", s.getRobotsTxt())
}

func (s *Server) handlerIndex() http.Handler {
	return playground.Handler("AgricolaDB playground", "/graphql")
}

func (s *Server) handleGraphql() http.Handler {
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: s.Client,
	}})
	srv := handler.NewDefaultServer(schema)
	srv.Use(extension.FixedComplexityLimit(s.MaxComplexity))
	return srv
}

func (s *Server) getRobotsTxt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nDisallow: /\n")) // nolint
	}
}
