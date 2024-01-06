package main

import (
	"fmt"
	"log/slog"

	"github.com/anhnmt/golang-clean-architecture/cmd/server/config"
	"github.com/anhnmt/golang-clean-architecture/pkg/logger"
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
	)
}
