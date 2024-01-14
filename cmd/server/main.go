package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
		Any("server", cfg.Server).
		Msg("Hello, World!")

	ctx, cancel := context.WithCancel(context.Background())

	srv, cleanup, err := InitServer(ctx, cfg)
	if err != nil {
		log.Panic().Err(err).Msg("failed to init server")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func(_srv server.Server) {
		err = _srv.Start(ctx)
		if err != nil {
			log.Err(err).Msg("failed to start server")
			cancel()
			<-ctx.Done()
		}
	}(srv)

	select {
	case v := <-quit:
		cleanup()
		log.Info().Any("v", v).Msg("signal.Notify")
	case done := <-ctx.Done():
		cleanup()
		log.Info().Any("done", done).Msg("ctx.Done")
	}

	log.Info().Msg("Gracefully shutting down")
}
