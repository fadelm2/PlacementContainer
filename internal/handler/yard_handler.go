package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"place-container/internal/dto"
	"place-container/internal/usecase"
)

type YardHandler struct {
	uc       *usecase.YardUsecase
	validate *validator.Validate
}

func NewYardHandler(u *usecase.YardUsecase) *YardHandler {
	return &YardHandler{uc: u, validate: validator.New()}
}

func (h *YardHandler) Suggestion(c *fiber.Ctx) error {
	var req dto.SuggestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := h.uc.SuggestPosition(context.Background(), req)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"suggested_position": res})
}

func (h *YardHandler) Placement(c *fiber.Ctx) error {
	var req dto.PlacementRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.uc.PlaceContainer(context.Background(), req); err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Success"})
}

func (h *YardHandler) Pickup(c *fiber.Ctx) error {
	var req dto.PickupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.uc.PickupContainer(context.Background(), req); err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Success"})
}
