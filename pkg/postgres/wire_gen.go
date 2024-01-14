// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package postgres

import (
	"context"
	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

import (
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Injectors from wire.go:

func NewDBEngine(ctx context.Context, cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	dbEngine, cleanup, err := DBEngineFunc(ctx, cfgApp, cfgPostgres)
	if err != nil {
		return nil, nil, err
	}
	return dbEngine, func() {
		cleanup()
	}, nil
}

// wire.go:

func DBEngineFunc(ctx context.Context, cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, func(), error) {
	db, err := New(ctx, cfgApp, cfgPostgres)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
