package service

import (
	"context"
	"fmt"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
	"github.com/VlanAni/nosql_over_postgres/internal/validator"
)

func (s *Service) Put(ctx context.Context, collectionName, id string, payload []byte) error {
	if len(payload) == 0 {
		return fmt.Errorf("incorrect payload")
	}

	if err := validator.ValidateCollectionName(collectionName); err != nil {
		return fmt.Errorf("incorrect collection name: %w", err)
	}

	return s.repo.Put(ctx, repository.NewPutRequest(
		collectionName,
		id,
		payload,
	))
}
