package main

import (
	"agricoladb/ent"
	"agricoladb/initdb"
	"context"
	"log"
	"net"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/go-sql-driver/mysql"
)

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

func (c *config) getDSN() string {
	mc := mysql.NewConfig()
	mc.DBName = c.DBName
	mc.User = c.DBUser
	mc.Passwd = c.DBPassword
	mc.Addr = net.JoinHostPort(c.DBHost, c.DBPort)
	mc.Net = "tcp"
	return mc.FormatDSN()
}
