package main

import (
	"github.com/bergsantana/go-contacts/internal/delivery/http"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/internal/usecase"
	"github.com/bergsantana/go-contacts/pkg/database"
	"github.com/bergsantana/go-contacts/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.NewSQLiteDB()
	repo := repository.NewContactGormRepository(db)
	uc := usecase.NewContactUsecase(repo)

	app := fiber.New()

	app.Use(middleware.SanitizeJSONBody())

	http.NewContactHandler(app, uc)

	app.Listen(":3000")
}
