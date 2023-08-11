package service

import (
	"github.com/bonNope/makvesTask/internal/contracts/repository"
	"github.com/bonNope/makvesTask/internal/models"
)

type EntityService struct {
	db repository.EntityRepositoryContract
}

func NewEntityService(db repository.EntityRepositoryContract) *EntityService {
	return &EntityService{db: db}
}

func (s *EntityService) GetEntitiesByIds(ids []int) []*models.Entity {
	entities := make([]*models.Entity, len(ids))
	for i, id := range ids {
		entities[i] = s.db.GetById(id)
	}

	return entities
}
