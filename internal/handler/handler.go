package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang_vast/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *fiber.App {
	routes := fiber.New()

	routes.Get("/vast/:id", h.generateVast)
	routes.Get("/open/:any", h.OpenFile)
	//routes.Get("/vmap", h.generateVMAP)

	return routes
}
