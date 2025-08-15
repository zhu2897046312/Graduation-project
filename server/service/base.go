package service

import (
	"server/repository"
)

type Service struct {
	repoFactory *repository.RepositoryFactory
}

func NewService(repoFactory *repository.RepositoryFactory) *Service {
	return &Service{
		repoFactory: repoFactory,
	}
} 