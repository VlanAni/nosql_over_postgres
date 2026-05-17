package service

import "github.com/VlanAni/nosql_over_postgres/internal/repository"

func NewService(repo repository.DocumentStorage) *Service {
	return &Service{
		repo: repo,
	}
}
