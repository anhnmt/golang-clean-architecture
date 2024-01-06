package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

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

	log.Info().
		Any("app", cfg.App).
		Any("log", cfg.Log).
		Any("postgres", cfg.Postgres).
		Msg("Hello, World!")

	_, err = postgres.New(cfg.Postgres)
	if err != nil {
		log.Err(err).Msg("failed get postgres")
		return
	}
}
