package postgres

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(cfg config.Postgres) (*Postgres, error) {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:   cfg.Database,
	}

	q := dsn.Query()
	q.Add("sslmode", cfg.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, fmt.Errorf("poolConfig - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns

	// Config maxConnIdleTime
	maxConnIdleTime, err := time.ParseDuration(cfg.MaxConnIdleTime)
	if err != nil {
		return nil, fmt.Errorf("maxConnIdleTime - time.ParseDuration: %w", err)
	}

	poolConfig.MaxConnIdleTime = maxConnIdleTime

	// Config maxConnLifetime
	maxConnLifetime, err := time.ParseDuration(cfg.MaxConnLifetime)
	if err != nil {
		return nil, fmt.Errorf("maxConnLifetime - time.ParseDuration: %w", err)
	}

	poolConfig.MaxConnLifetime = maxConnLifetime

	// Config connTimeout
	connTimeout, err := time.ParseDuration(cfg.ConnTimeout)
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
