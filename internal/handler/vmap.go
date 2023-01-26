package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) VmapXML(c *gin.Context) {
	if err := h.service.GenerateVMAP(); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}
