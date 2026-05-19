package service

import (
	"context"
	"errors"
	"testing"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
)

func TestDelete_Success(t *testing.T) {
	mock := &repository.MockStorage{
		DeleteFunc: func(ctx context.Context, req *repository.DeleteRequest) error {
			return nil
		},
	}
	svc := NewService(mock)

	err := svc.Delete(context.Background(), "users", "123")

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestDelete_InvalidName(t *testing.T) {
	mock := &repository.MockStorage{}
	svc := NewService(mock)

	err := svc.Delete(context.Background(), "DROP TABLE", "123")

	if err == nil {
		t.Error("expected validation error, got nil")
	}
}

func TestDelete_DatabaseError(t *testing.T) {
	mock := &repository.MockStorage{
		DeleteFunc: func(ctx context.Context, req *repository.DeleteRequest) error {
			return errors.New("timeout")
		},
	}
	svc := NewService(mock)

	err := svc.Delete(context.Background(), "users", "123")

	if err == nil {
		t.Error("expected database error, got nil")
	}
}
