package handler

import (
	"golang_vast/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) VastXML(c *gin.Context) {
	var input model.VastModel

	if err := c.BindJSON(&input); err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Vast.Generate_Vast(input)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}
