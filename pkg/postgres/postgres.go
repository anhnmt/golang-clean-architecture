package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(cfgApp config.App, cfgPostgres config.Postgres) (*Postgres, error) {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfgPostgres.User, cfgPostgres.Password),
		Host:   fmt.Sprintf("%s:%d", cfgPostgres.Host, cfgPostgres.Port),
		Path:   cfgPostgres.Database,
	}

	q := dsn.Query()
	q.Add("sslmode", cfgPostgres.SSLMode)
	q.Add("application_name", cfgApp.Name)

	// Migrate
	if cfgPostgres.Migrate {
		log.Info().Msg("Running migrations...")

		err := Migrate(dsn.String())
		if err != nil {
			return nil, err
		}

		log.Info().Msg("Migrations completed")
	}

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, fmt.Errorf("poolConfig - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = cfgPostgres.MaxConns
	poolConfig.MinConns = cfgPostgres.MinConns

	// Config maxConnIdleTime
	maxConnIdleTime, err := time.ParseDuration(cfgPostgres.MaxConnIdleTime)
	if err != nil {
		return nil, fmt.Errorf("maxConnIdleTime - time.ParseDuration: %w", err)
	}

	poolConfig.MaxConnIdleTime = maxConnIdleTime

	// Config maxConnLifetime
	maxConnLifetime, err := time.ParseDuration(cfgPostgres.MaxConnLifetime)
	if err != nil {
		return nil, fmt.Errorf("maxConnLifetime - time.ParseDuration: %w", err)
	}

	poolConfig.MaxConnLifetime = maxConnLifetime

	// Config connTimeout
	connTimeout, err := time.ParseDuration(cfgPostgres.ConnTimeout)
	if err != nil {
		return nil, fmt.Errorf("connTimeout - time.ParseDuration: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig err: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("postgres ping err: %w", err)
	}

	pg := &Postgres{
		Pool: pool,
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func Migrate(dbUrl string) error {
	dbUrl = strings.ReplaceAll(dbUrl, "postgres://", "pgx5://")
	m, err := migrate.New("file://db/migrations", dbUrl)
	if err != nil {
		return err
	}

	defer m.Close()

	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
