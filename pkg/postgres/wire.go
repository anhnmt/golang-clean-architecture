//go:build wireinject
// +build wireinject

package postgres

import (
	"context"

	"github.com/google/wire"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

func NewDBEngine(ctx context.Context, cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	panic(wire.Build(
		DBEngineFunc,
	))
}

func DBEngineFunc(ctx context.Context, cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	db, err := New(ctx, cfgApp, cfgPostgres)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
