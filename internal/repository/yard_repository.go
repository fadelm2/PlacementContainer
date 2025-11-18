package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type YardsRepository struct {
	Repository[entity.Yard]
	Log *logrus.Logger
}

func NewYardsRepository(log *logrus.Logger) *YardsRepository {
	return &YardsRepository{
		Log: log,
	}
}
