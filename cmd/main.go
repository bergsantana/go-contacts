package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bergsantana/go-contacts/internal/delivery/http"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/internal/tracing"
	"github.com/bergsantana/go-contacts/internal/usecase"
	"github.com/bergsantana/go-contacts/pkg/database"
	"github.com/bergsantana/go-contacts/pkg/middleware"
	"github.com/bergsantana/go-contacts/pkg/seed"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gorm.io/gorm"
)

func main() {
	shutdown := tracing.InitTracer()
	defer shutdown()
	db := database.NewSQLiteDB()
	handleArgs(db)

	repo := repository.NewContactGormRepository(db)
	uc := usecase.NewContactUsecase(repo)

	app := fiber.New()

	// Middleware
	app.Use(otelfiber.Middleware())
	app.Use(middleware.SanitizeJSONBody())
	app.Use(limiter.New(limiter.Config{
		Max:        15,              // Máximo de requests for minuto
		Expiration: 1 * time.Minute, // Tempo
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Limite de 15 requisições  por minuto atingido. Tente novamente mais tarde",
			})
		},
	}))

	// Rotas
	http.NewContactHandler(app, uc)
	fmt.Println("\n📌 Endpoints disponíveis:")
	for _, route := range app.GetRoutes(true) {
		fmt.Printf("   [%s] %s\n", route.Method, route.Path)
	}
	fmt.Println()

	app.Listen(":3000")
}

func handleArgs(db *gorm.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seed.SeedContacts(db)
			os.Exit(0)
		}
	}
}
