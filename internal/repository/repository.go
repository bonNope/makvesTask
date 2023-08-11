package repository

import (
	"github.com/bonNope/makvesTask/internal/contracts/repository"
	"github.com/bonNope/makvesTask/internal/models"
)

type Repository struct {
	repository.EntityRepositoryContract
}

func NewRepository(entities map[int]*models.Entity) *Repository {
	return &Repository{
		EntityRepositoryContract: NewEntityRepository(entities),
	}
}
