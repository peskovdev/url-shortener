package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/peskovdev/url-shortener/internal/config"
	mwLogger "github.com/peskovdev/url-shortener/internal/http-server/middleware"
	"github.com/peskovdev/url-shortener/internal/lib/logger/handlers/slogpretty"
	"github.com/peskovdev/url-shortener/internal/lib/logger/sl"
	"github.com/peskovdev/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", string(cfg.Env)), slog.String("version", "1.0"))
	log.Debug("debug message are enabled")
	log.Warn("Some warn")
	log.Error("ALARM")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("error opening db", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}

func setupLogger(env config.Env) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.EnvLocal:
		log = slog.New(
			slogpretty.NewPrettyHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}
	return log
}
