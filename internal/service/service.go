package service

import "github.com/VlanAni/nosql_over_postgres/internal/repository"

type Service struct {
	repo repository.DocumentStorage
}
