package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"place-container/internal/delivery/http"
	"place-container/internal/delivery/http/route"
	"place-container/internal/repository"
	"place-container/internal/usecase"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	YardRepository := repository.NewYardsRepository(config.Log)
	BlockRepository := repository.NewBlocksRepository(config.Log)
	// setup Usecase
	YardUseCase := usecase.NewYardUseCase(config.DB, config.Log, config.Validate, YardRepository)
	BlockUseCase := usecase.NewBlockUseCase(config.DB, config.Log, config.Validate, BlockRepository)

	// setup Controlle

	YardController := http.NewYardController(YardUseCase, config.Log)
	BlockController := http.NewBlockController(BlockUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:             config.App,
		YardController:  YardController,
		BlockController: BlockController,
	}
	routeConfig.Setup()

}
