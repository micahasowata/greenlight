package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spobly/greenlight/internal/config"
	"github.com/spobly/greenlight/internal/data"
)

const version = "1.0.0"

type application struct {
	config config.Config
	logger *log.Logger
}

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "development|staging|production")
	flag.StringVar(&cfg.Db.DSN, "db-dsn", os.Getenv("DB_DSN"), "PostgresQL dsn")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := data.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	logger.Println("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
