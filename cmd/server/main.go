package main

import (
	"agricoladb/ent"
	"agricoladb/graph"
	"agricoladb/graph/generated"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/caarlos0/env/v6"
	"github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

type config struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"agricoladb"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	// ent client
	client, err := ent.Open("mysql", cfg.getDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))

	http.Handle("/", playground.Handler("AgricolaDB playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (c *config) getDSN() string {
	mc := mysql.NewConfig()
	mc.DBName = c.DBName
	mc.User = c.DBUser
	mc.Passwd = c.DBPassword
	mc.Addr = net.JoinHostPort(c.DBHost, c.DBPort)
	mc.Net = "tcp"
	return mc.FormatDSN()
}
