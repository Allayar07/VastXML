package handler

import (
	"golang_vast/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.Default()

	routes.POST("/vast", h.PutUri)
	routes.GET("/vast/generate", h.VastXML)

	return routes
}
