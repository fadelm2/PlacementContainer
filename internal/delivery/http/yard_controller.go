package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"place-container/internal/model"
	"place-container/internal/usecase"
)

type YardController struct {
	UseCase *usecase.YardUseCase
	Log     *logrus.Logger
}

func NewYardController(useCase *usecase.YardUseCase, log *logrus.Logger) *YardController {
	return &YardController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *YardController) Create(ctx *fiber.Ctx) error {

	request := new(model.CreateYardRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Yard")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.YardResponse]{Data: response})
}

func (c *YardController) Get(ctx *fiber.Ctx) error {

	request := &model.GetYardRequest{
		ID: ctx.Params("yardId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Yard")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.YardResponse]{Data: response})
}

func (c *YardController) Update(ctx *fiber.Ctx) error {

	request := new(model.UpdateYardRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("YardId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating Yard")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.YardResponse]{Data: response})
}
