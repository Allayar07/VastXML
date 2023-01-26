package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) generateVMAP(c *gin.Context) {
	if err := h.service.Vast.GenerateVMAP(); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}
