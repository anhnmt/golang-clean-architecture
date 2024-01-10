package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-clean-architecture/cmd/server/config"
	"github.com/anhnmt/golang-clean-architecture/internal/server"
	"github.com/anhnmt/golang-clean-architecture/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("failed get config: %w", err))
	}

	logger.New(cfg.Log)

	log.Info().
		Any("app", cfg.App).
		Any("log", cfg.Log).
		Any("postgres", cfg.Postgres).
		Msg("Hello, World!")

	// _, err = postgres.New(cfg.App, cfg.Postgres)
	// if err != nil {
	//     log.Panic().Err(err).Msg("failed to connect to postgres")
	//     return
	// }

	srv := server.New(cfg.Server)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = srv.Start(ctx)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}
