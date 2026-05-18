package repository

import "context"

type MockStorage struct {
	PutFunc    func(ctx context.Context, req *PutRequest) error
	GetFunc    func(ctx context.Context, req *GetRequest) ([]byte, error)
	DeleteFunc func(ctx context.Context, req *DeleteRequest) error
}

func (m *MockStorage) Put(ctx context.Context, req *PutRequest) error {
	return m.PutFunc(ctx, req)
}

func (m *MockStorage) Get(ctx context.Context, req *GetRequest) ([]byte, error) {
	return m.GetFunc(ctx, req)
}

func (m *MockStorage) Delete(ctx context.Context, req *DeleteRequest) error {
	return m.DeleteFunc(ctx, req)
}

func (m *MockStorage) Close() {
}
