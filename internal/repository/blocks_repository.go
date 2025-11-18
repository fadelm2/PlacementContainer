package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type BlocksRepository struct {
	Repository[entity.Block]
	Log *logrus.Logger
}

func NewBlocksRepository(log *logrus.Logger) *BlocksRepository {
	return &BlocksRepository{
		Log: log,
	}
}
