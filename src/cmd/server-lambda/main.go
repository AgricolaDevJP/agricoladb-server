package main

import (
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/graph"
	"github.com/AgricolaDevJP/agricoladb-server/graph/generated"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
)

func init() {
	sqlite.Init()
}

func main() {
	client, err := ent.Open(dialect.SQLite, "file:/agricoladb.sqlite?cache=shared&mode=ro")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}})
	server := handler.NewDefaultServer(schema)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", server)
	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
