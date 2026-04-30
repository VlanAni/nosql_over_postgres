package repository

import (
	"context"
)

type DocumentStorage interface {
	Put(ctx context.Context, request *PutRequest) error
	Delete(ctx context.Context, request *DeleteRequest) error
	Get(ctx context.Context, request *GetRequest) ([]byte, error)
}
