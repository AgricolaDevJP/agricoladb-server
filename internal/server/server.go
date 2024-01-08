package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

type Server struct {
	AllowedOrigins []string
	Client         *ent.Client
	MaxComplexity  int
}

func (s *Server) Install(r *chi.Mux) {
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowCredentials: true,
	}).Handler)

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
