//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"github.com/anhnmt/golang-clean-architecture/cmd/server/config"
	"github.com/anhnmt/golang-clean-architecture/internal/grpc_server"
	"github.com/anhnmt/golang-clean-architecture/internal/server"
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

func InitServer(ctx context.Context, cfg config.Config) (server.Server, func(), error) {
	panic(wire.Build(
		initServerFunc,
	))
}

func initServerFunc(ctx context.Context, cfg config.Config) (server.Server, func(), error) {
	pg, cleanup, err := postgres.NewDBEngine(ctx, cfg.App, cfg.Postgres)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc_server.New(cfg.Server.Grpc)
	srv := server.NewServerEngine(cfg.Server, pg, grpcServer)

	return srv, func() {
		cleanup()
		srv.Close()
	}, nil
}
