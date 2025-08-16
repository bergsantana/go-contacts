package main

import (
	"fmt"

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

	// Middleware
	app.Use(middleware.SanitizeJSONBody())

	// Rotas
	http.NewContactHandler(app, uc)
	fmt.Println("\n📌 Available Endpoints:")
	for _, route := range app.GetRoutes(true) {
		fmt.Printf("   [%s] %s\n", route.Method, route.Path)
	}
	fmt.Println()

	app.Listen(":3000")
}
