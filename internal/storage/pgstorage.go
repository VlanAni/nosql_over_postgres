package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	pgsqlgenerator PGSQLGenerator
	pgpool         *pgxpool.Pool
}

func NewPostgresStorage(ctx context.Context) (*PostgresStorage, error) {
	pgpool, err := pgxpool.New(ctx, os.Getenv("PGDATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot create a connection pool: %w", err)
	}

	if err := pgpool.Ping(ctx); err != nil {
		pgpool.Close()
		return nil, fmt.Errorf("cannot ping the database from the pool: %w", err)
	}

	return &PostgresStorage{
		pgsqlgenerator: PGSQLGenerator{},
		pgpool:         pgpool,
	}, nil
}

func (ps *PostgresStorage) Put(ctx context.Context, request *repository.PutRequest) error {
	createRequest := ps.pgsqlgenerator.generateCreateRequest(request.CollectionName())

	if _, err := ps.pgpool.Exec(ctx, createRequest); err != nil {
		return fmt.Errorf("cannot execute the create query: %w", err)
	}

	insertRequest := ps.pgsqlgenerator.generateInsertRequest(request.CollectionName())

	if _, err := ps.pgpool.Exec(ctx, insertRequest, request.ID(), request.Payload()); err != nil {
		return fmt.Errorf("cannot execute the request query: %w", err)
	}

	return nil
}

func (ps *PostgresStorage) Get(ctx context.Context, request *repository.GetRequest) ([]byte, error) {
	getRequest := ps.pgsqlgenerator.generateGetRequest(request.CollectionName())

	rows := ps.pgpool.QueryRow(ctx, getRequest, request.ID())

	var payload []byte

	err := rows.Scan(&payload)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("there is no any record with such id: %s", request.ID())
		}
		return nil, fmt.Errorf("failed to scan payload: %w", err)
	}

	return payload, nil
}

func (ps *PostgresStorage) Delete(ctx context.Context, request *repository.DeleteRequest) error {
	deleteRequest := ps.pgsqlgenerator.generateDeleteRequest(request.CollectionName())

	if _, err := ps.pgpool.Exec(ctx, deleteRequest, request.ID()); err != nil {
		return fmt.Errorf("cannot execute the delete request: %w", err)
	}

	return nil
}

func (ps *PostgresStorage) Close() {
	ps.pgpool.Close()
}
