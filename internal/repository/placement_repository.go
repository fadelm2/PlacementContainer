package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type PlacementRepository struct {
	Repository[entity.Yard]
	Log *logrus.Logger
}

func NewPlacementRepositoryRepositoryRepository(log *logrus.Logger) *PlacementRepository {
	return &PlacementRepository{
		Log: log,
	}
}
