package converter

import (
	"place-container/internal/entity"
	"place-container/internal/model"
)

func BlockToResponse(Block *entity.Block) *model.BlockResponse {

	return &model.BlockResponse{
		ID:        Block.ID,
		YardID:    Block.YardID,
		Name:      Block.Name,
		TotalSlot: Block.TotalSlot,
		TotalRow:  Block.TotalRow,
		TotalTier: Block.TotalTier,
	}
}
