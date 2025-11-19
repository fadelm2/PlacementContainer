package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"place-container/internal/config"
	"place-container/internal/handler"
	"place-container/internal/repository/postgres"
	"place-container/internal/router"
	"place-container/internal/usecase"
)

func main() {
	db := config.NewPostgres()
	rd := config.NewRedis()

	// Auto migrate models

	repo := postgres.NewPostgresRepo(db)
	uc := usecase.NewYardUsecase(repo, rd)
	h := handler.NewYardHandler(uc)

	app := fiber.New()

	router.Setup(app, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("starting on :" + port)
	log.Fatal(app.Listen(":" + port))
}
