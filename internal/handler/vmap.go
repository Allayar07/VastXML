package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) generateVMAP(c *fiber.Ctx) error {
	if err := h.service.Vast.GenerateVMAP(); err != nil {
		return ErrorMessage(c, http.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON("OK")
}
