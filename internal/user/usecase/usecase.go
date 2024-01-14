package usecase

import (
	"github.com/google/wire"

	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

var UserUseCaseSet = wire.NewSet(New)

type UseCase interface {
}

type useCase struct {
	pg postgres.DBEngine
}

func New(
	pg postgres.DBEngine,
) UseCase {
	return &useCase{
		pg: pg,
	}
}
