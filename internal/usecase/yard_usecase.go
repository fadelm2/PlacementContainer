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

type YardUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	YardRepository *repository.YardsRepository
}

func NewYardUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	YardRepository *repository.YardsRepository) *YardUseCase {
	return &YardUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		YardRepository: YardRepository,
	}
}

func (c *YardUseCase) Create(ctx context.Context,
	request *model.CreateYardRequest) (*model.YardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Yard := &entity.Yard{
		ID:          request.ID, // harusnya open langsung default
		Name:        request.Name,
		Description: request.Description,
	}

	if err := c.YardRepository.Create(tx, Yard); err != nil {
		c.Log.WithError(err).Error("error creating Yard")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.YardToResponse(Yard), nil

}

func (c *YardUseCase) Update(ctx context.Context,
	request *model.UpdateYardRequest) (*model.YardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	yard := new(entity.Yard)
	if err := c.YardRepository.FindById(tx, yard, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Input yang dimasukan ada kesalahan")
	}

	if request.Name != "" {
		yard.Name = request.Name
	}

	if request.Description != "" {
		yard.Name = request.Description
	}

	if err := c.YardRepository.Update(tx, yard); err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.YardToResponse(yard), nil

}

func (c *YardUseCase) Get(ctx context.Context, request *model.GetYardRequest) (*model.YardResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	yard := new(entity.Yard)
	if err := c.YardRepository.FindById(tx, yard, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.YardToResponse(yard), nil
}
