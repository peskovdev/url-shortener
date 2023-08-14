package main

import (
	"github.com/peskovdev/url-shortener/internal/config"
)

func main() {
	config.MustLoad()

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: router: chi, "chi render"

	// TODO: run server
}
