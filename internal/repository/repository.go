package repository

import (
	"context"

	"place-container/internal/entity"
)

type Repository interface {
	// yard / block
	GetYardByCode(ctx context.Context, code string) (*entity.Yard, error)
	GetBlockByCode(ctx context.Context, yardCode, blockCode string) (*entity.Block, error)
	ListBlocksByYard(ctx context.Context, yardCode string) ([]entity.Block, error)

	// placement actions
	GetPlacementByPosition(ctx context.Context, yardCode, blockCode string, slot, row, tier int) (*entity.Placement, error)
	GetPlacementByContainer(ctx context.Context, containerNumber string) (*entity.Placement, error)
	ListPlacementsByBlock(ctx context.Context, yardCode, blockCode string) ([]entity.Placement, error)
	DeletePlacementByContainer(ctx context.Context, containerNumber string) error

	// yard plan
	ListYardPlansByBlock(ctx context.Context, yardCode, blockCode string) ([]entity.YardPlan, error)

	// overlap check (uses width of placement)
	CheckOverlap(ctx context.Context, yardCode, blockCode string, slot, row, tier, width int) (bool, error)

	// create placement using DB transaction
	CreatePlacementTx(ctx context.Context, p *entity.Placement) error
}
