package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"place-container/internal/entity"
	"place-container/internal/repository"
)

type pgRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) repository.Repository {
	return &pgRepo{db: db}
}

func (r *pgRepo) GetYardByCode(ctx context.Context, code string) (*entity.Yard, error) {
	var y entity.Yard
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&y).Error; err != nil {
		return nil, err
	}
	return &y, nil
}

func (r *pgRepo) GetBlockByCode(ctx context.Context, yardCode, blockCode string) (*entity.Block, error) {
	var b entity.Block
	if err := r.db.WithContext(ctx).
		Joins("JOIN yards on yards.id = blocks.yard_id").
		Where("yards.code = ? AND blocks.code = ?", yardCode, blockCode).
		First(&b).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *pgRepo) ListBlocksByYard(ctx context.Context, yardCode string) ([]entity.Block, error) {
	var blocks []entity.Block
	if err := r.db.WithContext(ctx).
		Joins("JOIN yards on yards.id = blocks.yard_id").
		Where("yards.code = ?", yardCode).Find(&blocks).Error; err != nil {
		return nil, err
	}
	return blocks, nil
}

func (r *pgRepo) GetPlacementByPosition(ctx context.Context, yardCode, blockCode string, slot, row, tier int) (*entity.Placement, error) {
	var p entity.Placement
	err := r.db.WithContext(ctx).
		Joins("JOIN blocks on placements.block_id = blocks.id").
		Joins("JOIN yards on placements.yard_id = yards.id").
		Where("yards.code = ? AND blocks.code = ? AND placements.slot = ? AND placements.row = ? AND placements.tier = ?",
			yardCode, blockCode, slot, row, tier).
		First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &p, err
}

func (r *pgRepo) GetPlacementByContainer(ctx context.Context, containerNumber string) (*entity.Placement, error) {
	var p entity.Placement
	if err := r.db.WithContext(ctx).Where("container_number = ?", containerNumber).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *pgRepo) ListPlacementsByBlock(ctx context.Context, yardCode, blockCode string) ([]entity.Placement, error) {
	var placements []entity.Placement
	if err := r.db.WithContext(ctx).
		Joins("JOIN yards on placements.yard_id = yards.id").
		Joins("JOIN blocks on placements.block_id = blocks.id").
		Where("yards.code = ? AND blocks.code = ?", yardCode, blockCode).
		Find(&placements).Error; err != nil {
		return nil, err
	}
	return placements, nil
}

func (r *pgRepo) DeletePlacementByContainer(ctx context.Context, containerNumber string) error {
	return r.db.WithContext(ctx).Where("container_number = ?", containerNumber).Delete(&entity.Placement{}).Error
}

func (r *pgRepo) ListYardPlansByBlock(ctx context.Context, yardCode, blockCode string) ([]entity.YardPlan, error) {
	var plans []entity.YardPlan
	if err := r.db.WithContext(ctx).
		Joins("JOIN yards on yard_plans.yard_id = yards.id").
		Joins("JOIN blocks on yard_plans.block_id = blocks.id").
		Where("yards.code = ? AND blocks.code = ?", yardCode, blockCode).
		Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// CheckOverlap checks if requested slot..slot+width-1 overlaps any existing placement ranges
// overlap condition: requestedSlot < existing.slot + existing.width AND existing.slot < requestedSlot+width
func (r *pgRepo) CheckOverlap(ctx context.Context, yardCode, blockCode string, slot, row, tier, width int) (bool, error) {
	var p entity.Placement
	err := r.db.WithContext(ctx).
		Joins("JOIN blocks on placements.block_id = blocks.id").
		Joins("JOIN yards on placements.yard_id = yards.id").
		Where("yards.code = ? AND blocks.code = ? AND placements.row = ? AND placements.tier = ?", yardCode, blockCode, row, tier).
		Where("? < placements.slot + placements.width AND placements.slot < ?", slot, slot+width).
		First(&p).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreatePlacementTx creates placement using a transaction and locks relevant rows to avoid race
func (r *pgRepo) CreatePlacementTx(ctx context.Context, p *entity.Placement) error {
	// Using transaction with serializable-ish locking
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Acquire a lock on blocks table row to reduce race (example). Adjust locking strategy per your workload.
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", p.BlockID).First(&entity.Block{}).Error; err != nil {
			return err
		}

		// Double-check overlap inside transaction
		var existing entity.Placement
		err := tx.Where("block_id = ? AND row = ? AND tier = ? AND ? < slot + width AND slot < ?",
			p.BlockID, p.Row, p.Tier, p.Slot, p.Slot+p.Width).First(&existing).Error
		if err == nil {
			return fmt.Errorf("position overlap inside transaction")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := tx.Create(p).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *pgRepo) HasContainerAbove(ctx context.Context, yard, block string, slot, row, tier, width int) (bool, error) {

	upperTier := tier + 1

	var count int64

	err := r.db.WithContext(ctx).
		Model(&entity.Placement{}).
		Where("yard_id = ? AND block_id = ?", yard, block).
		Where("row = ?", row).
		Where("tier = ?", upperTier).
		Where("slot <= ? AND slot + width - 1 >= ?", slot+width-1, slot).
		Count(&count).Error

	return count > 0, err
}
func (r *pgRepo) HasLeftContainer(ctx context.Context, yard, block string, row, tier, slot int) (bool, error) {
	if slot <= 1 {
		return false, nil
	}

	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Placement{}).
		Where("yard_id = ? AND block_id = ?", yard, block).
		Where("row = ? AND tier = ?", row, tier).
		Where("? BETWEEN slot AND slot + width - 1", slot-1).
		Count(&count).Error
	return count > 0, err
}

func (r *pgRepo) HasRightContainer(ctx context.Context, yard, block string, row, tier, slot, width int) (bool, error) {
	slotRight := slot + width

	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Placement{}).
		Where("yard_id = ? AND block_id = ?", yard, block).
		Where("row = ? AND tier = ?", row, tier).
		Where("? BETWEEN slot AND slot + width - 1", slotRight).
		Count(&count).Error
	return count > 0, err
}
