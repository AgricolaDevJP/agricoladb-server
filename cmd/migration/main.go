package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/AgricolaDevJP/agricoladb-server/ent"
	"github.com/AgricolaDevJP/agricoladb-server/initdb"
	"github.com/AgricolaDevJP/agricoladb-server/sqlite"
)

func init() {
	sqlite.Init()
}

func main() {
	// ent client
	client, err := ent.Open(dialect.SQLite, "file:agricoladb.sqlite?cache=shared")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	g := initdb.NewGenerator(client, true)

	if err := g.DropTablesIfNeed(); err != nil {
		log.Fatalf("failed creating schema resources1: %v", err)
	}

	// auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources:2 %v", err)
	}

	if err := g.Generate(); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("success")
}
