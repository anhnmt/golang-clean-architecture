package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBEngine interface {
	Pool() *pgxpool.Pool
	Close()
}
