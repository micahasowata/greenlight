package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spobly/greenlight/internal/config"
	"github.com/spobly/greenlight/internal/data"
)

const version = "1.0.0"

type application struct {
	config config.Config
	logger *slog.Logger
	models data.Models
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "development|staging|production")
	flag.StringVar(&cfg.Db.DSN, "db-dsn", os.Getenv("DB_DSN"), "PostgresQL dsn")
	flag.IntVar(&cfg.Db.MaxOpenConns, "db-max-open-conns", 25, "PostgresQL max open connections")
	flag.IntVar(&cfg.Db.MaxIdleConns, "db-max-idle-conns", 25, "PostgresQL max idle connections")
	flag.StringVar(&cfg.Db.MaxIdleTime, "db-max-idle-time", "15m", "PostgresQL max connection idle time")
	flag.Parse()

	db, err := data.OpenDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("starting server", slog.Group("properties", slog.String("env", cfg.Env), slog.String("addr", srv.Addr)))

	err = srv.ListenAndServe()
	logger.Error(err.Error())
}
