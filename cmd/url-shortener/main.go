package main

import (
	"github.com/peskovdev/url-shortener/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", string(cfg.Env)))
	log.Debug("debug message are enabled")

	// TODO: init storage: sqlite

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
