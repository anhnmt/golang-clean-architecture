package service

import (
	"github.com/google/wire"

	userv1 "github.com/anhnmt/golang-clean-architecture/api/protobuf/user/v1"
	"github.com/anhnmt/golang-clean-architecture/internal/user"
)

var ServiceProviderSet = wire.NewSet(
	New,
	user.UserProviderSet,
)

type Service interface {
}

type service struct {
}

func New(
	_ userv1.UserServiceServer,
) Service {
	return &service{}
}
