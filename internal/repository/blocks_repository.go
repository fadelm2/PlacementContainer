package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type BlocksRepository struct {
	Repository[entity.Yard]
	Log *logrus.Logger
}

func NewBlocksRepositoryRepositoryRepository(log *logrus.Logger) *BlocksRepository {
	return &BlocksRepository{
		Log: log,
	}
}
