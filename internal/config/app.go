package config

import (
	"article/internal/delivery/http"
	"article/internal/delivery/http/route"
	"article/internal/repository"
	"article/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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
	postRepository := repository.NewPostsRepository(config.Log)

	// setup Usecase
	postUseCase := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)

	// setup Controlle

	postController := http.NewPostController(postUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		PostController: postController,
	}
	routeConfig.Setup()

}
