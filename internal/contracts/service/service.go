package service

import "github.com/bonNope/makvesTask/internal/models"

type EntityServiceContract interface {
	GetEntitiesByIds(ids []int) []*models.Entity
}
