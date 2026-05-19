package service

import (
	"context"
	"errors"
	"testing"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
)

func TestPut_Success(t *testing.T) {
	mock := &repository.MockStorage{
		PutFunc: func(ctx context.Context, req *repository.PutRequest) error {
			return nil
		},
	}
	svc := NewService(mock)

	err := svc.Put(context.Background(), "users", "123", []byte(`{"name":"John"}`))

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestPut_InvalidName(t *testing.T) {
	mock := &repository.MockStorage{}
	svc := NewService(mock)

	err := svc.Put(context.Background(), "bad name!", "123", []byte(`{}`))

	if err == nil {
		t.Error("expected validation error, got nil")
	}
}

func TestPut_DatabaseError(t *testing.T) {
	mock := &repository.MockStorage{
		PutFunc: func(ctx context.Context, req *repository.PutRequest) error {
			return errors.New("db connection lost")
		},
	}
	svc := NewService(mock)

	err := svc.Put(context.Background(), "users", "123", []byte(`{}`))

	if err == nil {
		t.Error("expected database error, got nil")
	}
}
