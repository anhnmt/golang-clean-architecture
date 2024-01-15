package rpc

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	gen "github.com/anhnmt/golang-clean-architecture/api/protobuf/gen"
	"github.com/anhnmt/golang-clean-architecture/internal/user/usecase"
)

var _ gen.UserServiceServer = (*handler)(nil)

type handler struct {
	gen.UnimplementedUserServiceServer

	uc usecase.UseCase
}

func New(
	grpcServer *grpc.Server,
	uc usecase.UseCase,
) gen.UserServiceServer {
	svc := &handler{
		uc: uc,
	}

	gen.RegisterUserServiceServer(grpcServer, svc)
	return svc
}

func (h *handler) List(_ context.Context, request *gen.ListRequest) (*gen.ListResponse, error) {
	log.Info().Any("request", request).Msg("List")

	return nil, nil
}
