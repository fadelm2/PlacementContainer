package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"place-container/internal/model"
	"place-container/internal/usecase"
)

type BlockController struct {
	UseCase *usecase.BlockUseCase
	Log     *logrus.Logger
}

func NewBlockController(useCase *usecase.BlockUseCase, log *logrus.Logger) *BlockController {
	return &BlockController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *BlockController) Create(ctx *fiber.Ctx) error {

	request := new(model.CreateBlockRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Block")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BlockResponse]{Data: response})
}

func (c *BlockController) Get(ctx *fiber.Ctx) error {

	request := &model.GetBlockRequest{
		ID: ctx.Params("BlockId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Block")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BlockResponse]{Data: response})
}

func (c *BlockController) Update(ctx *fiber.Ctx) error {

	request := new(model.UpdateBlockRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("BlockId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating Block")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BlockResponse]{Data: response})
}
