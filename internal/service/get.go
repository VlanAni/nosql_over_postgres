package service

import (
	"context"
	"fmt"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
	"github.com/VlanAni/nosql_over_postgres/internal/validator"
)

func (s *Service) Get(ctx context.Context, collectionName, id string) ([]byte, error) {
	if err := validator.ValidateCollectionName(collectionName); err != nil {
		return nil, fmt.Errorf("incorrect collection name: %w", err)
	}

	return s.repo.Get(ctx, repository.NewGetRequest(
		collectionName,
		id,
	))
}
