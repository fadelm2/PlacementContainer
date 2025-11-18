package route

import (
	"github.com/gofiber/fiber/v2"
	"place-container/internal/delivery/http"
)

type RouteConfig struct {
	App             *fiber.App
	YardController  *http.YardController
	BlockController *http.BlockController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/yard/:PostId", c.YardController.Update)
	c.App.Post("/api/yard", c.YardController.Create)
	c.App.Put("/api/yard/:PostId", c.YardController.Update)

}
