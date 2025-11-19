package router

import (
	"github.com/gofiber/fiber/v2"
	"place-container/internal/handler"
)

func Setup(app *fiber.App, h *handler.YardHandler) {
	api := app.Group("/api")
	api.Post("/suggestion", h.Suggestion)
	api.Post("/placement", h.Placement)
	api.Post("/pickup", h.Pickup)
}
