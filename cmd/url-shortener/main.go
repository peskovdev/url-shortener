package main

import (
	"github.com/peskovdev/url-shortener/internal/config"
	"github.com/peskovdev/url-shortener/internal/lib/logger/sl"
	"github.com/peskovdev/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", string(cfg.Env)))
	log.Debug("debug message are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("error opening db", sl.Err(err))
		os.Exit(1)
	}

	// check that saveURL works
	alias := "rick"
	_, err = storage.SaveURL("https://youtu.be/abc", alias)
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}
	urlOriginal, err := storage.GetOriginalURL(alias)
	if err != nil {
		log.Error("failed to get url", sl.Err(err))
		os.Exit(1)
	}
	log.Info(urlOriginal)
	err = storage.DeleteURL(alias)
	if err != nil {
		log.Error("failed to delete url", sl.Err(err))
		os.Exit(1)
	}
	err = storage.DeleteURL(alias)
	if err != nil {
		log.Error("failed to delete url", sl.Err(err))
		os.Exit(1)
	}

	// TODO: router: chi, "chi render"

	// TODO: run server
}

func setupLogger(env config.Env) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(
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
