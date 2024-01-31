package main

import (
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spobly/greenlight/internal/config"
	"github.com/spobly/greenlight/internal/data"
	"github.com/spobly/greenlight/internal/mailer"
)

const version = "1.0.0"

type application struct {
	config config.Config
	logger *slog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	cfg := &config.Config{}

	displayVersion := flag.Bool("version", false, "Display version and exit")
	cfg.Parse()

	if *displayVersion {
		fmt.Printf("Version: \t%s\n", version)
		os.Exit(1)
	}

	db, err := data.OpenDB(*cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	expvar.NewString("version").Set(version)

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	app := &application{
		config: *cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.SMTP.Port, cfg.SMTP.Host, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Sender),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
	}
}
