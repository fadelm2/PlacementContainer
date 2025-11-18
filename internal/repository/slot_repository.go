package repository

import (
	"github.com/sirupsen/logrus"
	"place-container/internal/entity"
)

type SlotRepository struct {
	Repository[entity.Yard]
	Log *logrus.Logger
}

func NewSlotRepositoryRepository(log *logrus.Logger) *SlotRepository {
	return &SlotRepository{
		Log: log,
	}
}
