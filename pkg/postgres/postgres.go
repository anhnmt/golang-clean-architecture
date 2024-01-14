package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

var _ DBEngine = (*postgres)(nil)

type postgres struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfgApp config.App, cfgPostgres config.Postgres) (DBEngine, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfgPostgres.User, cfgPostgres.Password),
		Host:   fmt.Sprintf("%s:%d", cfgPostgres.Host, cfgPostgres.Port),
		Path:   cfgPostgres.Database,
	}

	q := dsn.Query()
	q.Add("sslmode", cfgPostgres.SSLMode)
	q.Add("application_name", cfgApp.Name)

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

	newCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(newCtx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig err: %w", err)
	}

	if err = pool.Ping(newCtx); err != nil {
		return nil, fmt.Errorf("postgres ping err: %w", err)
	}

	// Migrate
	if cfgPostgres.Migrate {
		log.Info().Msg("Running migrations...")

		if err = Migrate(dsn); err != nil {
			return nil, fmt.Errorf("postgres migrate err: %w", err)
		}

		log.Info().Msg("Migrations completed")
	}

	pg := &postgres{
		pool: pool,
	}

	return pg, nil
}

func (p *postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}

func (p *postgres) Pool() *pgxpool.Pool {
	return p.pool
}

func Migrate(db url.URL) error {
	db.Scheme = "pgx5"
	m, err := migrate.New("file://db/migrations", db.String())
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
