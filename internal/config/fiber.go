package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.GetBool("web.prefork"),
		ColorScheme:  fiber.DefaultColors,
	})
	corsSettings := cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173, http://localhost:8080, https://test.fadelweb.site, https://testapi.fadelweb.site, http://localhost:3000",
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept,  Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		//   ExposeHeaders:    "Origin",
	})
	app.Use(corsSettings)

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
