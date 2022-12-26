package main

import (
	"agricoladb/ent"
	"agricoladb/initdb"
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ent client
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/agricoladb")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		os.Exit(1)
	}

	g := initdb.NewGenerator(client, "./masterdata/", true)
	if err := g.Generate(); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		os.Exit(1)
	}

	log.Println("success")
}
