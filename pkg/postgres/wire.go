//go:build wireinject
// +build wireinject

package postgres

import (
	"github.com/google/wire"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

func NewDBEngine(cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	panic(wire.Build(
		DBEngineFunc,
	))
}

func DBEngineFunc(cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	db, err := New(cfgApp, cfgPostgres)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
