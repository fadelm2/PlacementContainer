package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"place-container/internal/dto"
	"place-container/internal/entity"
	"place-container/internal/repository"
)

type YardUsecase struct {
	repo  repository.Repository
	redis *redis.Client
}

func NewYardUsecase(r repository.Repository, rd *redis.Client) *YardUsecase {
	return &YardUsecase{repo: r, redis: rd}
}

func (u *YardUsecase) SuggestPosition(ctx context.Context, req dto.SuggestionRequest) (*dto.SuggestionResponse, error) {
	cacheKey := "suggestion:" + req.Yard + ":" + req.ContainerNumber
	if u.redis != nil {
		if v, err := u.redis.Get(ctx, cacheKey).Result(); err == nil && v != "" {
			parts := strings.Split(v, ":")
			if len(parts) >= 4 {
				slot, _ := strconv.Atoi(parts[1])
				row, _ := strconv.Atoi(parts[2])
				tier, _ := strconv.Atoi(parts[3])
				width := 1
				if len(parts) >= 5 {
					width, _ = strconv.Atoi(parts[4])
				}
				return &dto.SuggestionResponse{Block: parts[0], Slot: slot, Row: row, Tier: tier, Width: width}, nil
			}
		}
	}

	blocks, err := u.repo.ListBlocksByYard(ctx, req.Yard)
	if err != nil {
		return nil, err
	}
	if len(blocks) == 0 {
		return nil, errors.New("no blocks found for yard")
	}

	for _, b := range blocks {
		plans, err := u.repo.ListYardPlansByBlock(ctx, req.Yard, b.Code)
		if err != nil {
			return nil, err
		}

		candidates := make([][3]int, 0)

		if len(plans) == 0 {
			for s := 1; s <= b.TotalSlot; s++ {
				for r := 1; r <= b.TotalRow; r++ {
					for t := 1; t <= b.TotalTier; t++ {
						candidates = append(candidates, [3]int{s, r, t})
					}
				}
			}
		} else {
			for _, p := range plans {
				if p.ContainerSize != req.ContainerSize || p.ContainerType != req.ContainerType {
					continue
				}
				for s := p.FromSlot; s <= p.ToSlot; s++ {
					for r := p.FromRow; r <= p.ToRow; r++ {
						for t := 1; t <= b.TotalTier; t++ {
							candidates = append(candidates, [3]int{s, r, t})
						}
					}
				}
			}
		}

		width := 1
		if req.ContainerSize == 40 {
			width = 2
		}

		for _, cell := range candidates {
			// ensure within bounds
			if cell[0]+width-1 > b.TotalSlot {
				continue
			}
			overlap, err := u.repo.CheckOverlap(ctx, req.Yard, b.Code, cell[0], cell[1], cell[2], width)
			if err != nil {
				return nil, err
			}
			if !overlap {
				resp := &dto.SuggestionResponse{
					Block: b.Code,
					Slot:  cell[0],
					Row:   cell[1],
					Tier:  cell[2],
					Width: width,
				}
				if u.redis != nil {
					_ = u.redis.Set(ctx, cacheKey, fmt.Sprintf("%s:%d:%d:%d:%d", b.Code, cell[0], cell[1], cell[2], width), 30*60).Err()
				}
				return resp, nil
			}
		}
	}
	return nil, errors.New("no available position")
}

func (u *YardUsecase) PlaceContainer(ctx context.Context, req dto.PlacementRequest) error {
	// check if container already placed
	if existing, _ := u.repo.GetPlacementByContainer(ctx, req.ContainerNumber); existing != nil {
		return errors.New("container already placed")
	}

	width := 1
	if req.Size == 40 {
		width = 2
	}

	// check occupancy
	overlap, err := u.repo.CheckOverlap(ctx, req.Yard, req.Block, req.Slot, req.Row, req.Tier, width)
	if err != nil {
		return err
	}
	if overlap {
		return errors.New("position occupied")
	}

	place := &entity.Placement{
		ID:              uuid.New().String(),
		YardID:          req.Yard,
		BlockID:         req.Block,
		ContainerNumber: req.ContainerNumber,
		Slot:            req.Slot,
		Row:             req.Row,
		Tier:            req.Tier,
		Width:           width,
		Size:            req.Size,
		Type:            req.Type,
		Height:          req.Height,
	}

	if err := u.repo.CreatePlacementTx(ctx, place); err != nil {
		return err
	}

	// purge suggestion cache if exists
	if u.redis != nil {
		_ = u.redis.Del(ctx, "suggestion:"+req.Yard+":"+req.ContainerNumber).Err()
	}
	return nil
}

func (u *YardUsecase) PickupContainer(ctx context.Context, req dto.PickupRequest) error {
	p, err := u.repo.GetPlacementByContainer(ctx, req.ContainerNumber)
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("container not found")
	}
	if err := u.repo.DeletePlacementByContainer(ctx, req.ContainerNumber); err != nil {
		return err
	}
	if u.redis != nil {
		_ = u.redis.Del(ctx, "suggestion:"+p.YardID+":"+req.ContainerNumber).Err()
	}
	return nil
}
