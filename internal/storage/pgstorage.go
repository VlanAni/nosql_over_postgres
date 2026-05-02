package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	pgsqlgenerator PGSQLGenerator
	pgpool         *pgxpool.Pool
}
