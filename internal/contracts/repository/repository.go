package repository

import "github.com/bonNope/makvesTask/internal/models"

type EntityRepositoryContract interface {
	GetById(id int) *models.Entity
}
