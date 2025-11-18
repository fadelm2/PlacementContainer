package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"place-container/internal/entity"
	"place-container/internal/model"
	"place-container/internal/model/converter"
	"place-container/internal/repository"
)

type BlockUseCase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	Validate        *validator.Validate
	BlockRepository *repository.BlocksRepository
}

func NewBlockUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	BlockRepository *repository.BlocksRepository) *BlockUseCase {
	return &BlockUseCase{
		DB:              db,
		Log:             logger,
		Validate:        validate,
		BlockRepository: BlockRepository,
	}
}

func (c *BlockUseCase) Create(ctx context.Context,
	request *model.CreateBlockRequest) (*model.BlockResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	block := &entity.Block{
		ID:        request.ID, // harusnya open langsung default
		YardID:    request.YardID,
		Name:      request.Name,
		TotalRow:  request.TotalRow,
		TotalTier: request.TotalTier,
		TotalSlot: request.TotalSlot,
	}

	if err := c.BlockRepository.Create(tx, block); err != nil {
		c.Log.WithError(err).Error("error creating Block")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.BlockToResponse(block), nil

}

func (c *BlockUseCase) Update(ctx context.Context,
	request *model.UpdateBlockRequest) (*model.BlockResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	Block := new(entity.Block)
	if err := c.BlockRepository.FindById(tx, Block, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Input yang dimasukan ada kesalahan")
	}

	if request.YardID != "" {
		Block.YardID = request.YardID
	}
	if request.Name != "" {
		Block.Name = request.Name
	}
	if request.TotalSlot != nil {
		Block.TotalSlot = *request.TotalSlot
	}
	if request.TotalRow != nil {
		Block.TotalRow = *request.TotalRow
	}
	if request.TotalTier != nil {
		Block.TotalTier = *request.TotalTier
	}

	if err := c.BlockRepository.Update(tx, Block); err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.BlockToResponse(Block), nil

}

func (c *BlockUseCase) Get(ctx context.Context, request *model.GetBlockRequest) (*model.BlockResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Block := new(entity.Block)
	if err := c.BlockRepository.FindById(tx, Block, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.BlockToResponse(Block), nil
}
