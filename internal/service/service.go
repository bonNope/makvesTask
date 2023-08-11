package service

import (
	"github.com/bonNope/makvesTask/internal/contracts/service"
	"github.com/bonNope/makvesTask/internal/repository"
)

type Service struct {
	service.EntityServiceContract
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		EntityServiceContract: NewEntityService(repo.EntityRepositoryContract),
	}
}
