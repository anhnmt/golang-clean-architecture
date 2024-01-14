package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-clean-architecture/cmd/server/config"
	"github.com/anhnmt/golang-clean-architecture/internal/grpc_server"
	"github.com/anhnmt/golang-clean-architecture/internal/server"
	"github.com/anhnmt/golang-clean-architecture/pkg/logger"
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
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
		Any("server", cfg.Server).
		Msg("Hello, World!")

	pg, cleanup, err := postgres.NewDBEngine(cfg.App, cfg.Postgres)
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to postgres")
	}
	defer cleanup()

	grpcServer := grpc_server.New(cfg.Server.Grpc)

	srv := server.NewServerEngine(cfg.Server, pg, grpcServer)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = srv.Start(ctx)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}
