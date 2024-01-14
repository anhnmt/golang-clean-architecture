//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/anhnmt/golang-clean-architecture/internal/user/delivery/rpc"
	"github.com/anhnmt/golang-clean-architecture/internal/user/usecase"
	"github.com/anhnmt/golang-clean-architecture/pkg/config"
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

func NewServerEngine(
	cfg config.Server,
	pg postgres.DBEngine,
	grpcServer *grpc.Server,
) Server {
	panic(wire.Build(
		New,
		usecase.UserUseCaseSet,
		rpc.UserHandlerSet,
	))
}
