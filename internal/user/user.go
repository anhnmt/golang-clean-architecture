package user

import (
	"github.com/google/wire"

	"github.com/anhnmt/golang-clean-architecture/internal/user/delivery/rpc"
	"github.com/anhnmt/golang-clean-architecture/internal/user/usecase"
)

var UserProviderSet = wire.NewSet(
	rpc.New,
	usecase.New,
)
