package usecase

import (
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

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
