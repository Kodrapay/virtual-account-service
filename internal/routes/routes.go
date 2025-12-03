package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/virtual-account-service/internal/handlers"
	"github.com/kodra-pay/virtual-account-service/internal/services"
)

func Register(app *fiber.App, service string) {
	health := handlers.NewHealthHandler(service)
	health.Register(app)

	svc := services.NewVirtualAccountService()
	h := handlers.NewVirtualAccountHandler(svc)
	api := app.Group("/virtual-accounts")
	api.Post("/", h.Create)
	api.Get("/:id", h.Get)
}
