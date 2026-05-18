package service

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
)

func TestGet_Success(t *testing.T) {
	expectedData := []byte(`{"name":"John"}`)

	mock := &repository.MockStorage{
		GetFunc: func(ctx context.Context, req *repository.GetRequest) ([]byte, error) {
			return expectedData, nil
		},
	}
	svc := NewService(mock)

	result, err := svc.Get(context.Background(), "users", "123")

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if !bytes.Equal(result, expectedData) {
		t.Errorf("expected %s, got %s", expectedData, result)
	}
}

func TestGet_InvalidName(t *testing.T) {
	mock := &repository.MockStorage{}
	svc := NewService(mock)

	_, err := svc.Get(context.Background(), "invalid-name", "123")

	if err == nil {
		t.Error("expected validation error, got nil")
	}
}

func TestGet_NotFound(t *testing.T) {
	mock := &repository.MockStorage{
		GetFunc: func(ctx context.Context, req *repository.GetRequest) ([]byte, error) {
			return nil, errors.New("not found")
		},
	}
	svc := NewService(mock)

	_, err := svc.Get(context.Background(), "users", "999")

	if err == nil {
		t.Error("expected not found error, got nil")
	}
}
