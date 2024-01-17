package rpc

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	userv1 "github.com/anhnmt/golang-clean-architecture/api/protobuf/user/v1"
	"github.com/anhnmt/golang-clean-architecture/internal/user/usecase"
)

var _ userv1.UserServiceServer = (*handler)(nil)

type handler struct {
	userv1.UnimplementedUserServiceServer

	uc usecase.UseCase
}

func New(
	grpcServer *grpc.Server,
	uc usecase.UseCase,
) userv1.UserServiceServer {
	svc := &handler{
		uc: uc,
	}

	userv1.RegisterUserServiceServer(grpcServer, svc)
	return svc
}

func (h *handler) List(_ context.Context, request *userv1.ListRequest) (*userv1.ListResponse, error) {
	log.Info().Any("request", request).Msg("List")

	return nil, nil
}
