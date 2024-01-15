package service

import (
	"github.com/google/wire"

	gen "github.com/anhnmt/golang-clean-architecture/api/protobuf/gen"
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
	_ gen.UserServiceServer,
) Service {
	return &service{}
}
