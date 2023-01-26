package handler

import (
	"golang_vast/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Put_To_Db_VastInfo(c *gin.Context) {
	var input model.VastModel

	if err := c.BindJSON(&input); err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Vast.PutVastInfo(input)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id ": id,
	})
}

func (h *Handler) VastXML(c *gin.Context) {
	var input model.VastModel

	if err := c.BindJSON(&input); err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Vast.Generate_Vast(input.ID)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}
