package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) generateVast(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorMessage(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err = h.service.Vast.GenerateVast(id); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
