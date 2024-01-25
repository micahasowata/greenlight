package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutDownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.Info("shutting down server", slog.Group("properties", slog.String("signal", s.String())))

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		shutDownError <- srv.Shutdown(ctx)
	}()

	app.logger.Info("starting server", slog.Group("properties", slog.String("env", app.config.Env), slog.String("addr", srv.Addr)))

	err := srv.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	err = <-shutDownError
	if err != nil {
		return err
	}

	app.logger.Info("stopped server", slog.Group("properties", slog.String("addr", srv.Addr)))
	return nil
}
