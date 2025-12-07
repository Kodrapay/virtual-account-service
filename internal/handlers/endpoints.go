package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/virtual-account-service/internal/dto"
	"github.com/kodra-pay/virtual-account-service/internal/services"
)

type VirtualAccountHandler struct {
	svc *services.VirtualAccountService
}

func NewVirtualAccountHandler(svc *services.VirtualAccountService) *VirtualAccountHandler {
	return &VirtualAccountHandler{svc: svc}
}

func (h *VirtualAccountHandler) Create(c *fiber.Ctx) error {
	var req dto.VirtualAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	return c.JSON(h.svc.Create(c.Context(), req))
}

func (h *VirtualAccountHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // Use c.ParamsInt
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid virtual account ID")
	}
	return c.JSON(h.svc.Get(c.Context(), id))
}
