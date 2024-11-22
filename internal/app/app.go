package app

import (
	"fmt"
	"log/slog"

	"github.com/Bakhram74/effective-mobile-test-work.git/config"
	db "github.com/Bakhram74/effective-mobile-test-work.git/db/sqlc"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/http"
	"github.com/Bakhram74/effective-mobile-test-work.git/internal/service"
	"github.com/Bakhram74/effective-mobile-test-work.git/pkg/client/postgres"
	"github.com/Bakhram74/effective-mobile-test-work.git/pkg/httpServer"
)

func Run(config *config.Config) {
	slog.Info("Postgresql initializing")
	pool, err := postgres.New(config.DBSource)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	err = RunMigration(config)
	if err != nil {
		panic(fmt.Sprintf("Migration error: %s", err.Error()))
	}

	queries := db.New(pool)

	service := service.NewService(queries)

	handler := http.NewHandler(config, service).Init()

	slog.Info("Runnig app server at", slog.String("addr", config.HTTPServerAddress))
	srv := httpServer.NewServer(config, handler)
	if err := srv.Run(); err != nil {
		slog.Error("Error occurred while running http server", slog.String("error", err.Error()))
	}

}
