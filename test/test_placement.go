package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"place-container/internal/entity"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetPlacementByContainer(ctx context.Context, container string) (*entity.Placement, error) {
	args := m.Called(ctx, container)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Placement), args.Error(1)
}

func (m *MockRepo) CheckOverlap(ctx context.Context, yard, block string, slot, row, tier, width int) (bool, error) {
	args := m.Called(ctx, yard, block, slot, row, tier, width)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepo) CreatePlacementTx(ctx context.Context, p *entity.Placement) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.Called(ctx, keys)
	cmd := redis.NewIntCmd(ctx)
	cmd.SetVal(1)
	return cmd
}
