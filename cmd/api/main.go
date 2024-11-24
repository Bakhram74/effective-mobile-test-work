package main

import (
	"log/slog"
	"os"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/app"
)

// @title      Music info
// @version 0.0.1
// @description     API docs for Effective-mobile-test-work.
// @host      localhost:8080
// @BasePath  /song
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	app.Run(&cfg)
}
