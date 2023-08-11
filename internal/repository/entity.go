package repository

import "github.com/bonNope/makvesTask/internal/models"

type EntityRepository struct {
	entities map[int]*models.Entity
}

func NewEntityRepository(entities map[int]*models.Entity) *EntityRepository {
	return &EntityRepository{entities: entities}
}

func (r *EntityRepository) GetById(id int) *models.Entity {
	return r.entities[id]
}
