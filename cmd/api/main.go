package main

import (
	"log/slog"
	"os"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/app"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	app.Run(&cfg)
}
