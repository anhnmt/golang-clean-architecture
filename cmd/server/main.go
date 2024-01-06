package main

import (
	"fmt"
	"log/slog"

	"github.com/anhnmt/golang-clean-architecture/cmd/server/config"
	"github.com/anhnmt/golang-clean-architecture/pkg/logger"
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("failed get config: %w", err))
	}

	logger.New(cfg.Log)

	slog.Info("Hello, World!",
		slog.Any("app", cfg.App),
		slog.Any("log", cfg.Log),
		slog.Any("postgres", cfg.Postgres),
	)

	_, err = postgres.New(cfg.Postgres)
	if err != nil {
		slog.Error("failed get postgres", slog.Any("error", err))
		return
	}
}
