package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type ContainerRepository struct {
	Repository[entity.Yard]
	Log *logrus.Logger
}

func NewContainerRepositoryRepositoryRepository(log *logrus.Logger) *ContainerRepository {
	return &ContainerRepository{
		Log: log,
	}
}
