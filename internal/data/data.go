package data

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/spobly/greenlight/internal/config"
)

func OpenDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Db.MaxOpenConns)

	db.SetMaxIdleConns(cfg.Db.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
