package service

import (
	"context"
	"fmt"

	"github.com/VlanAni/nosql_over_postgres/internal/repository"
	"github.com/VlanAni/nosql_over_postgres/internal/validator"
)

func (s *Service) Delete(ctx context.Context, collectionName, id string) error {
	if err := validator.ValidateCollectionName(collectionName); err != nil {
		return fmt.Errorf("incorrect collection name: %w", err)
	}

	return s.repo.Delete(ctx, repository.NewDeleteRequest(
		collectionName,
		id,
	))
}
